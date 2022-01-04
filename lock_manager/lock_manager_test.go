package lock_manager

import (
	"testing"
)

// Let's model transactions as goroutines with sleep.Timeout in the body
func TestLockManager(t *testing.T) {
	transactionsMap = map[int64]func() int64{}
	manager := NewLockManager()

	//tr1 := func() {
	//	defer manager.
	//	time.Sleep(2000)
	//	fmt.Println(1)
	//	return
	//}

	manager.ProcessLockRequest(1, LockShared, 1234)
	manager.ProcessLockRequest(2, LockShared, 1234)
}
