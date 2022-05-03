package main

import (
	"concurrent2/ping"
	"testing"
	"time"
)

var staticTestHostList = []ping.Host{
	{"1", 0, nil},
	{"2", 10, nil},
	{"3", 10, nil},
	{"4", 10, nil},
	{"5", 10, nil},
	{"6", 10, nil},
	{"7", 10, nil},
	{"8", 0, nil},
}

func Test_pingAll(t *testing.T) {
	// arrange
	ch := make(chan ping.Host)

	go func(hlist []ping.Host, c chan ping.Host) {
		for _, h := range hlist {
			c <- h
		}
		close(c)
	}(staticTestHostList, ch)

	// act
	reachCount, hostCount := pingAll(ch)

	// assert
	if hostCount != 8 {
		t.Errorf("invalid host count %v, expected %v", hostCount, 8)
	}

	if reachCount != 6 {
		t.Errorf("invalid reachable count %v, expected %v", reachCount, 6)
	}
}

func Test_pingAll_Concurrency(t *testing.T) {
	// arrange
	ch := make(chan ping.Host)

	go func(hlist []ping.Host, c chan ping.Host) {
		for _, h := range hlist {
			c <- h
		}
		close(c)
	}(staticTestHostList, ch)

	start := time.Now()

	// act
	hostCount, reachCount := pingAll(ch)

	// assert
	end := time.Now()
	duration := end.Sub(start).Milliseconds()

	if reachCount <= 0 || hostCount <= 0 {
		t.Error("invalid counters (negative or zero)")
	}

	if duration > 15 {
		t.Errorf("pinging all hosts took too long: %v ms, expected to complete within 15 ms tops", duration)
	}

	if duration < 8 {
		t.Errorf("pinging all hosts completed too early after %v ms, should take at least 10 ms", duration)
	}
}
