package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
)

type obtainedData struct {
	min   float64
	max   float64
	sum   float64
	total float64
}

const maxChanels = 63
const melink = "https://api.mercadolibre.com/sites/MLA/search?limit=200&category="

func download(url string) map[string]interface{} {
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
	resp.Body.Close()
	return body
}
func GetPricesAndSold(results []interface{}) obtainedData {
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
	return obtainedData{min, max, prices, total}

}
func GetALLLLL(url string, c chan obtainedData) {
	body := download(url)
	results, ok := body["results"].([]interface{})
	if !ok {
		c <- obtainedData{0.0, 0.0, 0.0, 0.0}
	}

	c <- GetPricesAndSold(results)
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
func GetTotal(body *map[string]interface{}) int {
	return int((*body)["paging"].(map[string]interface{})["total"].(float64))
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func PreciosYVentas(category string, ch chan Cosa) obtainedData {
	url := melink + category
	body := download(url)
	results := body["results"].([]interface{})
	total := GetTotal(&body)
	res := GetPricesAndSold(results)
	chanels := min(maxChanels, total/200) + 1
	channs := make([]chan obtainedData, chanels)
	//Mando la primera tanda a Descargar y calcular
	for c := range channs {
		channs[c] = make(chan obtainedData)
		// go GetALLLLL(melink+category+"&offset="+strconv.Itoa(200*c), channs[c])
		ch <- Cosa{channs[c], melink + category + "&offset=" + strconv.Itoa(200*c)}
	}

	if maxChanels == chanels {
		return res
	}
	done := chanels * 200
	respondio := make([]bool, chanels)
	for done < total || all(respondio) {
		for c := range channs {
			select {
			case resi := <-channs[c]:
				brezolver(&res, &resi)
				done += 200
				if done < total {
					ch <- Cosa{channs[c], melink + category + "&offset=" + strconv.Itoa(200*c)}
					// go GetALLLLL(melink+category+"&offset="+strconv.Itoa(done), channs[c])
					respondio[c] = false
				} else {
					respondio[c] = true

				}
			default:
				continue
			}
		}
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
