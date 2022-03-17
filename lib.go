package iter

type Iterator[T any] interface {
	// Next yields a new value from the Iterator.
	Next() (Option[T],error)
}

func Slice[T any](slice []T)Iterator[T]{
	return &sliceIter[T]{
		slice:slice,
	}
}

func ForEach[T any](it Iterator[T],do func(T))error{
	v,err := it.Next()
	if err != nil{
		return err
	}
	for v.IsSome() {
		do(v.Unwrap())
		v,err = it.Next()
		if err != nil{
			return err
		}
	}
	return nil
}

func Map[From,To any](from Iterator[From],do func(From)(To,error))Iterator[To]{
	return &mapIter[From, To]{
		inner: from,
		fn:    do,
	}
}
