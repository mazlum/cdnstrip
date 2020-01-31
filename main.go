package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/user"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/mazlum/cdnstrip/cdn"

	"github.com/briandowns/spinner"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Run faster !
}

// Initialize global variables
var cdnRanges []*net.IPNet
var mutex sync.Mutex
var wg sync.WaitGroup
var validIP int
var invalidIP int
var cdnIP int
var s *spinner.Spinner = spinner.New(spinner.CharSets[11], 100*time.Millisecond)

func main() {

	cacheFilePath := getCacheFilePath()

	thread := flag.Int("t", 1, "Number of threads")
	input := flag.String("i", "", "Input [FileName|IP|CIDR]")
	out := flag.String("o", "filtered.txt", "Output file name")
	skipCache := flag.Bool("skip-cache", false, "Skip loading cache file for CDN IP ranges")
	flag.Parse()

	if *input == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Start spinner
	print("\n")
	s.Writer = os.Stdout
	s.Start()

	// First check if cache exists
	s.Suffix = " Loading for cache file..."
	cahceFile, err := ioutil.ReadFile(cacheFilePath)
	if err == nil || *skipCache {
		c := strings.Split(string(cahceFile), "\n")
		if len(c) == 0 {
			fatal(errors.New("empty cache file"))
		}
		for _, i := range c {
			_, cidr, err := net.ParseCIDR(i)
			fatal(err)
			cdnRanges = append(cdnRanges, cidr)
		}
	} else {
		// Create new cache file
		s.Suffix = " Loading all CDN ranges..."
		ranges, err := cdn.LoadAll()
		fatal(err)
		cdnRanges = ranges
		s.Suffix = " Creating new cache file..."
		cahceFile, err := os.OpenFile(cacheFilePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0664)
		fatal(err)
		for i, r := range cdnRanges {
			cahceFile.WriteString(r.String())
			if i != len(cdnRanges)-1 {
				cahceFile.WriteString("\n")
			}
		}
		cahceFile.Close()
	}

	outFile, err := os.Create(*out)
	fatal(err)
	defer outFile.Close()

	list := loadInput(*input)
	channel := make(chan string, len(list))
	for _, ip := range list {
		channel <- ip
	}
	close(channel)
	for i := 0; i < *thread; i++ {
		wg.Add(1)
		go strip(channel, outFile)
	}
	wg.Wait()

	s.Stop()
	print("[✔]" + s.Suffix + "\n")
}

func strip(channel chan string, file *os.File) {
	defer wg.Done()
	for ip := range channel {
		i := net.ParseIP(ip)
		if i != nil {
			if cdn.Check(cdnRanges, i) {
				mutex.Lock()
				cdnIP++
				mutex.Unlock()
			} else {
				mutex.Lock()
				validIP++
				file.WriteString(i.String() + "\n")
				mutex.Unlock()
			}
		} else {
			mutex.Lock()
			invalidIP++
			mutex.Unlock()
		}

		// Update spinner
		updateSpinnerStats()

	}
}

func updateSpinnerStats() {
	mutex.Lock()
	s.Suffix = "  [ VALID: " + strconv.Itoa(validIP) + " | INVALID: " + strconv.Itoa(invalidIP) + " | CDN: " + strconv.Itoa(cdnIP) + " ]"
	mutex.Unlock()
}

func getCacheFilePath() string {
	usr, err := user.Current()
	if err != nil {
		fatal(err)
	}
	return usr.HomeDir + "/.config/cdnstrip.cache"
}

func loadInput(param string) []string {
	s.Suffix = " Loading input..."

	ip := net.ParseIP(param)
	if ip != nil {
		return []string{ip.String()}
	}

	ips, err := expandCIDR(param)
	if err == nil {
		return ips
	}

	file, err := ioutil.ReadFile(param)
	fatal(err)
	return strings.Split(string(file), "\n")
}

func expandCIDR(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); incIP(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}

func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func fatal(err error) {
	if err != nil {
		s.Stop()
		pc, _, _, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		if ok && details != nil {
			log.Printf("[%s] ERROR: %s\n", strings.ToUpper(strings.Split(details.Name(), ".")[1]), err)
		} else {
			log.Printf("[UNKOWN] ERROR: %s\n", err)
		}
		os.Exit(1)
	}
}
