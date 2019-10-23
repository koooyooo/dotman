package main

import (
	"flag"
	"fmt"

	"github.com/koooyooo/fasthttp/control"
)

func main() {
	reqPerSec := flag.Int("r", 100, "requests per sec")
	sec := flag.Int("s", 0, "num sec")
	numWorkers := flag.Int("w", 1, "num workers")
	method := flag.String("m", "GET", "method")
	flag.Parse()
	url := flag.Arg(0)

	fmt.Println("requests-per-sec", *reqPerSec, "sec", *sec, "num-workers", *numWorkers, "method", *method, "url", url)
	control.Run(false, *method, url, *reqPerSec, *sec, *numWorkers)
}
