package gochan

import (
	"github.com/haziha/golist"
	"sync"
)

func New[T any](_cap int) (goChan *GoChan[T]) {
	chanCap := _cap
	if chanCap < 0 {
		chanCap = 0
	}
	goChan = &GoChan[T]{
		list:     golist.New[T](),
		event:    make(chan struct{}, 0),
		dataChan: make(chan T, chanCap),
		closed:   false,
		cap:      _cap,
	}
	go goChan.goroutine()
	return
}

type GoChan[T any] struct {
	lock     sync.RWMutex
	list     *golist.List[T]
	event    chan struct{}
	dataChan chan T
	cap      int

	closed bool
	once   sync.Once
}

func (gc *GoChan[T]) Cap() (c int) {
	c = gc.cap
	return
}

func (gc *GoChan[T]) Len() (l int) {
	gc.lock.RLock()
	defer gc.lock.RUnlock()

	l = len(gc.dataChan)
	if gc.cap < 0 {
		l += gc.list.Len()
	}

	return
}

func (gc *GoChan[T]) goroutine() {
	defer func() {
		recover()
	}()
	for !gc.closed {
		gc.lock.Lock()
		if gc.list.Len() == 0 {
			gc.lock.Unlock()
			<-gc.event
		} else {
			val := gc.list.Front().Value
			gc.list.Remove(gc.list.Front())
			gc.lock.Unlock()
			gc.dataChan <- val
		}
	}
}

func (gc *GoChan[T]) Close() {
	gc.once.Do(func() {
		gc.closed = true
		close(gc.event)
		close(gc.dataChan)
	})
}
