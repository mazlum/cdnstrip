# CDN Strip
[![Go Report Card](https://goreportcard.com/badge/github.com/mazlum/cdnstrip)](https://goreportcard.com/report/github.com/mazlum/cdnstrip)  [![License](https://img.shields.io/packagist/l/doctrine/orm.svg)](https://raw.githubusercontent.com/mazlum/cdnstrip/master/LICENCE)

Go module for striping CDN IP ranges.

<img align="center" src="https://github.com/mazlum/cdnstrip/raw/master/usage.gif" alt="DEMO"/>

**Currently Stripping**
- [x] Cloudflare
- [x] Cloudfront
- [x] Akamai
- [x] Azure CDN
- [x] DDOS Guard
- [x] Fastly
- [x] Incapsula
- [x] Max CDN

## Install
```
go get github.com/mazlum/cdnstrip
```

## Usage Parameters
```
  -i string
    	Input [FileName|IP|CIDR]
  -o string
    	Output file name (default "filtered.txt")
  -skip-cache
    	Skip loading cache file for CDN IP ranges
  -t int
    	Number of threads (default 1)

```

## Example Code
```
package main

import (
	"log"

	"github.com/mazlum/cdnstrip/cdn"
)

func main() {

	ip := "1.1.1.1"

	cdnRanges, err := cdn.LoadAll()
	if err != nil {
		log.Fatal(err)
	}

	if cdn.Check(ip, cdnRanges) {
		print("It's CDN IP !")
	} else {
		print("It's not CDN IP !")
	}

}
```


**Authors**
- [Mazlum Ağar](https://twitter.com/mazlumagar)
- [Ege Balcı](https://twitter.com/egeblc) 
