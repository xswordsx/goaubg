package chaoscloud

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Result is a simple <Name, Value> tuple.
type Result struct {
	Name  string
	Value string
}

func (r Result) String() string { return r.Name + "\n" }

var (
	Machine = FakeLookup("machine", "vm-vraycloud-1b-pool-4", "120.157.185.89")
	Quota   = FakeLookup("quota", "artist-senior-limit", "50000.00")
	Blob    = FakeLookup("scene", "gopher_final.vrscene", "gs://vraycloud-production/10c278e2a799ca4d8c0fb89")

	ReplicatedMachine = First(
		FakeLookup("machine_1", "vm-vraycloud-1b-pool-4", "120.157.185.89"),
		FakeLookup("machine_2", "vm-vraycloud-1b-pool-4", "120.157.185.89"),
	)
	ReplicatedQuota = First(
		FakeLookup("quota_1", "artist-senior-limit", "50000.00"),
		FakeLookup("quota_2", "artist-senior-limit", "50000.00"),
	)
	ReplicatedBlob = First(
		FakeLookup("scene_1", "gopher_final.vrscene", "gs://vraycloud-production/10c278e2a799ca4d8c0fb89"),
		FakeLookup("scene_2", "gopher_final.vrscene", "gs://vraycloud-production/10c278e2a799ca4d8c0fb89"),
	)
)

// LookupFunc is a convinent shorthand for a function that
// returns a result based on a query.
type LookupFunc func(query string) Result

// FakeLookup returns a result for resource of type kind
// after no longer than 100 milliseconds.
func FakeLookup(kind, name, value string) LookupFunc {
	return func(q string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result{
			Name:  fmt.Sprintf("%s(%q): %s", kind, q, name),
			Value: value,
		}
	}
}

// Resources looks up various required resources to start a V-Ray render.
func Resources(query string) ([]Result, error) {
	results := []Result{
		Machine(query),
		Quota(query),
		Blob(query),
	}
	return results, nil
}

// ResourcesParallel looks up various required resources to
// start a V-Ray render in parallel.
func ResourcesParallel(query string) ([]Result, error) {
	c := make(chan Result)
	go func() { c <- Machine(query) }()
	go func() { c <- Quota(query) }()
	go func() { c <- Blob(query) }()
	return []Result{<-c, <-c, <-c}, nil
}

// ResourcesTimeout looks up various required resources to
// start a V-Ray render in parallel but waits no longer than
// timeout.
func ResourcesTimeout(query string, timeout time.Duration) ([]Result, error) {
	t := time.After(timeout)
	c := make(chan Result, 3)
	go func() { c <- Machine(query) }()
	go func() { c <- Quota(query) }()
	go func() { c <- Blob(query) }()

	var results []Result
	for i := 0; i < 3; i++ {
		select {
		case <-t:
			return results, errors.New("timed out")
		case result := <-c:
			results = append(results, result)
		}
	}

	return results, nil
}

// ResourcesReplicated looks up various required resources
// to start a V-Ray render by querying multiple redundant
// backends for the information. The method waits for a
// response no longer than timeout.
func ResourcesReplicated(query string, timeout time.Duration) ([]Result, error) {
	t := time.After(timeout)
	c := make(chan Result, 3)
	go func() { c <- ReplicatedMachine(query) }()
	go func() { c <- ReplicatedQuota(query) }()
	go func() { c <- ReplicatedBlob(query) }()

	var results []Result
	for i := 0; i < 3; i++ {
		select {
		case <-t:
			return results, errors.New("timed out")
		case result := <-c:
			results = append(results, result)
		}
	}

	return results, nil
}

// First returns the result from the fastest replica from
// a given replica list.
func First(replicas ...LookupFunc) LookupFunc {
	return func(query string) Result {
		c := make(chan Result, len(replicas))
		for i := range replicas {
			go func(i int) { c <- replicas[i](query) }(i)
		}

		return <-c
	}
}
