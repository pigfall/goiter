package iter

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


type flatIter[T any] struct{
	inner Iterator[[]T]
	index int
	curList []T
}

func (this *flatIter[T]) Next()(Option[T],error){
	if this.curList == nil{
		list,err := this.inner.Next()
		if err != nil{
			return None[T](),nil
		}
		if list.IsNone(){
			return None[T](),nil

		}
		this.curList = list.Unwrap()
		this.index = 0
	}
	if this.index >= len(this.curList){
		this.curList = nil
		this.index =0
		return this.Next()
	}
	v := this.curList[this.index]
	this.index++
	return Some(v),nil
}
