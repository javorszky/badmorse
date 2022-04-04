package main

import "testing"

func BenchmarkBadMorse(b *testing.B) {
	in := " aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ0123456789?.,"
	benchmarks := []struct {
		name string
		fn   func(string) string
	}{
		{
			name: "badMorse",
			fn:   badMorse,
		},
		{
			name: "badmorsewithslice",
			fn:   badMorseWithSlice,
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.fn(in)
			}
		})
	}
}
