package ping

import "time"

// Ping dummy implementation to simulate pinging with delay caused by round trip time (RTT)
func (h *Host) Ping() {
	reachable := false
	if h.RTTms > 0 && h.RTTms < 20000 {
		time.Sleep(time.Duration(h.RTTms) * time.Millisecond)
		reachable = true
		h.Reachable = &reachable
	}
}
