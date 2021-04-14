package lock

import (
	"context"
	"fmt"
)

func (l *EtcdLock) Release() error {
	if l == nil {
		return fmt.Errorf("nil lock")
	}
	l.Lock()
	defer l.Unlock()

	defer l.session.Close()
	return l.mutex.Unlock(context.Background())
}
