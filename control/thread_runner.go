package control

import (
	"time"

	"github.com/koooyooo/dotman/model"
	"github.com/koooyooo/dotman/prodcon"
)

func RunThread(isAtOnceMode bool, req model.Request, reqPerSec, sec, numWorkers int, config model.Config) {
	var client prodcon.ClientFunc
	client = prodcon.FastHttp

	for s := 0; s < sec; s++ {
		st := time.Now()
		for r := 0; r < reqPerSec; r++ {
			go client(req, config)
		}
		ed := time.Now()
		invokeSec := ed.Sub(st).Milliseconds()
		time.Sleep(time.Duration(1000-invokeSec) * time.Millisecond)
	}
}
