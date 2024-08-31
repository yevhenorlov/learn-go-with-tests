package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("runs", func(t *testing.T) {
		got := Repeat("a", 5)
		expected := "aaaaa"
		assert(t, got, expected)
	})
	t.Run("repeats the specified amount", func(t *testing.T) {
		got := Repeat("a", 3)
		expected := "aaa"
		assert(t, got, expected)
	})
}

// to run, use ` go test -bench=.`
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}

func assert(t testing.TB, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected '%s' but got '%s'", expected, got)
	}
}
