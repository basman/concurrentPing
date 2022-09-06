package main

import (
	"fmt"
	"sync"
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
	// TODO for every job provided by the input channel call Host.Ping() once
	// TODO return the number of hosts that are reachable according to Host.Reachable and the number of all jobs
	var hostCount, reachableCount int
	out := make(chan bool)
	wg := sync.WaitGroup{}
	for h := range jobs {
		hostCount++

		wg.Add(1)
		go func(h1 ping.Host) {
			h1.Ping()

			if h1.Reachable != nil && *h1.Reachable {
				out <- true
			}
			wg.Done()
		}(h)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for res := range out {
		if res {
			reachableCount++
		}
	}
	return hostCount, reachableCount
}

const hostsFilename = "hosts.csv.bz2"

func main() {
	startProg := time.Now()
	jobCh := ping.GetJobs(hostsFilename)

	startPing := time.Now()
	fmt.Printf("loaded hosts from %v\n", hostsFilename)

	hostCount, reachableCount := pingAll(jobCh)
	stopPing := time.Now()

	fmt.Printf("TIMING %v to parse %v hosts\n", startPing.Sub(startProg), hostCount)
	fmt.Printf("TIMING %v to ping %v hosts\n", stopPing.Sub(startPing), hostCount)
	fmt.Printf("RESULT %v/%v (%.6f%%) hosts reachable\n", reachableCount, hostCount, float64(reachableCount)/float64(hostCount)*100)
}
