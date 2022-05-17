package lock_manager

import (
	"sync"
	"time"
)

type LeaseManager struct {
	leases map[string]lease // Map of chunks to leases. Should we remove entry when lease expires? Or just have a boolean flag
}

type lease struct {
	m       *sync.RWMutex // A pointer to a RW Lock for a chunk
	endTime time.Time     // System Time to release lock
}

func (lm *LeaseManager) AcquireReadLease(filepath string) {

}

func (lm *LeaseManager) AcquireWriteLease(filepath string) {

}
