package eventStore

import "testing"

const msg = "0123456789"

func BenchmarkPopcount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteEvent(msg)
	}
}

