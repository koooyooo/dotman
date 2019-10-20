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

	run(*method, url, *reqPerSec, *sec, *numWorkers, FastHttp)
}

func run(method, url string, reqPerSec, sec, numWorkers int, f ClientFunc) {
	// 終了通知、ワーク通知のストリーム準備
	doneStream := make(chan struct{}, numWorkers)
	workStream := make(chan struct{}, 0)
	defer close(doneStream)
	defer close(workStream)

	// 処理件数を管理
	var wg sync.WaitGroup

	// 消費者を稼働
	for c := 0; c < numWorkers; c++ {
		go func(name string, d, w chan struct{}) {
			for {
				select {
				case <-w:
					f(method, url)
					wg.Done()
				case <-d:
					return
				default:
				}
			}
		}("worker-"+strconv.Itoa(c), doneStream, workStream)
	}
	st := time.Now()
	if sec == 0 {
		// 全ワークを一気に投入
		wg.Add(reqPerSec)
		for j := 0; j < reqPerSec; j++ {
			workStream <- struct{}{}
		}
	} else {
		wg.Add(reqPerSec * sec)
		// 指定のタイミングで ワーク通知を投入
		for i := 0; i < sec; i++ {
			for j := 0; j < reqPerSec; j++ {
				workStream <- struct{}{}
				time.Sleep(time.Duration(1000/reqPerSec) * time.Millisecond)
			}
		}
	}
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
