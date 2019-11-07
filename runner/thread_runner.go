package runner

import (
	"sync"
	"time"

	"github.com/koooyooo/dotman/model"
	"github.com/koooyooo/dotman/prodcon"
)

func RunThread(isAtOnceMode bool, req model.Request, reqPerSec, sec, numWorkers int, config model.Config) {
	var client prodcon.ClientFunc
	client = prodcon.FastHttp

	var wg = sync.WaitGroup{}
	wg.Add(reqPerSec * sec)
	st := time.Now()
	for s := 0; s < sec; s++ {
		st := time.Now()
		for r := 0; r < reqPerSec; r++ {
			go func() {
				client(req, config)
				wg.Done()
			}()
		}
		ed := time.Now()
		invokeSec := ed.Sub(st).Milliseconds()
		time.Sleep(time.Duration(1000-invokeSec) * time.Millisecond)
	}
	wg.Wait()
	ed := time.Now()
	outputResult(st, ed, reqPerSec*sec)
}
