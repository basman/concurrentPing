package ping

import "testing"

func TestGetJobs(t *testing.T) {
	// arrange
	hostNum := 1

	// act
	hch := GetJobs("test_data/hosts_test.csv.bz2")

	var hosts []Host
	for h := range hch {
		hosts = append(hosts, h)
	}

	// assert
	if hostNum != len(hosts) {
		t.Errorf("invalid number of hosts loaded: %v, expected %v", len(hosts), hostNum)
	}

	host := hosts[0]
	if host.Address != "1.2.3.4" {
		t.Errorf("wrong host address: %v, expected %v", host.Address, "1.2.3.4")
	}

	if host.RTTms != 120 {
		t.Errorf("wrong RTT: %v, expected %v", host.RTTms, 120)
	}
}
