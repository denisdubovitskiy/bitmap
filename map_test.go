package bitmap_test

import (
	"github.com/denisdubovitskiy/bitmap"
	"testing"
)

func TestBitmap(t *testing.T) {
	const size = 10000000

	c := bitmap.New(size)

	for i := 0; i < size; i++ {
		if c.Has(i) {
			t.Fatalf("already enabled: %d", i)
		}

		c.Set(i)

		if !c.Has(i) {
			t.Fatalf("not enabled: %d", i)
		}

		c.Clear(i)

		if c.Has(i) {
			t.Fatalf("still enabled: %d", i)
		}
	}
}
