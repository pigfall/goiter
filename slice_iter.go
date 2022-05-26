package iter

type sliceIter[T any] struct{
	slice []T
}

func (this *sliceIter[T]) Next()(Option[T],error){
	if len(this.slice) == 0 {
		return None[T](),nil
	}
	first := this.slice[0]
	this.slice = this.slice[1:]
	return Some[T](first),nil
}

func (this *sliceIter[T]) IntoIter()Iterator[T]{
	return this
}

