package ip

import (
	"slices"
	"testing"
)

func TestStringScorer(t *testing.T) {
	ss := NewStringScorer(10)
	ss.Set("a")
	ss.Set("b")
	ss.Set("c")
	ss.AddScore("a")
	ss.AddScore("b")
	ss.AddScore("b")
	ss.AddScore("c")
	ss.AddScore("c")
	ss.AddScore("c")

	if r := ss.data; !slices.Equal(r, []string{"c", "b", "a"}) {
		t.Fatalf("expected [c b a], got %v", r)
	}

	ss.AddScore("a")
	if r := ss.data; !slices.Equal(r, []string{"c", "b", "a"}) {
		t.Fatalf("expected [c b a], got %v", r)
	}
	ss.AddScore("a")
	if r := ss.data; !slices.Equal(r, []string{"c", "b", "a"}) {
		t.Fatalf("expected [c b a], got %v", r)
	}
	ss.AddScore("a")
	if r := ss.data; !slices.Equal(r, []string{"a", "c", "b"}) {
		t.Fatalf("expected [a c b], got %v", r)
	}
}
