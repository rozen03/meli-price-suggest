package main

func MinMaxAvgData(data obtainedData) Respuesta {
	return Respuesta{data.max, data.sum / data.total, data.min}
}

func crecentList(tam int) []int {
	list := make([]int, tam)
	for i := 0; i < tam; i++ {
		list[i] = i
	}
	return list
}

func crecentListt(tam int) []float64 {
	list := make([]float64, tam)
	for i := 0; i < tam; i++ {
		list[i] = float64(i)
	}
	return list
}
