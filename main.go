package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	ip := flag.String("ip", "", "IP address to scan")
	dns := flag.String("dns", "", "DNS name to scan")
	flag.Parse()

	if *ip == "" && *dns == "" {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf(`{"name": "example-risk", "ip": "%s", "dns": "%s", "severity": "high"}`, *ip, *dns)
}
