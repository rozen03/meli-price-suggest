package main

import (
	// "encoding/json"
	// "fmt"
	"reflect"
	"testing"
	// "math/rand"
	// "time"
)

func TestMeli(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

}
func TestCategoryChilds(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	id := "MLA5725"
	var cosos = []string{"MLA4711", "MLA6520", "MLA6070", "MLA86360", "MLA3381", "MLA4610", "MLA2227", "MLA86838", "MLA6537", "MLA8531", "MLA400928", "MLA1747", "MLA1771", "MLA86080", "MLA377674", "MLA4589", "MLA6177"}
	hijos := Hijos(id)
	if !reflect.DeepEqual(cosos, hijos) {
		t.Error(id + " childs have changed or there's an error getting it's childs")
	}

}

// MLA1403
// MLA1071
// MLA1367
// MLA1368
// MLA1743
// MLA1384
// MLA1246
// MLA1039
// MLA1051
// MLA1798
// MLA1648
// MLA1144
// MLA1276
// MLA5726

func BenchmarkMLA1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Suggest("MLA1000", PreciosYVentas)
	}
}

func BenchmarkMLA2547(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Suggest("MLA2547", PreciosYVentas)
	}
}

func BenchmarkMLA407134(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Suggest("MLA407134", PreciosYVentas)
	}
}

func BenchmarkMLA1574(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Suggest("MLA1574", PreciosYVentas)
	}
}

func BenchmarkMLA1499(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Suggest("MLA1499", PreciosYVentas)
	}
}

func BenchmarkMLA1459(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Suggest("MLA1459", PreciosYVentas)
	}
}

func BenchmarkMLA1182(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Suggest("MLA1182", PreciosYVentas)
	}
}

func BenchmarkMLA1168(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Suggest("MLA1168", PreciosYVentas)
	}
}

func BenchmarkMLA3025(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Suggest("MLA3025", PreciosYVentas)
	}
}
func BenchmarkMLA1132(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Suggest("MLA1132", PreciosYVentas)
	}
}
func BenchmarkMLA3937(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Suggest("MLA3937", PreciosYVentas)
	}
}

func BenchmarkMLA409431(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Suggest("MLA409431", PreciosYVentas)
	}
}
func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Suggest("MLA1430", PreciosYVentas)
	}
}
func BenchmarkMLA1540(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Suggest("MLA1540", PreciosYVentas)
	}
}
func vBenchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Suggest("", PreciosYVentas)
	}
}
func BenchmarkMLA1953(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Suggest("MLA1953", PreciosYVentas)
	}
}
func Benchmark5725(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Suggest("MLA5725", PreciosYVentas)
	}
}
