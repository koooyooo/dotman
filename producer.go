package main

import (
	"fmt"
	"sync"
	"time"
)

type workProducer func(int, int, chan<- struct{}, *sync.WaitGroup)

func allAtOnceProducer(reqPerSec, _ int, workStream chan<- struct{}, wg *sync.WaitGroup) {
	wg.Add(reqPerSec)
	for j := 0; j < reqPerSec; j++ {
		workStream <- struct{}{}
	}
}

func perSecDistributionProducer(reqPerSec, sec int, workStream chan<- struct{}, wg *sync.WaitGroup) {
	wg.Add(reqPerSec * sec)
	for i := 0; i < sec; i++ {
		for j := 0; j < reqPerSec; j++ {
			workStream <- struct{}{}
			time.Sleep(time.Duration(1000/reqPerSec) * time.Millisecond)
		}
		fmt.Print(" ")
	}
}
