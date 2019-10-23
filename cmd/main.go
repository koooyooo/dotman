package main

import (
	"flag"
	"fmt"
	"github.com/koooyooo/fasthttp/prodcons"
	"strconv"
	"sync"
	"time"
)

// main 処理
func main() {
	reqPerSec := flag.Int("r", 100, "num request per sec")
	sec := flag.Int("s", 0, "num sec")
	numWorkers := flag.Int("w", 1, "num numWorkers")
	method := flag.String("m", "GET", "method")
	flag.Parse()
	url := flag.Arg(0)

	if *sec == 0 {
		fmt.Println("Mode: All at once")
		fmt.Println("total-requests", *reqPerSec, "sec", *sec, "num-workers", *numWorkers, "method", *method, "url", url)
	} else {
		fmt.Println("Mode: Request per sec")
		fmt.Println("requests-per-sec", *reqPerSec, "sec", *sec, "num-workers", *numWorkers, "method", *method, "url", url)
	}
	run(*method, url, *reqPerSec, *sec, *numWorkers)
}

func run(method, url string, reqPerSec, sec, numWorkers int) {
	// 終了通知、ワーク通知のストリーム準備
	doneStream := make(chan struct{}, numWorkers)
	workStream := make(chan struct{}, 0)
	defer close(doneStream)
	defer close(workStream)

	var wg sync.WaitGroup

	// ワーカーを起動
	workConsumer := prodcons.CreateConsumer(method, url, prodcons.FastHttp, &wg)
	for c := 0; c < numWorkers; c++ {
		name := "worker-" + strconv.Itoa(c)
		go workConsumer(name, doneStream, workStream)
	}

	st := time.Now()

	// ワークを追加
	var workProd prodcons.WorkProducer = prodcons.PerSecDistributionProducer
	if sec == 0 {
		workProd = prodcons.AllAtOnceProducer
	}
	workProd(reqPerSec, sec, workStream, &wg)

	// 処理完了を待機
	wg.Wait()
	ed := time.Now()

	// 終了通知を投入
	for c := 0; c < numWorkers; c++ {
		doneStream <- struct{}{}
	}

	var totalReq = sec * reqPerSec
	if sec == 0 {
		totalReq = reqPerSec
	}
	outputResult(st, ed, totalReq)
}

func outputResult(st time.Time, ed time.Time, totalReq int) {
	fmt.Println()
	fmt.Println("Sec:", ed.Sub(st).Seconds())
	fmt.Println("Req/Sec:", float64(totalReq)/(ed.Sub(st).Seconds()))
}
