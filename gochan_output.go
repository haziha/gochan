package gochan

func (gc *GoChan[T]) Pop() (val T) {
	return <-gc.dataChan
}
