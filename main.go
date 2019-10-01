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
	flag.Parse()
	url := flag.Arg(0)
	run(url, false)
	run(url, true)
}

func run(url string, fast bool) {
	reqNum := 100
	st := time.Now()
	for i := 0; i < reqNum; i++ {
		switch fast {
		case true:
			fastHttp(url)
		case false:
			normalHttp(url)
		}
	}
	ed := time.Now()
	fmt.Println()
	fmt.Println("Sec:", ed.Sub(st).Seconds())
	sec := ed.Sub(st).Seconds()
	psec := float64(reqNum) / sec
	fmt.Println("Req/Sec:", psec)

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
	fmt.Print(".")
	//fmt.Print(len(body), ",")
}
