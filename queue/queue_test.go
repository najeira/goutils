package queue

import "testing"

func TestQueueLength(t *testing.T) {
	q := New()

	if q.Length() != 0 {
		t.Error("empty queue length not 0")
	}

	for i := 0; i < 1000; i++ {
		q.Add(i)
		if q.Length() != i+1 {
			t.Error("adding: queue with", i, "elements has length", q.Length())
		}
	}
	for i := 0; i < 1000; i++ {
		q.Pop()
		if q.Length() != 1000-i-1 {
			t.Error("removing: queue with", 1000-i-i, "elements has length", q.Length())
		}
	}
}

func TestQueuePeekOutOfRangePanics(t *testing.T) {
	q := New()

	assertPanics(t, "should panic when peeking empty queue", func() {
		q.Peek()
	})

	q.Add(1)
	q.Pop()

	assertPanics(t, "should panic when peeking emptied queue", func() {
		q.Peek()
	})
}

func TestQueueRemoveOutOfRangePanics(t *testing.T) {
	q := New()

	assertPanics(t, "should panic when removing empty queue", func() {
		q.Pop()
	})

	q.Add(1)
	q.Pop()

	assertPanics(t, "should panic when removing emptied queue", func() {
		q.Pop()
	})
}

func assertPanics(t *testing.T, name string, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%s: didn't panic as expected", name)
		} else {
			t.Logf("%s: got panic as expected: %v", name, r)
		}
	}()

	f()
}

// General warning: Go's benchmark utility (go test -bench .) increases the number of
// iterations until the benchmarks take a reasonable amount of time to run; memory usage
// is *NOT* considered. On my machine, these benchmarks hit around ~1GB before they've had
// enough, but if you have less than that available and start swapping, then all bets are off.

func BenchmarkQueueSerial(b *testing.B) {
	q := New()
	for i := 0; i < b.N; i++ {
		q.Add(nil)
	}
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func BenchmarkQueueTickTock(b *testing.B) {
	q := New()
	for i := 0; i < b.N; i++ {
		q.Add(nil)
		q.Pop()
	}
}
