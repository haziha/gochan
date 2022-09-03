package gochan

func (gc *GoChan[T]) Push(val T) {
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
