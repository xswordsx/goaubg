package chaoscloud

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	Name  string
	Value string
}

var (
	Machine = FakeLookup("machine", "vm-vraycloud-1b-pool-4", "120.157.185.89")
	Quota   = FakeLookup("quota", "artist-senior-limit", "50000.00")
	Blob    = FakeLookup("scene", "gopher_final.vrscene", "gs://vraycloud-production/10c278e2a799ca4d8c0fb89")
)

type LookupFunc func(query string) Result

func FakeLookup(kind, name, value string) LookupFunc {
	return func(q string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result{
			Name:  fmt.Sprintf("%s(%q): %s", kind, q, name),
			Value: value,
		}
	}
}

func Resources(query string) ([]Result, error) {
	results := []Result{
		Machine(query),
		Quota(query),
		Blob(query),
	}
	return results, nil
}
