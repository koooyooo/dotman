package prodcon

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

// workConsumer (name, done channel, work channel) consumes works
type workConsumer func(string, <-chan struct{}, <-chan struct{})

// createConsumer create workConsumer
func CreateConsumer(headers map[string][]string, method string, url string, f ClientFunc, wg *sync.WaitGroup) workConsumer {
	return func(name string, d, w <-chan struct{}) {
		for {
			select {
			case <-w:
				f(headers, method, url)
				wg.Done()
			case <-d:
				return
			default:
			}
		}
	}
}

//
type ClientFunc = func(headers map[string][]string, method, url string)

func FastHttp(headers map[string][]string, method, url string) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	for k, vs := range headers {
		for _, v := range vs {
			req.Header.Add(k, v)
		}
	}

	req.Header.SetMethod(method)
	req.SetRequestURI(url)

	fasthttp.Do(req, resp)

	bodyBytes := resp.Body()
	output(bodyBytes)
}

func NetHttp(headers map[string][]string, method, url string) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	for k, vs := range headers {
		for _, v := range vs {
			req.Header.Add(k, v)
		}
	}
	tr := &http.Transport{}
	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	output(body)
}

func output(body []byte) {
	fmt.Print(".")
	fmt.Print(string(body), ",")
}
