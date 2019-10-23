package main

import (
	"flag"
	"fmt"

	"github.com/koooyooo/fasthttp/common"

	"github.com/koooyooo/fasthttp/control"
)

func main() {
	totalReqs := flag.Int("r", 100, "total requests")
	numWorkers := flag.Int("w", 1, "num workers")
	method := flag.String("m", "GET", "method")
	headers := flag.String("h", "", "headers: key1:value1,key2:value2")
	debug := flag.Bool("d", false, "debugMode")

	flag.Parse()
	url := flag.Arg(0)

	fmt.Println("total-requests", *totalReqs, "sec", 0, "num-workers", *numWorkers, "method", *method, "url", url)
	control.Run(true, common.ParseHeader(*headers), *method, url, *totalReqs, 0, *numWorkers, *debug)
}
