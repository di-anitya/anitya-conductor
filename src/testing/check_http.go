package testing

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
	"strings"
)

// transport is an http.RoundTripper that keeps track of the in-flight
// request and implements hooks to report HTTP tracing events.
type transport struct {
	current *http.Request
}

// CheckHTTPSCertificate check SSL certification
func CheckHTTPSCertificate(url string) {
	protocol := strings.Split(url, "://")[0]
	if protocol == "https" {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		result, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	} else {
		fmt.Println("skipped for CheckHTTPSCertificate because of http")
	}
}

// RoundTrip wraps http.DefaultTransport.RoundTrip to keep track
// of the current request.
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.current = req
	return http.DefaultTransport.RoundTrip(req)
}

// GotConn prints whether the connection has been used previously
// for the current request.
func (t *transport) GotConn(info httptrace.GotConnInfo) {
	fmt.Printf("Target URL: %v (refused: %v, idle time: %v)\n", t.current.URL, info.Reused, info.IdleTime)
}

// CheckHTTP is ..
func CheckHTTP() {
	url := "https://nttdata.com"
	//url := "http://www.intellilink.co.jp"
	//url := "https://news.yahoo.co.jp"
	//url := "https://mogemoge.com/"

	CheckHTTPSCertificate(url)
	t := &transport{}

	req, _ := http.NewRequest("GET", url, nil)
	trace := &httptrace.ClientTrace{
		GotConn: t.GotConn,
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	client := &http.Client{Transport: t}
	if _, err := client.Do(req); err != nil {
		log.Fatal(err)
	}
}
