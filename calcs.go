package main

import "math"

/*
**returns min, max and the average
 */
func MinMaxAvgData(data obtainedData) Respuesta {
	return Respuesta{data.max, data.sum / data.total, data.min}
}

/*
**get two references to obtainedData
**and saves in the first the min between it's max valures,
**the max between it's max valures
**the sum between the two total counts and prices
 */
func MergeObainedData(res *obtainedData, resi *obtainedData) {
	(*res).sum += (*resi).sum
	(*res).total += (*resi).total
	(*res).max = math.Max((*res).max, (*resi).max)
	(*res).min = math.Min((*res).min, (*resi).min)
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

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
