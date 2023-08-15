package tuple

func First[F any] (f F, _... any) F {
	return f
}

func Second[F any, S any] (_ F, s S, _... any) S {
	return s
}

