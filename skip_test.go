package pp2

import "testing"

func TestSkip(t *testing.T) {
	var values Values[int]
	if err := Copy(
		SkipSrc(
			Seq[IntSrc](1, 2, 3, 4, 5),
			1,
			nil,
		),
		CollectValues[IntSink](&values),
	); err != nil {
		t.Fatal(err)
	}
	if len(values) != 4 {
		t.Fatal()
	}
	if values[0] != 2 {
		t.Fatal()
	}
}
