package control

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/koooyooo/fasthttp/prodcon"
)

func Run(isAtOnceMode bool, headers map[string][]string, method, url string, reqPerSec, sec, numWorkers int, debug bool) {
	// 終了通知、ワーク通知のストリーム準備
	doneStream := make(chan struct{}, numWorkers)
	workStream := make(chan struct{}, 0)
	defer close(doneStream)
	defer close(workStream)

	var wg sync.WaitGroup

	// ワーカーを起動
	workConsumer := prodcon.CreateConsumer(headers, method, url, prodcon.FastHttp, &wg, debug)
	for c := 0; c < numWorkers; c++ {
		name := "worker-" + strconv.Itoa(c)
		go workConsumer(name, doneStream, workStream)
	}

	st := time.Now()

	// ワークを追加
	var workProd prodcon.WorkProducer = prodcon.PerSecProducer
	if isAtOnceMode {
		workProd = prodcon.AtOnceProducer
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
	if isAtOnceMode {
		totalReq = reqPerSec
	}
	outputResult(st, ed, totalReq)
}

func outputResult(st time.Time, ed time.Time, totalReq int) {
	fmt.Println()
	fmt.Println("Sec:", ed.Sub(st).Seconds())
	fmt.Println("Req/Sec:", float64(totalReq)/(ed.Sub(st).Seconds()))
}
