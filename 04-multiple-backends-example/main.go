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
	res := chaoscloud.First(
		chaoscloud.FakeLookup("machine", "vm-1", "1"),
		chaoscloud.FakeLookup("machine", "vm-2", "2"),
	)("golang")
	elapsed := time.Since(start)
	fmt.Println(res)
	fmt.Printf("Took: %v\n", elapsed)
}
