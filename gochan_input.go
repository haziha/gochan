package gochan

import "fmt"

func (gc *GoChan[T]) Push(val T) (err error) {
	defer func() {
		if err1 := recover(); err1 != nil {
			err = fmt.Errorf("gochan Push: %v", err1)
		}
	}()
	gc.lock.Lock()
	gc.list.PushBack(val)
	gc.lock.Unlock()

	if gc.cap < 0 {
		select {
		case gc.event <- struct{}{}:
		default:
		}
		return
	}

	gc.event <- struct{}{}

	return
}
