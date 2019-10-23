package main

import (
	"flag"
	"fmt"

	"github.com/koooyooo/fasthttp/control"
)

func main() {
	totalReqs := flag.Int("r", 100, "total requests")
	numWorkers := flag.Int("w", 1, "num workers")
	method := flag.String("m", "GET", "method")
	flag.Parse()
	url := flag.Arg(0)

	fmt.Println("total-requests", *totalReqs, "sec", 0, "num-workers", *numWorkers, "method", *method, "url", url)
	control.Run(true, *method, url, *totalReqs, 0, *numWorkers)
}
