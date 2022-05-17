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

func ForEach[T any](it Iterator[T],do func(T)error)error{
	v,err := it.Next()
	if err != nil{
		return err
	}
	for v.IsSome() {
		err := do(v.Unwrap())
		if err != nil{
			return err
		}
		v,err = it.Next()
		if err != nil{
			return err
		}
	}
	return nil
}

func ForEachWithStop[T any](it Iterator[T],do func(T)(stop bool,err error))error{
	v,err := it.Next()
	if err != nil{
		return err
	}
	for v.IsSome() {
		stop,err := do(v.Unwrap())
		if err != nil{
			return err
		}
		if stop{
			return nil
		}
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

func Collect[T any](iter Iterator[T])([]T,error){
	var length =0
	sliceIt ,ok := iter.(*sliceIter[T])
	if ok {
		length = len(sliceIt.slice)
	}
	ret := make([]T,0,length)
	err := ForEach(iter,func(value T)error{
		ret = append(ret,value)
		return nil
	})
	if err != nil{
		return nil, err
	}
	return ret,nil
}

func MustCollect[T any](iter Iterator[T])([]T){
	list,err := Collect(iter)
	if err != nil{
		panic(err)
	}
	return list
}

type Ts interface{
	float32 | ~string
}

func Filter[T any](it Iterator[T],filter func(v T)(include bool,err error))Iterator[T]{
	return &filterIter[T]{
		raw:it,
		filter:filter,
	}
}

func Exist[T any](it Iterator[T],f func(v T)(ok bool,err error))(bool,error){
	var exist bool
	err := ForEachWithStop(
		it,
		func(v T)(stop bool,err error){
			ok,err := f(v)
			if err != nil{
				return false,err
			}
			if ok {
				exist = true
				return true,nil
			}
			return false,nil
		},
	)
	return exist,err
}

func FindFisrt[T any](it Iterator[T],f func(v T)(match bool,err error))(T,bool,error){
	var find T
	var findOne bool
	err := ForEachWithStop(
		it,
		func(v T)(stop bool,err error){
			match,err := f(v)
			if err != nil{
				return false,err
			}
			if match {
				find = v
				findOne = true
				return true,nil
			}
			return false,nil
		},
	)
	return find,findOne,err
}
