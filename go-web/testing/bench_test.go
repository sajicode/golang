package main

import (
	"testing"
)

func BenchmarkDecode(b *testing.B) {
	//* loop through function to be benchmarked b.N times
	for i := 0; i < b.N; i++ {
		unmarshal("post.json")
	}
}

//* go test -v -cover -short -bench .
//* go test -run x -bench .
