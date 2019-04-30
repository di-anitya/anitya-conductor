package testing

import (
	"fmt"
	"net"
	"os"
	"strings"
)

// CheckDNSLookup for checking DNS lookup
func CheckDNSLookup(url string) {
	dns := strings.Split(url, "://")[1]

	ips, err := net.LookupIP(dns)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		os.Exit(1)
	}
	for _, ip := range ips {
		fmt.Printf("%s IN A %s\n", dns, ip.String())
	}
}

// CheckDNS is ..
func CheckDNS() {
	url := "https://nttdata.com"
	//url := "http://www.intellilink.co.jp"
	//url := "https://news.yahoo.co.jp"
	//url := "https://mogemoge.com/"

	CheckDNSLookup(url)
}
