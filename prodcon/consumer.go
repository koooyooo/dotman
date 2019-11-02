package prodcon

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/koooyooo/dotman/model"

	"github.com/valyala/fasthttp"
)

// workConsumer (name, done channel, work channel) consumes works
type workConsumer func(string, <-chan struct{}, <-chan struct{})

// createConsumer create workConsumer
func CreateConsumer(req model.Request, f ClientFunc, wg *sync.WaitGroup, config model.Config) workConsumer {
	return func(name string, d, w <-chan struct{}) {
		for {
			select {
			case <-w:
				f(req, config)
				wg.Done()
			case <-d:
				return
			}
		}
	}
}

//
type ClientFunc = func(req model.Request, c model.Config)

func FastHttp(r model.Request, c model.Config) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	for k, vs := range r.Headers {
		for _, v := range vs {
			req.Header.Add(k, v)
		}
	}

	req.Header.SetMethod(r.Method)
	req.SetRequestURI(r.Url)
	st := time.Now()
	fasthttp.Do(req, resp)
	ed := time.Now()
	if c.VerboseTime {
		fmt.Printf("[%d]", ed.Sub(st).Milliseconds())
	}
	bodyBytes := resp.Body()
	output(bodyBytes, c.VerboseResponse)
}

func NetHttp(r model.Request, c model.Config) {
	req, err := http.NewRequest(r.Method, r.Url, nil)
	if err != nil {
		log.Fatal(err)
	}
	for k, vs := range r.Headers {
		for _, v := range vs {
			req.Header.Add(k, v)
		}
	}
	tr := &http.Transport{}
	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	output(body, c.VerboseResponse)
}

func output(body []byte, debug bool) {
	out := "."
	if debug {
		out = string(body)
	}
	fmt.Print(out)
}
