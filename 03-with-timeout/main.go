package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/xswordsx/goaubg/chaoscloud"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	start := time.Now()
	res, err := chaoscloud.ResourcesTimeout("job-id-12345", 80*time.Millisecond)
	elapsed := time.Since(start)
	fmt.Println(res, err)
	fmt.Printf("Took: %v\n", elapsed)
}
