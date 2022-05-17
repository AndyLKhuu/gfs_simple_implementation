package lease_manager

import (
	"time"
)

type LeaseManager struct {
	leases map[uint64]leaseInfo
}

type leaseInfo struct {
	primary string    // Chunkserver which holds the lease
	endTime time.Time // System Time to release lock
}

func (lm *LeaseManager) AcquireOrCreateLease(ch uint64, chunkServers []string) {
	l, ok := lm.leases[ch]
	if !ok {
		// TO:DO Better heuristic to give lease to chunkserver under lower load.
		lease := leaseInfo{primary: chunkServers[0], endTime: time.Now().Add(time.Second * 60)}
		lm.leases[ch] = lease
	}
}

// TO:DO Create a easy switch to toggle between current lease manager and a lease manager which gives leases on reads?
