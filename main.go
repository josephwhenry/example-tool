package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	ip := flag.String("ip", "", "IP address to scan")
	dns := flag.String("dns", "", "DNS name to scan")
	severity := flag.String("severity", "", "Severity of the risk")
	flag.Parse()

	if *ip == "" && *dns == "" {
		flag.Usage()
		os.Exit(1)
	}

	if *severity == "" {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf(`{"name": "example-risk", "ip": "%s", "dns": "%s", "severity": "%s"}`, *ip, *dns, *severity)
}
