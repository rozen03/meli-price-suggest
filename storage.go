package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
)

//"github.com/nobonobo/unqlitego"
// "github.com/tpotlog/unqlitego"
// "github.com/tpotlog/unqlitego/collections"
//"github.com/nobonobo/unqlitego/collections"
//"collections"

// "fmt"
//"errors"

func Categorias() []string {
	url := "https://api.mercadolibre.com/sites/MLA/categories"
	resp, err := http.Get(url)
	if err != nil {
		panic("Explotur")
	}
	defer resp.Body.Close()
	var body []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		panic("at the dis...")
	}

	res := make([]string, len(body))
	// var totales float64
	for i := range body {
		// resi := body[i].(map[string]interface{})
		res[i] = body[i]["id"].(string)
		// totales += resi["total_items_in_this_category"].(float64)
	}
	return res
}
func Hijos(category string) []string {
	url := "https://api.mercadolibre.com/categories/" + category
	resp, err := http.Get(url)
	if err != nil {
		panic("Explotur")
	}
	defer resp.Body.Close()
	var body map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		panic("at the dis...")
	}

	results := body["children_categories"].([]interface{})
	// total := body["total_items_in_this_category"].(float64)
	res := make([]string, len(results))
	// var totales float64
	for i := range results {
		resi := results[i].(map[string]interface{})
		res[i] = resi["id"].(string)
		// totales += resi["total_items_in_this_category"].(float64)
	}
	return res
}

type obtainedData struct {
	min   float64
	max   float64
	sum   float64
	total float64
}

func GetPreciosYVentas(results []interface{}) obtainedData {
	prices := 0.0
	total := 0.0
	max := 0.0
	min := math.MaxFloat64
	for i := range results {
		resi := results[i].(map[string]interface{})
		price, ok := resi["price"].(float64)
		if !ok {
			continue
		}
		sold, ok := resi["sold_quantity"].(float64)
		if !ok {
			continue
		}
		max = math.Max(price, max)
		min = math.Min(price, min)
		price += price * (sold + 1)
		total += sold
	}
	// fmt.Println(res)
	return obtainedData{min, max, prices, total}

}
func GetALLLLL(category string, offset int) obtainedData {
	url := "https://api.mercadolibre.com/sites/MLA/search?limit=200&category=" + category + "&offset=" + strconv.Itoa(offset)
	// fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		panic("Explotur")
	}
	var body map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		panic("at the dis...")
	}
	defer resp.Body.Close()
	results := body["results"].([]interface{})
	// total := body["paging"].(map[string]interface{})["total"].(float64)
	// fmt.Println(total, reflect.TypeOf(total))
	return GetPreciosYVentas(results)
}
func PreciosYVentas(category string) obtainedData {
	url := "https://api.mercadolibre.com/sites/MLA/search?limit=200&category=" + category
	resp, err := http.Get(url)
	if err != nil {
		panic("Explotur")
	}
	var body map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		fmt.Println(resp.Body)
		panic("at the dis...")
	}
	results := body["results"].([]interface{})
	total := int(body["paging"].(map[string]interface{})["total"].(float64))
	// fmt.Println(total, reflect.TypeOf(total))
	res := GetPreciosYVentas(results)
	resp.Body.Close()
	for i := 200; i < total; i += 200 {
		resi := GetALLLLL(category, i)
		res.sum += resi.sum
		res.total += resi.total
		res.max = math.Max(res.max, resi.max)
		res.min = math.Max(res.min, resi.min)
	}
	return res
}

// func PreciosYVentasPorHijo(category string) obtainedData {
// hijos := Hijos(category)
// if len(hijos) > 1 {
// fmt.Println("Hijos de", category, hijos)
// var res []item
// for _, hijo := range hijos {
// res = append(res, PreciosYVentasPorHijo(hijo)...)
// }
// fmt.Println(category, "Longitud:", len(res))
// return res
// } else {
// return PreciosYVentas(category)
// }
// }
