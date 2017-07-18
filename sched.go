package main

import "time"
import "fmt"

type ArgsAndResult struct {
	res      chan obtainedData
	args     string
	download Downloader
}

const maxChanelsSched = 64

/*
**A Task Worker is a Goroutine that is listening to a channel waiting for
**a new task to do, in this case the Task would be to download, canculate
**and send to the thread creator of the channel
**the minimum, maximum, total count and the sum of all proces
**
 */

func taskWorker(ch chan ArgsAndResult, workerId int) {
	for true {
		select {
		case resi := <-ch:
			GetObtainedData(resi.args, resi.res, resi.download)
		default:
			fmt.Println("soy", workerId, "y estoy al pedo")
			time.Sleep(time.Second / 2)

		}
	}
}

func startWorkers(workers int) chan ArgsAndResult {
	ch := make(chan ArgsAndResult)
	for i := 0; i < workers; i++ {
		go taskWorker(ch, i)
	}
	return ch
}
