package main

import (
	"testing"
	"marmoteduTest2/log"
)

func TestMain_SimpleLog(t *testing.T) {
	log.Infof("%s\n","test")
}

func Benchmark_SimpleLog(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n ++
		log.Infof("%d\n", n)
	}
}
