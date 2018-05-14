/*
Exercise 1.3 Experiment to measure the dif ference in running time bet ween our potentially inefficient versions and the one that uses strings.Join. (Section 1.6 illustrates part of the time package, and Section 11.4 shows how to write benchmark tests for systematic performance evaluation.)
*/
package main

import (
	"strings"
	"testing"
)

var args = []string{"hello", "world"}

func concat(args []string) {
	r, sep := "", ""
	for _, a := range args {
		r += sep + a
		sep = " "
	}
}

func BenchmarkIsConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat(args)
	}
}

func BenchmarkIsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(args, " ")
	}
}
