package pp2

// Put puts an element to Sink
func Put[
	Sink ~func(*T) (Sink, error),
	T any,
](sink *Sink, value *T) (err error) {
	if sink == nil || *sink == nil {
		return nil
	}
	*sink, err = (*sink)(value)
	return
}
