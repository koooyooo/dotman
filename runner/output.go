package runner

import (
	"fmt"
	"time"
)

func outputResult(st time.Time, ed time.Time, totalReq int) {
	fmt.Println()
	fmt.Println("Sec:", ed.Sub(st).Seconds())
	fmt.Println("Req/Sec:", float64(totalReq)/(ed.Sub(st).Seconds()))
}
