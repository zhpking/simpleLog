package main

import (
	"testing"
	"log"
)

func TestMain_PkgLog(t *testing.T) {
	log.Printf("%s\n", "test")
}

func Benchmark_PkgLog(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n ++
		log.Printf("%d\n", n)
	}
}
