package pp2

import "testing"

func TestMapFilterSrc(t *testing.T) {
	var values Values[int]
	if err := Copy[IntSrc, IntSink](
		MapFilterSrc(
			Seq[IntSrc](1, 2, 3),
			func(i int) *int {
				if i == 2 {
					return nil
				}
				return &i
			},
			nil,
		),
		CollectValues[IntSink](&values),
	); err != nil {
		t.Fatal(err)
	}
	if len(values) != 2 {
		t.Fatal()
	}
	if values[0] != 1 {
		t.Fatal()
	}
	if values[1] != 3 {
		t.Fatal()
	}
}

func TestMapFilterSink(t *testing.T) {
	var values Values[int]
	if err := Copy[IntSrc, IntSink](
		Seq[IntSrc](1, 2, 3),
		MapFilterSink(
			CollectValues[IntSink](&values),
			func(i int) *int {
				if i == 2 {
					return nil
				}
				return &i
			},
		),
	); err != nil {
		t.Fatal(err)
	}
	if len(values) != 2 {
		t.Fatal()
	}
	if values[0] != 1 {
		t.Fatal()
	}
	if values[1] != 3 {
		t.Fatal()
	}
}
