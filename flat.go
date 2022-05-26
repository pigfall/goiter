package iter

type flatIter[T any] struct{
	iter Iterator[Iterator[T]]
	curIter Iterator[T]
}

func Flatten[T any](iter Iterator[Iterator[T]])Iterator[T]{
	return &flatIter[T]{iter:iter}
}

func (this *flatIter[T]) Next()(Option[T],error){
	if this.curIter == nil{
		iter,err := this.iter.Next()
		if err != nil{
			return None[T](),err
		}
		if iter.IsNone(){
			return None[T](),nil
		}
		this.curIter = iter.Unwrap()
	}

	i,err := this.curIter.Next()
	if err != nil{
		return None[T](),nil
	}
	if i.IsNone(){
		this.curIter = nil
		return this.Next()
	}
	return i,nil
}

func FlattenSlice[T any](list [][]T)[]T{
	ret := make([]T,0,len(list))
	for _,v := range list{
		ret = append(ret, v...)
	}
	return ret
}

func ListSliceToIter[T any](list [][]T)[]Iterator[T]{
	ret := make([]Iterator[T],0,len(list))
	for _,v := range list{
		ret = append(ret,Slice(v))
	}
	return ret
}
