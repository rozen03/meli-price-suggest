package main

type Respuesta struct {
	max       float64 `json:max`
	suggested float64 `json:suggested`
	min       float64 `json:min`
}
type ItemsPorCategoria func(category string) obtainedData
type Items struct {
	PreciosYVentas ItemsPorCategoria
}

func Suggest(category string, PreciosYVentas ItemsPorCategoria) Respuesta {
	data := PreciosYVentas(category)
	// fmt.Println(data)
	return MinMaxAvgData(data)
}
