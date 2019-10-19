package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
)

func main() {

	reqPerSec := flag.Int("r", 100, "request per sec")
	sec := flag.Int("s", 1, "sec")
	concReq := flag.Int("c", 1, "concurrent requests")
	method := flag.String("m", "GET", "method")
	flag.Parse()
	url := flag.Arg(0)
	run(*method, url, *reqPerSec, *sec, *concReq, fastHttp)
	run(*method, url, *reqPerSec, *sec, *concReq, netHttp)
}

func run(method, url string, reqPerSec, sec, concReq int, f func(string, string)) {
	doneStream := make(chan struct{})
	workStream := make(chan struct{})
	for c := 0; c < concReq; c++ {
		go func(name string, d, w chan struct{}) {
			for {
				select {
				case <-w:
					f(method, url)
				case <-d:
					return
				default:
				}
			}
		}("worker-"+strconv.Itoa(c), doneStream, workStream)
	}
	st := time.Now()
	for i := 0; i < sec; i++ {
		for j := 0; j < reqPerSec; j++ {
			workStream <- struct{}{}
			time.Sleep(time.Duration(1000/reqPerSec) * time.Millisecond)
		}
	}
	for c := 0; c < concReq; c++ {
		doneStream <- struct{}{}
	}
	ed := time.Now()
	fmt.Println()
	fmt.Println("Sec:", ed.Sub(st).Seconds())
	time := ed.Sub(st).Seconds()
	numReq := sec * reqPerSec
	fmt.Println("Req/Sec:", float64(numReq)/time)
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
