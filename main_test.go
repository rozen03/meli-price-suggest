package main

import
// "encoding/json"

(
	"math"
	"testing"
)

// "reflect"
const limit = 200

func TestWithOnes(t *testing.T) {
	GenerateSameTest(t, 4000.0, 1, 0)
	GenerateSameTest(t, 4000.0, 1, 50)
	GenerateSameTest(t, 4000.0, 1, 200)
	GenerateSameTest(t, 4000.0, 1, 1000)
}
func TestWith50(t *testing.T) {
	GenerateSameTest(t, 4000.0, 50, 0)
	GenerateSameTest(t, 4000.0, 50, 50)
	GenerateSameTest(t, 4000.0, 50, 200)
	GenerateSameTest(t, 4000.0, 50, 1000)
}
func TestWith100(t *testing.T) {
	GenerateSameTest(t, 4000.0, 100, 0)
	GenerateSameTest(t, 4000.0, 100, 50)
	GenerateSameTest(t, 4000.0, 100, 200)
	GenerateSameTest(t, 4000.0, 100, 1000)
}
func TestMiddleTo100(t *testing.T) {
	GenerateMiddleTest(t, 100, 0)
	GenerateMiddleTest(t, 100, 50)
	GenerateMiddleTest(t, 100, 200)
	GenerateMiddleTest(t, 100, 1000)
}
func TestMiddleTo1k(t *testing.T) {
	GenerateMiddleTest(t, 1000, 0)
	GenerateMiddleTest(t, 1000, 50)
	GenerateMiddleTest(t, 1000, 200)
	GenerateMiddleTest(t, 1000, 1000)
}

// func TestMiddleTo1M(t *testing.T) {
// GenerateMiddleTest(t, 1000000, 0)
// GenerateMiddleTest(t, 1000000, 50)
// GenerateMiddleTest(t, 1000000, 200)
// GenerateMiddleTest(t, 1000000, 1000)
// }
func TestCreciente(t *testing.T) {
	GenerateIncreasingTest(t, 50)
	GenerateIncreasingTest(t, 100)
	GenerateIncreasingTest(t, 500)
	GenerateIncreasingTest(t, 1000)
}

func TestDecreciente(t *testing.T) {
	GenerateDecreasingTest(t, 50)
	GenerateDecreasingTest(t, 100)
	GenerateDecreasingTest(t, 500)
	GenerateDecreasingTest(t, 1000)
}
func GenerateSameTest(t *testing.T, total float64, price float64, sold float64) {
	ch := startWorkers(1000)
	res := Suggest("23123", ch, func(s string) map[string]interface{} { return GenerarMismo(total, price, sold) })

	if res.max != price {
		t.Error("Max should be", price, "got", res.max)
	}
	if res.min != price {
		t.Error("Min should be ", price, " got", res.min)
	}
	if res.suggested != price {
		t.Error("Suggested should be ", price, " got", res.suggested)
	}
}

const TOLERANCE = 0.00000001

func GenerateMiddleTest(t *testing.T, hasta float64, sold float64) {
	ch := startWorkers(1000)
	res := Suggest("23123", ch, GeneradorDelMedio(hasta, sold))

	if res.max != hasta {
		t.Error("Max should be", hasta, "got", res.max)
	}
	if res.min != 0 {
		t.Error("Min should be ", 0, " got", res.min)
	}
	if diff := math.Abs(res.suggested - hasta/2); diff > TOLERANCE {
		t.Error("Suggested should be ", hasta/2, " got", res.suggested)
	}
}
func resIncreasing(hasta float64) float64 {
	//it's not super precise but works fine for this
	return (2 * (2*hasta + 1)) / 6 // = sum i*i for i in 0 to n)/(n*(n+1))/2

}
func resDecreasing(hasta float64) float64 {
	//it's not super precise but works for this
	return 2 * (hasta - 1) / 6 // = sum i*(n-i) for i in 0 to n)/(n*(n+1))/2
}
func GenerateIncreasingTest(t *testing.T, hasta float64) {
	ch := startWorkers(1000)
	res := Suggest("23123", ch, GeneradorCreciente(hasta))
	if res.max != hasta {
		t.Error("Max should be", hasta, "got", res.max)
	}
	if res.min != 0 {
		t.Error("Min should be ", 0, " got", res.min)
	}
	if diff := math.Abs(res.suggested - resIncreasing(hasta)); diff > 1 {
		t.Error("Suggested should be ", resIncreasing(hasta), " got", res.suggested)
	}
}
func GenerateDecreasingTest(t *testing.T, hasta float64) {
	ch := startWorkers(1000)
	res := Suggest("23123", ch, GeneradorDecreciente(hasta))
	if res.max != hasta {
		t.Error("Max should be", hasta, "got", res.max)
	}
	if res.min != 0 {
		t.Error("Min should be ", 0, " got", res.min)
	}
	if diff := math.Abs(res.suggested - resDecreasing(hasta)); diff > 1 {
		t.Error("Suggested should be ", resDecreasing(hasta), " got", res.suggested)
	}
}
func GenerarMismo(total float64, price float64, soldCount float64) map[string]interface{} {
	var prices [200]float64
	var sold [200]float64
	for i := range prices {
		prices[i] = price
		sold[i] = soldCount
	}
	return Generar(total, prices, sold)
}
func GeneradorDelMedio(hasta float64, soldCount float64) func(s string) map[string]interface{} {
	total := (hasta + 1) * 200
	contador := -1.0
	return func(s string) map[string]interface{} {
		contador++
		var prices [200]float64
		var sold [200]float64
		for i := range prices {
			prices[i] = contador
			sold[i] = soldCount
		}
		return Generar(total, prices, sold)
	}
}
func GeneradorCreciente(hasta float64) func(s string) map[string]interface{} {
	total := (hasta + 1) * 200
	contador := -1.0
	return func(s string) map[string]interface{} {
		contador++
		var prices [200]float64
		var sold [200]float64
		for i := range prices {
			prices[i] = contador
			sold[i] = contador
		}
		return Generar(total, prices, sold)
	}
}
func GeneradorDecreciente(hasta float64) func(s string) map[string]interface{} {
	total := (hasta + 1) * 200
	contador := -1.0
	return func(s string) map[string]interface{} {
		contador++
		var prices [200]float64
		var sold [200]float64
		for i := range prices {
			prices[i] = contador
			sold[i] = hasta - contador
		}
		return Generar(total, prices, sold)
	}
}
func Generar(total float64, prices [200]float64, sold [200]float64) map[string]interface{} {
	maa := make(map[string]interface{})
	paging := make(map[string]interface{})
	paging["total"] = total
	results := make([]interface{}, limit)
	for i := 0; i < limit; i++ {
		results[i] = make(map[string]interface{})
		results[i].(map[string]interface{})["price"] = prices[i]
		results[i].(map[string]interface{})["sold_quantity"] = sold[i]
	}
	maa["paging"] = paging
	maa["results"] = results
	// muu := maa.(map[string]interface{})
	return maa
}
