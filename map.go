package pp2

func MapSrc[
	Src ~func() (*T, Src, error),
	T any,
](
	src Src,
	fn func(T) T,
	cont Src,
) Src {
	var ret Src
	ret = func() (*T, Src, error) {
		value, err := Get(&src)
		if err != nil {
			return nil, nil, err
		}
		if value == nil && src == nil {
			return nil, cont, nil
		}
		if value != nil {
			*value = fn(*value)
		}
		return value, ret, nil
	}
	return ret
}

func MapSink[
	Sink ~func(*T) (Sink, error),
	T any,
](
	sink Sink,
	fn func(T) T,
) Sink {
	var ret Sink
	ret = func(value *T) (Sink, error) {
		if value != nil && sink == nil {
			return nil, ErrShortSink
		}
		var err error
		if value != nil {
			sink, err = sink(PtrOf(fn(*value)))
		} else {
			sink, err = sink(nil)
		}
		if err != nil {
			return nil, err
		}
		if sink == nil {
			return nil, nil
		}
		return ret, nil
	}
	return ret
}
