package runner

import (
	"strconv"
	"sync"
	"time"

	"github.com/koooyooo/dotman/model"

	"github.com/koooyooo/dotman/prodcon"
)

func RunWorker(isAtOnceMode bool, req model.Request, reqPerSec, sec, numWorkers int, config model.Config) {
	// 終了通知、ワーク通知のストリーム準備
	doneStream := make(chan struct{}, numWorkers)
	workStream := make(chan struct{}, 0)
	defer close(doneStream)
	defer close(workStream)

	var wg sync.WaitGroup

	// ワーカーを起動
	workConsumer := prodcon.CreateConsumer(req, prodcon.FastHttp, &wg, config)
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
