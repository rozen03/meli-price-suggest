package main

import (
	// "encoding/json"
	"fmt"
	"testing"
	// "reflect"
	"math/rand"
	"time"
)

const meliId = "MLA5726"

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

//
// func Benchmark10Workers(b *testing.B) {
// ch := startWorkers(10)
// for i := 0; i < b.N; i++ {
// Suggest(meliId, ch, Download)
// }
// }
// func Benchmark20Workers(b *testing.B) {
// ch := startWorkers(20)
// for i := 0; i < b.N; i++ {
// Suggest(meliId, ch, Download)
// }
// }
// func Benchmark30Workers(b *testing.B) {
// ch := startWorkers(30)
// for i := 0; i < b.N; i++ {
// Suggest(meliId, ch, Download)
// }
// }
func Benchmark40Workers(b *testing.B) {
	ch := startWorkers(40)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, Download)
	}
}
func Benchmark50Workers(b *testing.B) {
	ch := startWorkers(50)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, Download)
	}
}
func Benchmark60Workers(b *testing.B) {
	ch := startWorkers(60)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, Download)
	}
}
func Benchmark70Workers(b *testing.B) {
	ch := startWorkers(70)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, Download)
	}
}
func Benchmark100Workers(b *testing.B) {
	ch := startWorkers(100)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, Download)
	}
}
func Benchmark200Workers(b *testing.B) {
	ch := startWorkers(200)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, Download)
	}
}
func Benchmark500Workers(b *testing.B) {
	ch := startWorkers(500)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, Download)
	}
}
func Benchmark1000Workers(b *testing.B) {
	ch := startWorkers(1000)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, Download)
	}
}
func Benchmark2000Workers(b *testing.B) {
	ch := startWorkers(2000)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, Download)
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
func TestWithOnes(t *testing.T) {
	ch := startWorkers(1000)
	fmt.Println(Suggest("23123", ch, GenerarUnos))

}

func GenerarUnos(category string) map[string]interface{} {
	maa := make(map[string]interface{})
	paging := make(map[string]interface{})
	paging["total"] = 400000.0
	limit := 200
	results := make([]interface{}, limit)
	for i := 0; i < limit; i++ {
		results[i] = make(map[string]interface{})
		results[i].(map[string]interface{})["price"] = 1.0
		results[i].(map[string]interface{})["sold_quantity"] = 0.0
	}
	maa["paging"] = paging
	maa["results"] = results
	// muu := maa.(map[string]interface{})
	return maa
}
