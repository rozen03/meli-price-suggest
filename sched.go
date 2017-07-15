package main

import (
	"fmt"
	"meli-price-suggest/fifo"
)

type Cosa struct {
	res chan obtainedData
	url string
}

const maxChanelsSched = 300

func scheduler(ch chan Cosa) {
	queue := fifo.NewRingqueue()
	channs := make([]chan bool, maxChanelsSched)
	contador := 0
	for c := range channs {
		channs[c] = make(chan bool)
	}
	for true {
		// fmt.Println("Queue", queue)
		select {
		case resi := <-ch:
			fmt.Println("ch", <-ch)
			queue.Add(resi)
		default:
		}
		// fmt.Println(a)
		for c := range channs {
			// fmt.Print(c)
			select {
			case <-channs[c]:

				if contador < maxChanelsSched {
					mandar, ok := queue.Remove()
					//Mando la primera tanda a Descargar y calcular
					if ok {
						go GetALLLLL(mandar.(Cosa).url, mandar.(Cosa).res, channs[c])
						contador += 1
					}
				}
				contador -= 1
			default:
				continue
			}
		}
		if contador < maxChanelsSched {
			mandar, ok := queue.Remove()
			//Mando la primera tanda a Descargar y calcular
			if ok {
				go GetALLLLL(mandar.(Cosa).url, mandar.(Cosa).res, channs[contador])
				contador += 1
			}
		}

		// time.Sleep(time.Second)
	}
}
