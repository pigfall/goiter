package iter

type Option[T any]	struct{
	value *T
}

func Some[T any](v T) Option[T]{
	return Option[T]{value:&v}
}

func None[T any]() Option[T]{
	return Option[T]{}
}


func (opt Option[T]) IsSome() bool {
	return !opt.IsNone()
}

func (opt Option[T]) IsNone() bool {
	return opt.value == nil
}


func (opt Option[T]) Unwrap() T {
	if opt.IsNone() {
		panic("Attempted to unwrap an empty Option.")
	}
	return *opt.value
}

// MapOption applies a function fn to the contained value if it exists.
func MapOption[T any, R any](opt Option[T], fn func(T) (R,error)) (Option[R],error) {
	if !opt.IsSome() {
		return None[R](),nil
	}
	v,err := fn(opt.Unwrap())
	if err !=nil{
		return None[R](),err
	}
	return Some(v),nil
}
