package main

import (
	"flag"
	"fmt"

	"github.com/koooyooo/dotman/common"
	"github.com/koooyooo/dotman/model"

	"github.com/koooyooo/dotman/control"
)

func main() {
	totalReqs := flag.Int("r", 100, "total requests")
	numWorkers := flag.Int("w", 1, "num workers")
	method := flag.String("m", "GET", "method")
	headers := flag.String("h", "", "headers: key1:value1,key2:value2")
	verboseResponse := flag.Bool("vr", false, "verbose output of response")
	verboseTime := flag.Bool("vt", false, "verbose output of time")

	flag.Parse()
	url := flag.Arg(0)

	fmt.Println("total-requests", *totalReqs, "sec", 0, "num-workers", *numWorkers, "method", *method, "url", url)
	control.RunWorker(
		true,
		model.Request{
			Headers: common.ParseHeader(*headers),
			Method:  *method,
			Url:     url,
		},
		*totalReqs,
		0,
		*numWorkers,
		model.Config{
			VerboseResponse: *verboseResponse,
			VerboseTime:     *verboseTime,
		})
}
