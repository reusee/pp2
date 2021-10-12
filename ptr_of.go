package pp2

func PtrOf[T any](v T) *T {
	return &v
}
