package iter

type Iterator[T any] interface {
	// Next yields a new value from the Iterator.
	Next() Option[T]
}

