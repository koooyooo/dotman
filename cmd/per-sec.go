package main

import (
	"flag"
	"fmt"

	"github.com/koooyooo/dotman/model"

	"github.com/koooyooo/dotman/common"

	"github.com/koooyooo/dotman/runner"
)

func main() {
	reqPerSec := flag.Int("r", 10, "requests per sec")
	sec := flag.Int("s", 10, "num sec")
	numWorkers := flag.Int("w", 1, "num workers, 0 workers mean thread per message mode")
	method := flag.String("m", "GET", "method")
	headers := flag.String("h", "", "headers: key1:value1,key2:value2")
	body := flag.String("b", "", "request body")
	verboseResponse := flag.Bool("vr", false, "verbose output of response")
	verboseTime := flag.Bool("vt", false, "verbose output of time")

	flag.Parse()
	url := flag.Arg(0)

	threadPerMessageMode := *numWorkers <= 0

	fmt.Println("thread-per-msg-mode", threadPerMessageMode, "requests-per-sec", *reqPerSec, "sec", *sec, "num-workers", *numWorkers, "method", *method, "url", url)
	if threadPerMessageMode {
		runner.RunThread(
			true,
			model.Request{
				Headers: common.ParseHeader(*headers),
				Method:  *method,
				Url:     url,
				Body:    []byte(*body),
			},
			*reqPerSec,
			*sec,
			*numWorkers,
			model.Config{
				VerboseResponse: *verboseResponse,
				VerboseTime:     *verboseTime,
			})
	} else {
		runner.RunWorker(
			false,
			model.Request{
				Headers: common.ParseHeader(*headers),
				Method:  *method,
				Url:     url,
				Body:    []byte(*body),
			},
			*reqPerSec,
			*sec,
			*numWorkers,
			model.Config{
				VerboseResponse: *verboseResponse,
				VerboseTime:     *verboseTime,
			})
	}
}
