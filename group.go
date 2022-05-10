package iter

func GroupAsMap[T any ,I comparable](it Iterator[T], groupGetter func(v T) (I) ) (map[I][]T,error){
	m := make(map[I][]T)
	err := ForEach(
			it,
			func(v T)(error){
				id := groupGetter(v)
				if m[id] == nil{
					m[id] = make([]T,0,1)
				}
				m[id] = append(m[id], v)
				return nil
			},
	)
	if err != nil{
		return nil, err
	}
	return m, nil
}

func GroupAsMap_A [T any ,I comparable, V any] (it Iterator[T],groupGetter func(t T)(I,V)) (map[I][]V,error){
	m := make(map[I][]V)
	err := ForEach(
			it,
			func(t T)(error){
				id,v := groupGetter(t)
				if m[id] == nil{
					m[id] = make([]V,0,1)
				}
				m[id] = append(m[id], v)
				return nil
			},
	)
	if err != nil{
		return nil, err
	}
	return m, nil
}


func GroupAsSlice[T any, I comparable](it Iterator[T],groupGetter func(v T) (I) )([][]T,error){
	groupAsMap,err := GroupAsMap(it,groupGetter)
	if err != nil{
		return nil, err
	}
	list := make([][]T,0,len(groupAsMap))
	for _,g := range groupAsMap{
		list = append(list,g)
	}
	return list,nil
}
