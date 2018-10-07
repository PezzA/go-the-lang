package main

import "testing"

var items []string

func init() {
	items = []string{"quite", "a", "few", "arguments", "", "", "some", "extra", "arguments", "to", "do"}
}

func BenchmarkPrintArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printArgs(items)
	}
}

func BenchmarkTurboArgsPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TurboArgsPrint(items)
	}
}
