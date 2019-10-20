package main

import (
	"flag"
	"fmt"
	"strconv"
	"sync"
	"time"
)

// main 処理
func main() {
	reqPerSec := flag.Int("r", 100, "num request per sec")
	sec := flag.Int("s", 0, "num sec")
	numWorkers := flag.Int("c", 1, "num numWorkers")
	method := flag.String("m", "GET", "method")
	flag.Parse()
	url := flag.Arg(0)

	fmt.Println("ReqPerSec", *reqPerSec, "sec", *sec, "numWorkers", *numWorkers, "method", *method, "url", url)

	run(*method, url, *reqPerSec, *sec, *numWorkers)
}

func run(method, url string, reqPerSec, sec, numWorkers int) {
	// 終了通知、ワーク通知のストリーム準備
	doneStream := make(chan struct{}, numWorkers)
	workStream := make(chan struct{}, 0)
	defer close(doneStream)
	defer close(workStream)

	var wg sync.WaitGroup

	workConsumer := createConsumer(method, url, FastHttp, &wg)
	for c := 0; c < numWorkers; c++ {
		go workConsumer("worker-"+strconv.Itoa(c), doneStream, workStream)
	}
	st := time.Now()
	var prod workProducer
	if sec == 0 {
		prod = allAtOnceProducer
	} else {
		prod = perSecDistributionProducer
	}
	prod(reqPerSec, sec, workStream, &wg)
	// 処理完了を待機
	wg.Wait()
	ed := time.Now()

	// 終了通知を投入
	for c := 0; c < numWorkers; c++ {
		doneStream <- struct{}{}
	}
	outputResult(st, ed, sec*reqPerSec)
}

func outputResult(st time.Time, ed time.Time, totalReq int) {
	fmt.Println()
	fmt.Println("Sec:", ed.Sub(st).Seconds())
	time := ed.Sub(st).Seconds()
	fmt.Println("Req/Sec:", float64(totalReq)/time)
}
