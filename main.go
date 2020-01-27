package main

import (
	"fmt"

	cdn "github.com/mazlum/cdnstrip/cdn"
)

func main() {
	// Test
	// cf, err := cdn.LoadCloudflare()

	// cf, err := cdn.LoadCloudFront()
	// fmt.Print(cf)
	// fmt.Println(err)

	// fs, err2 := cdn.LoadFastly()
	// fmt.Print(fs)
	// fmt.Println(err2)

	dg, err := cdn.LoadAkamai()
	fmt.Print(dg)
	fmt.Println(err)

}
