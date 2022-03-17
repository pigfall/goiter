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

