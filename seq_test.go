package pp2

import "testing"

func TestSeq(t *testing.T) {
	var n int
	if err := Copy(
		Seq[IntSrc](1, 2, 3),
		CountSink[IntSink](&n),
	); err != nil {
		t.Fatal(err)
	}
	if n != 3 {
		t.Fatalf("got %d", n)
	}

	if err := Copy[IntSrc, IntSink](
		Seq[IntSrc](1, 2, 3),
		Discard[IntSink],
	); err != nil {
		t.Fatal(err)
	}
}
