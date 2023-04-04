package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Request struct {
	Method   string
	URL      *url.URL
	Header   http.Header
	Body     io.ReadCloser
	Close    bool
	Response *http.Response
}

type Response struct {
	Status     string
	StatusCode int // 200
	Header     http.Header
	Body       io.ReadCloser
	Request    *http.Request
}

type ReadCloser interface {
	Reader
	Closer
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

func main() {
	resp, err := http.Get("https://www.google.com/search?q=dog&hl=en&ei=Eo8iZO_tJeiF4-EPv--yaA&ved=0ahUKEwjvvvechf79AhXowjgGHb-3DA0Q4dUDCA8&uact=5&oq=dog&gs_lcp=Cgxnd3Mtd2l6LXNlcnAQAzIQCC4QgwEQ1AIQsQMQigUQQzIQCC4QigUQsQMQgwEQ1AIQQzIHCAAQigUQQzIHCAAQigUQQzIHCAAQigUQQzIHCAAQigUQQzIHCAAQigUQQzIICC4QgAQQ1AIyBAgAEAMyBQgAEIAEOg8IABCKBRDqAhC0AhBDGAE6DwguEIoFEOoCELQCEEMYAToSCC4QigUQ1AIQ6gIQtAIQQxgBOggILhCKBRCRAjoICAAQigUQkQI6CAgAEIoFELEDOhEILhCKBRCxAxCDARDHARDRAzoLCAAQigUQsQMQgwE6EQguEIoFELEDEIMBEMcBEK8BOg0IABCKBRCxAxCDARBDOgoIABCKBRCxAxBDOgcILhCKBRBDOg0ILhCKBRDHARDRAxBDOg0ILhCKBRCxAxDUAhBDSgQIQRgAUABYtghg0QloAXABeACAAZ4BiAHAA5IBAzAuM5gBAKABAbABCsABAdoBBAgBGAc&sclient=gws-wiz-serp")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	byteSlice := make([]byte, 999999)
	_, myErr := resp.Body.Read(byteSlice)
	// fmt.Println(myErr)

	if myErr == nil {
		fmt.Println(string(byteSlice))
	}
	fmt.Println(hello)
}

var hello = "string"
