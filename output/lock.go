package output

import "sync"

var screenMutex = sync.Mutex{}

func LockScreen() {
	screenMutex.Lock()
}

func UnlockScreen() {
	screenMutex.Unlock()
}
