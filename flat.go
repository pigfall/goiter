package iter

// TODO
//type flatIter[T any,I Iterator[[]T]] struct{
//	iter Iterator[I]
//	curIter Iterator[T]
//}
//
//func Flat[T any,I Iterator[[]T]](iter Iterator[I])Iterator[T]{
//	return &flatIter[T,I]{iter:iter}
//}
//
//func (this *flatIter[T,I]) Next()(Option[T],error){
//	if this.curIter == nil{
//		iter,err := this.iter.Next()
//		if err != nil{
//			return None[T](),err
//		}
//		if iter.IsNone(){
//			return None[T](),nil
//		}
//		this.curIter = iter.Unwrap()
//	}
//
//	i,err := this.curIter.Next()
//	if err != nil{
//		return None[T](),nil
//	}
//	if i.IsNone(){
//		this.curIter = nil
//		return this.Next()
//	}
//	return i,nil
//}

func Flatten[T any](list [][]T)[]T{
	ret := make([]T,0,len(list))
	for _,v := range list{
		ret = append(ret,v...)
	}

	return ret
}
