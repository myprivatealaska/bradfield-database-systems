package lock_manager

import (
	"fmt"
	"sync"
)

type LockType int

var transactionsMap map[int64]func() int64

const (
	LockShared LockType = iota
	LockExclusive
)

type lockMeta struct {
	Granted []int64
	// an array of transaction ids, ordered
	WaitQueue []int64
}

type LockManager struct {
	mu    sync.Mutex
	locks map[int64]map[LockType]lockMeta
}

func NewLockManager() *LockManager {
	return &LockManager{
		mu:    sync.Mutex{},
		locks: map[int64]map[LockType]lockMeta{},
	}
}

func (lm *LockManager) ProcessLockRequest(transactionID int64, lockType LockType, resourceID int64) bool {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	if locks, exists := lm.locks[resourceID]; exists {
		switch lockType {
		case LockShared:
			if locksShared, ok := locks[LockShared]; ok {
				newGranted := append(locksShared.Granted, transactionID)
				lm.locks[resourceID][LockShared] = lockMeta{
					Granted:   newGranted,
					WaitQueue: lm.locks[resourceID][LockShared].WaitQueue,
				}
			} else {
				lm.locks[resourceID][LockShared] = lockMeta{
					Granted:   []int64{transactionID},
					WaitQueue: []int64{},
				}
			}
			return true
		case LockExclusive:
			if locksExclusive, ok := locks[LockExclusive]; ok {
				newWaitQ := append(locksExclusive.WaitQueue, transactionID)
				lm.locks[resourceID][LockExclusive] = lockMeta{
					Granted:   lm.locks[resourceID][LockExclusive].Granted,
					WaitQueue: newWaitQ,
				}
				copy(lm.locks[resourceID][LockExclusive].WaitQueue, newWaitQ)
				return false
			} else {
				lm.locks[resourceID][LockExclusive] = lockMeta{
					Granted:   []int64{transactionID},
					WaitQueue: []int64{},
				}
				return true
			}
		default:
			panic(fmt.Sprintf("Lock type not supported: %v", lockType))
		}
	} else {
		lm.locks[resourceID] = map[LockType]lockMeta{}
		lm.locks[resourceID][lockType] = lockMeta{
			Granted:   []int64{transactionID},
			WaitQueue: []int64{},
		}
		return true
	}
}

// TODO: it lock granted, execute the operation and release the lock
// 	     if lock enqueued, enter a for loop with timeout, check repeatedly until the lock can be granted, only then - return
//		 build a job for dependency graph check. how do I kill a transaction if it's causing a deadlock?

//func (lm *LockManager) ReleaseLock(lockType LockType, resourceID int64) {
//	for _, l := range lm.locks[resourceID] {
//		if l.
//	}
//}
