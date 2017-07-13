package main

func MinMaxAvgData(data obtainedData) Respuesta {
	return Respuesta{data.max, data.sum / data.total, data.min}
}

func MinMaxAvg(list []float64) Respuesta {
	sum := 0.0
	min := list[0]
	max := list[0]
	for _, val := range list {
		sum += val
		if val > max {
			max = val
		} else if val < min {
			min = val
		}
	}
	return Respuesta{max, sum / float64(len(list)), min}
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
