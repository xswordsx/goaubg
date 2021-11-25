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
	res, err := chaoscloud.ResourcesReplicated("golang", 80*time.Millisecond)
	elapsed := time.Since(start)
	fmt.Println(res)
	fmt.Printf("Took: %v %v\n", elapsed, err)
}
