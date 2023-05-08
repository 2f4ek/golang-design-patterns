package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Singleton struct {
}

var inst *Singleton

func GetInstance() *Singleton {
	if inst == nil {
		lock.Lock()
		defer lock.Unlock()
		inst = &Singleton{}

		fmt.Println("Creating singleton instance now.")
	} else {
		fmt.Println("singleton instance already created.")
	}

	return inst
}
