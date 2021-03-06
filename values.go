package pp2

type Values[T any] []T

func IterValues[
	Src ~func() (*T, Src, error),
	T any,
](v Values[T], cont Src) Src {
	n := 0
	var src Src
	src = func() (*T, Src, error) {
		if n >= len(v) {
			return nil, cont, nil
		}
		value := &v[n]
		n++
		return value, src, nil
	}
	return src
}

func CollectValues[
	Sink ~func(*T) (Sink, error),
	T any,
](p *Values[T]) Sink {
	return Tap[Sink](func(v T) error {
		*p = append(*p, v)
		return nil
	})
}
