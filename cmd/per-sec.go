package main

import (
	"flag"
	"fmt"

	"github.com/koooyooo/dotman/model"

	"github.com/koooyooo/dotman/common"

	"github.com/koooyooo/dotman/control"
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
	control.RunWorker(
		false,
		model.Request{
			Headers: common.ParseHeader(*headers),
			Method:  *method,
			Url:     url,
		},
		*reqPerSec,
		*sec,
		*numWorkers,
		*debug)
}
