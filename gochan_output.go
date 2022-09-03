package gochan

func (gc *GoChan[T]) Pop() (val T, ok bool) {
	val, ok = <-gc.dataChan
	return
}
