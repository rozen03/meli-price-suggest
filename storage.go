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

type obtainedData struct {
	min   float64
	max   float64
	sum   float64
	total float64
}

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
		fmt.Println(resp.Body)
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
		prices += price * (sold + 1)
		total += (sold + 1)
	}
	// fmt.Println(res)
	return obtainedData{min, max, prices, total}

}
func GetALLLLL(category string, offset int, c chan obtainedData) {
	url := "https://api.mercadolibre.com/sites/MLA/search?limit=200&category=" + category + "&offset=" + strconv.Itoa(offset)
	// fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		for i := 0; i < 5 && err != nil; i++ {
			fmt.Println(err)
			resp, err = http.Get(url)
		}

		if err != nil {
			panic("Explotur")
		}
	}
	var body map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		fmt.Println(resp.Body)
		panic("at the dis...")
	}

	results, ok := body["results"].([]interface{})
	resp.Body.Close()
	if !ok {
		c <- obtainedData{0.0, 0.0, 0.0, 0.0}
	}
	// total := body["paging"].(map[string]interface{})["total"].(float64)
	// fmt.Println(total, reflect.TypeOf(total))
	c <- GetPreciosYVentas(results)
}
func brezolver(res *obtainedData, resi *obtainedData) {
	(*res).sum += (*resi).sum
	(*res).total += (*resi).total
	(*res).max = math.Max((*res).max, (*resi).max)
	(*res).min = math.Max((*res).min, (*resi).min)
}
func all(respuestas []bool) bool {
	for b := range respuestas {
		if !respuestas[b] {
			return false
		}
	}
	return true
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
	var chanels int
	maxChanels := 100
	if maxChanels < total/200 {
		chanels = maxChanels
	} else {
		chanels = total / 200
	}
	chanels++
	// for i := 200; i < total; i += 200 * (chanels) {
	channs := make([]chan obtainedData, chanels)
	for c := range channs {
		channs[c] = make(chan obtainedData)
		go GetALLLLL(category, 200*c, channs[c])
	}
	// for i := 200; i < chanels; i += 200 {
	chans := chanels * 200
	respondio := make([]bool, chanels)
	for chans < total || all(respondio) {
		for c := range channs {
			select {
			case resi := <-channs[c]:
				brezolver(&res, &resi)
				chans += 200
				if chans < total {
					go GetALLLLL(category, chans, channs[c])
					respondio[c] = false
				} else {
					respondio[c] = true

				}
			default:
				continue
			}
		}

	}
	// }
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
