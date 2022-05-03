package ping

import (
	"testing"
	"time"
)

func TestHost_Ping(t *testing.T) {
	// arrange
	h1 := Host{
		RTTms:     120,
		Reachable: nil,
	}

	ping := make(chan Host)
	go func(h2 Host) {
		h2.Ping()
		ping <- h2
		close(ping)
	}(h1)

	start := time.Now()

	// act
	var h3 Host
	select {
	case <-time.After(150 * time.Millisecond):
		t.Errorf("timeout after 150ms")
		return
	case h3 = <-ping:
		duration := time.Now().Sub(start).Milliseconds()
		if duration < 115 {
			t.Errorf("ping completed too early, after %v ms", duration)
		}
	}

	// assert
	if h3.Reachable == nil || !*h3.Reachable {
		t.Errorf("reachability not set or wrong value: %v", h3.Reachable)
	}
}
