package main

import (
	"flag"
	"fmt"

	"github.com/koooyooo/fasthttp/common"

	"github.com/koooyooo/fasthttp/control"
)

func main() {
	reqPerSec := flag.Int("r", 10, "requests per sec")
	sec := flag.Int("s", 10, "num sec")
	numWorkers := flag.Int("w", 1, "num workers")
	method := flag.String("m", "GET", "method")
	headers := flag.String("h", "", "headers: key1:value1,key2:value2")
	debug := flag.Bool("d", false, "debugMode")

	flag.Parse()
	url := flag.Arg(0)

	fmt.Println("requests-per-sec", *reqPerSec, "sec", *sec, "num-workers", *numWorkers, "method", *method, "url", url)
	control.Run(false, common.ParseHeader(*headers), *method, url, *reqPerSec, *sec, *numWorkers, *debug)
}
