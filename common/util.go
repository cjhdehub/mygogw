package common

import (
	"fmt"
	"sync"
)

var mu = sync.Mutex{}

var UUIDMAP = make(map[string]int)

func UUID(key string) string {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := UUIDMAP[key]; !ok {
		UUIDMAP[key] = 0
	}

	UUIDMAP[key]++
	return fmt.Sprint(UUIDMAP[key])
}
