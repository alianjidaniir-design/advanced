package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		return
	}
	URL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println("Error in parsing URL:", err)
		return
	}
	c := &http.Client{
		Timeout: time.Second * 10,
	}
	request, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		fmt.Println("Error in creating request:", err)
		return
	}
	httpData, err := c.Do(request)
	if err != nil {
		fmt.Println("Error in doing request:", err)
		return
	}

	fmt.Println("Status code:", httpData.Status)
	header, _ := httputil.DumpResponse(httpData, true)
	fmt.Println(string(header))
	contentType := httpData.Header.Get("Content-Type")
	charset := strings.SplitAfter(contentType, "charest=")
	fmt.Println(charset)
	if len(charset) > 1 {
		fmt.Println("Charset:", charset[1])
	}
	if httpData.ContentLength == -1 {
		fmt.Println("ContentLength is invalid")
	} else {
		fmt.Println("ContentLength:", httpData.ContentLength)
	}
	leng := 0
	var buffer [1024]byte
	r := httpData.Body
	for {
		n, err := r.Read(buffer[0:])
		if err != nil {
			fmt.Println("Error in reading body:", err)
			break
		}
		leng += n
	}
	fmt.Println("Calculated response data length:", leng)
}
