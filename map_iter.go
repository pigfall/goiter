package iter

type mapIter[T, R any] struct {
	inner Iterator[T]
	fn    func(T) (R,error)
}

func (it *mapIter[T, R]) Next() (Option[R],error) {
	v,err := it.inner.Next()
	if err != nil{
		return None[R](),err
	}
	return MapOption(v, it.fn)
}
