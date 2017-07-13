package main

import (
	// "encoding/json"
	"fmt"
	"testing"
	// "reflect"
	"math/rand"
	"time"
)

func Test001(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

}

func BenchmarkXxx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

func BenchmarkShuffle(b *testing.B) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for i := 0; i < b.N; i++ {
		list = shuffle(list)
	}
}

func shuffle(slice []int) []int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	for i := range slice {
		j := random.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
