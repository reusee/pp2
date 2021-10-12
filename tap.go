package pp2

func Tap[
	Sink interface {
		~func(*T) (Sink, error)
	},
	T any,
](fn func(T) error) Sink {
	var sink Sink
	sink = func(v *T) (Sink, error) {
		if v == nil {
			return nil, nil
		}
		if v != nil {
			if err := fn(*v); err != nil {
				return nil, err
			}
		}
		return sink, nil
	}
	return sink
}