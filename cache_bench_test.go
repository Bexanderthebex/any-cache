package anycache

import (
	"testing"
)

func BenchmarkGet(b *testing.B) {
	mep := make(map[any]any)
	for i := 0; i < 1_000_000; i++ {
		mep[i] = "sample"
	}
	cache := New()
	cache.ReloadSnapshot(mep)

	b.ResetTimer()
	for i := 0; i < b.N; i += 4 {
		cache.Get(i)
	}
}

func BenchmarkSample(b *testing.B) {
	mep := make(map[int]string)
	for i := 0; i < 1_000_000; i++ {
		mep[i] = "sample"
	}
	cache := NewMap()
	cache.ReloadSnapshot(mep)

	b.ResetTimer()
	for i := 0; i < b.N; i += 4 {
		cache.Get(i)
	}
}

func BenchmarkReloadSnapshotAnyCache(b *testing.B) {
	mep := make(map[int]string)
	for i := 0; i < 1_000_000; i++ {
		mep[i] = "sample"
	}
	cache := New()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.ReloadSnapshot(mep)
	}
}

func BenchmarkReloadSnapshotNative(b *testing.B) {
	mep := make(map[int]string)
	for i := 0; i < 1_000_000; i++ {
		mep[i] = "sample"
	}
	cache := NewMap()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.ReloadSnapshot(mep)
	}
}
