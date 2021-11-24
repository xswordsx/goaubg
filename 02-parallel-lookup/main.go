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
	res, err := chaoscloud.ResourcesParallel("job-id-12345")
	elapsed := time.Since(start)
	fmt.Println(res, err)
	fmt.Printf("Took: %v\n", elapsed)
}
