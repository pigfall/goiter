package iter

import(
	"fmt"
)

func ToMapWithNoRepeatId[T any,I  comparable](it Iterator[T],idGetter func(v T)(I))(map[I]T,error){
	m := make(map[I]T)
	err := ForEach(
		it,
		func(v T)error{
			id:=idGetter(v)
			_,ok := m[id]
			if ok{
				// repeated id
				return fmt.Errorf("repeated id %+v",id)
			}
			m[id] = v
			return nil
		},
	)
	if err != nil{
		return nil, err
	}

	return m, nil
}
