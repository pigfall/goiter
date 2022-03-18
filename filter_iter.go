package iter

type filterIter[T any] struct{
	raw Iterator[T]
	filter func(value T)(include bool,err error)
}

func (this filterIter[T])Next()(Option[T],error){
	for{
		v,err := this.raw.Next()
		if err != nil{
			return None[T](),err
		}
		if v.IsNone(){
			return None[T](),nil
		}

		include,err := this.filter(v.Unwrap())
		if err != nil{
			return None[T](),nil
		}
		if !include {
			continue
		}

		return Some[T](v.Unwrap()),nil
	}
}
