package ping

import (
	"compress/bzip2"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

type Host struct {
	Address   string // typically an IP address
	RTTms     int    // round trip time for from sending a ICMP echo request until the response has been received (0 means infinite RTT)
	Reachable *bool  // indicate host's reachability as an outcome of the response time
}

var jobs []Host
var jobsMu sync.Mutex

func GetJobs(filename string) chan Host {
	jobsMu.Lock()
	defer jobsMu.Unlock()

	if len(jobs) == 0 {
		var err error
		jobs, err = loadJobs(filename)
		if err != nil {
			log.Fatalf("failed to load jobs from %v: %v", filename, err)
		}
	}

	ch := make(chan Host)

	go func(c chan Host) {
		for _, h := range jobs {
			c <- h
		}
		close(c)
	}(ch)

	return ch
}

func loadJobs(filename string) ([]Host, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	bz := bzip2.NewReader(f)
	r := csv.NewReader(bz)

	var hosts []Host
	recIdx := 0
	for {
		recIdx++
		fields, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("CSV line %v: %v", recIdx, err)
		}

		addr, rttStr := fields[0], fields[1]

		rtt, err := strconv.Atoi(rttStr)
		if err != nil {
			return nil, fmt.Errorf("CSV line %v: could not parse RTT '%v' as a number", recIdx, rttStr)
		}

		hosts = append(hosts, Host{
			Address: addr,
			RTTms:   rtt,
		})
	}

	return hosts, nil
}
