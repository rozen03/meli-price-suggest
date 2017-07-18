package main

import
// "encoding/json"

"testing"

// "reflect"
const limit = 200

func TestWithOnes0Sold(t *testing.T) {
	ch := startWorkers(1000)
	res := Suggest("23123", ch, func(s string) map[string]interface{} { return GenerarUnos(4000.0, 0) })

	if res.max != 1.0 {
		t.Error("Max should be 1 got", res.max)
	}
	if res.min != 1.0 {
		t.Error("Min should be 1 got", res.min)
	}
	if res.suggested != 1.0 {
		t.Error("Suggested should be 1 got", res.suggested)
	}
}
func GenerarUnos(total float64, soldCount float64) map[string]interface{} {
	var prices [200]float64
	var sold [200]float64
	for i := range prices {
		prices[i] = 1
		sold[i] = soldCount
	}
	return Generar(total, prices, sold)
}
func GeneradorCreciente(hasta float64) func(s string) map[string]interface{} {
	total := hasta * 200
	contador := 0.0
	return func(s string) map[string]interface{} {
		contador++
		var prices [200]float64
		var sold [200]float64
		for i := range prices {
			prices[i] = 10 + contador
			sold[i] = contador
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
