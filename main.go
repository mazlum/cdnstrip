package main

import (
	"fmt"
	cdn "github.com/mazlum/cdnstrip/cdn"
)

func main() {
	// Test
	// cf, err := cdn.LoadCloudflare()
	cf, err := cdn.LoadMaxCdn()
	fmt.Print(cf)
	fmt.Println(err)
}