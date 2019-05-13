package monitoring

import (
	"fmt"
	"net"
	"os"
	"strings"
)

// CheckDNSLookup for checking DNS lookup
func CheckDNSLookup(url string) (bool, string) {
	var status bool
	var message string
	var result string

	_dns := strings.Split(url, "://")[1]
	dns := strings.Split(_dns, "/")[0]

	fmt.Println("start CheckDNSLookup", dns)

	ips, err := net.LookupIP(dns)
	if err != nil {
		//fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		status = false
		message = err.Error()
		os.Exit(1)
	}
	for _, ip := range ips {
		result += dns + " IN A " + ip.String() + "\\n"
		//fmt.Printf("%s IN A %s\n", dns, ip.String())
	}
	status = true
	message = result
	return status, message
}

// RunDNSValification is ..
func RunDNSValification(url string) (bool, string) {
	fmt.Println("start RunDNSValification")
	status, message := CheckDNSLookup(url)
	return status, message
}
