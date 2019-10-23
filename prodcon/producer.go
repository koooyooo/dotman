package prodcon

import (
	"sync"
	"time"
)

type WorkProducer func(int, int, chan<- struct{}, *sync.Cond, *sync.WaitGroup)

func AtOnceProducer(reqPerSec, _ int, workStream chan<- struct{}, cond *sync.Cond, wg *sync.WaitGroup) {
	wg.Add(reqPerSec)
	for j := 0; j < reqPerSec; j++ {
		cond.L.Lock()
		workStream <- struct{}{}
		cond.L.Unlock()
		cond.Broadcast()
	}
}

func PerSecProducer(reqPerSec, sec int, workStream chan<- struct{}, cond *sync.Cond, wg *sync.WaitGroup) {
	wg.Add(reqPerSec * sec)
	for i := 0; i < sec; i++ {
		for j := 0; j < reqPerSec; j++ {
			cond.L.Lock()
			workStream <- struct{}{}
			cond.L.Unlock()
			cond.Broadcast()
			time.Sleep(time.Duration(1000/reqPerSec) * time.Millisecond)
		}
		//fmt.Print(" ")
	}
}
