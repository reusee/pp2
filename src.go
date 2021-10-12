package pp2

type IntSrc func() (*int, IntSrc, error)
