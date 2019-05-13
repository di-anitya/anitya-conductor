package monitoring

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
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
func CheckHTTPSCertificate(url string) (bool, string) {
	var status bool
	var message string

	protocol := strings.Split(url, "://")[0]
	if protocol == "https" {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		result, err := http.Get(url)
		if err != nil {
			//fmt.Println(err)
			status = false
			message = err.Error()
		}
		//fmt.Println(result)

		responseData, err := ioutil.ReadAll(result.Body)
		if err != nil {
			status = false
			message = err.Error()
		}

		status = true
		message = string(responseData)
	} else {
		//fmt.Println("skipped for CheckHTTPSCertificate because of http")
		status = true
		message = "skipped for CheckHTTPSCertificate because of http"
	}
	return status, message
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

// CheckHTTPRequest check SSL certification
func CheckHTTPRequest(url string) (bool, string) {
	var status bool
	var message string

	t := &transport{}

	req, _ := http.NewRequest("GET", url, nil)
	trace := &httptrace.ClientTrace{
		GotConn: t.GotConn,
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	client := &http.Client{Transport: t}
	result, err := client.Do(req)
	if err != nil {
		//log.Fatal(err)
		status = false
		message = err.Error()
	} else {
		status = true
		responseData, err := ioutil.ReadAll(result.Body)
		if err != nil {
			status = false
			message = err.Error()
		}

		status = true
		message = string(responseData)
	}
	return status, message
}

// RunHTTPValification is ..
func RunHTTPValification(url string) (bool, string) {
	status_https_certificate, result_https_certificate := CheckHTTPSCertificate(url)
	status_http_request, result_http_request := CheckHTTPRequest(url)

	var status bool
	var result string
	result += "CheckHTTPSCertificate:\n" + result_https_certificate
	result += "CheckHTTPRequest:\n" + result_http_request

	if status_https_certificate && status_http_request {
		status = true
	} else {
		status = false
	}

	return status, result
}
