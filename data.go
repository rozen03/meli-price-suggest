package main

import "net/http"

type Respuesta struct {
	max       float64 `json:max`
	suggested float64 `json:suggested`
	min       float64 `json:min`
}
type httpInterface func(category string) (*http.Response, error)

/*
**Download the recieved category prices
**returns min, max and the average  of all prices in this  category
 */
func Suggest(category string, ch chan ArgsAndResult, get httpInterface) Respuesta {
	data := PreciosYVentas(category, ch, get)
	return MinMaxAvgData(data)
}
