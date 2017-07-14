package main

import "meli-price-suggest/fifo"

type Cosa struct {
	res chan obtainedData
	url string
}

func scheduler(ch chan Cosa) {
	queue := fifo.NewRingqueue()
	for true {

		select {
		case resi := <-ch:
			queue.Add(resi)
		default:
			continue
		}
		mandar, ok := queue.Remove()
		if ok {
			go GetALLLLL(mandar.(Cosa).url, mandar.(Cosa).res)
		}
	}
}
