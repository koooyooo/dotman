package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/valyala/fasthttp"
)

func main() {
	numReq := flag.Int("n", 100, "numRequest")
	concReq := flag.Int("c", 1, "concurrent requests")
	method := flag.String("m", "GET", "method")
	flag.Parse()
	url := flag.Arg(0)
	run(*method, url, *numReq, *concReq, fastHttp)
	run(*method, url, *numReq, *concReq, netHttp)
}

func run(method, url string, numReq, concReq int, f func(string, string)) {
	//workStream := make(chan struct{})

	st := time.Now()
	for i := 0; i < numReq; i++ {
		f(method, url)
	}
	ed := time.Now()
	fmt.Println()
	fmt.Println("Sec:", ed.Sub(st).Seconds())
	sec := ed.Sub(st).Seconds()
	fmt.Println("Req/Sec:", float64(numReq)/sec)
}

func fastHttp(method, url string) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.Header.SetMethod(method)
	req.SetRequestURI(url)

	fasthttp.Do(req, resp)

	bodyBytes := resp.Body()
	operation(bodyBytes)
}

func netHttp(method, url string) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	tr := &http.Transport{}
	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	operation(body)
}

func operation(body []byte) {
	fmt.Print(".")
	//fmt.Print(len(body), ",")
}
