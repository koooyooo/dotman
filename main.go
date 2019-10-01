package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/valyala/fasthttp"
)

func main() {
	url := ""
	run(url, false)
	run(url, true)
}

func run(url string, fast bool) {
	st := time.Now()
	for i := 0; i < 100; i++ {
		switch fast {
		case true:
			fastHttp(url)
		case false:
			normalHttp(url)
		}
	}
	ed := time.Now()
	fmt.Println()
	fmt.Println(ed.Sub(st))
}

func fastHttp(url string) {

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.Header.SetMethod("GET")
	req.SetRequestURI(url)

	fasthttp.Do(req, resp)

	bodyBytes := resp.Body()
	operation(bodyBytes)
}

func normalHttp(url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	tr := &http.Transport{}
	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	operation(body)
}

func operation(body []byte) {
	fmt.Print(len(body), ",")
}
