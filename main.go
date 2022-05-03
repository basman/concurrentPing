package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
   Challenge: Implement concurrency in pingAll() to complete all jobs as fast as possible.
   Rule:      Only change main.go, do not change ANY OTHER file.
   Definition of done: ALL unit tests in this project show green, ping/*_test.go included.
*/

import (
	"concurrent2/ping"
)

func pingAll(jobs chan ping.Host) (int, int) {
	var hostCount int

	wg := sync.WaitGroup{}
	var reachCount int64
	for h := range jobs {
		hostCount++
		wg.Add(1)

		go func(h1 ping.Host) {
			h1.Ping()
			if h1.Reachable != nil && *h1.Reachable {
				atomic.AddInt64(&reachCount, 1)
			}
			wg.Done()
		}(h)
	}

	wg.Wait()
	return int(reachCount), hostCount
}

const hostsFilename = "hosts.csv.bz2"

func main() {
	startProg := time.Now()
	jobCh := ping.GetJobs(hostsFilename)

	startPing := time.Now()
	fmt.Printf("loaded hosts from %v\n", hostsFilename)

	reachableCount, hostCount := pingAll(jobCh)
	stopPing := time.Now()

	fmt.Printf("TIMING %v to parse %v hosts\n", startPing.Sub(startProg), hostCount)
	fmt.Printf("TIMING %v to ping %v hosts\n", stopPing.Sub(startPing), hostCount)
	fmt.Printf("RESULT %v/%v (%.6f%%) hosts reachable\n", reachableCount, hostCount, float64(reachableCount)/float64(hostCount)*100)
}
