package pp2

func Discard[
	Sink ~func(*T) (Sink, error),
	T any,
](v *T) (Sink, error) {
	if v == nil {
		return nil, nil
	}
	return Discard[Sink], nil
}
