// The race condition in this code occurs when the shared variable is accessed and modified by 
// the goroutines at the same time under the concurrent execution. This causes duplicate IDs.
// We use mutex lock to synchronize access to the shared variable 'current'. The mutex is unlocked using defer when the function exits.
// Mutex ensures only one goroutine runs the critical section at a time, which is the read-increment-write operation of this variable.

package counter

import (
	"sync"
	"time"
)

var current int64
var mutex sync.Mutex

func NextID() int64 {
	mutex.Lock()
	defer mutex.Unlock()
	id := current
	time.Sleep(0)
	current++
	return id
}
