package main

type Respuesta struct {
	max       float64 `json:max`
	suggested float64 `json:suggested`
	min       float64 `json:min`
}
type Downloader func(category string) map[string]interface{}

/*
**Download the recieved category prices
**returns min, max and the average  of all prices in this  category
 */
func Suggest(category string, ch chan ArgsAndResult, download Downloader) Respuesta {
	data := PreciosYVentas(category, ch, download)
	return MinMaxAvgData(data)
}
