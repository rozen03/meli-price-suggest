package main

import (
	// "encoding/json"
	"fmt"
	"testing"
	// "reflect"
	"math/rand"
	"time"
)

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
	fmt.Println(Suggest("23123", ch, func(s string) map[string]interface{} { return GenerarUnos(400000.0) }))
}
func GenerarUnos(total float64) map[string]interface{} {
	prices := make([]float64, 200)
	sold := make([]float64, 200)
	for i := range prices {
		prices[i] = 1
		sold[i] = 0
	}
	return
}
func Generar(total float64, prices [200]float64, sold [200]float64) map[string]interface{} {
	maa := make(map[string]interface{})
	paging := make(map[string]interface{})
	paging["total"] = total
	limit := 200
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
