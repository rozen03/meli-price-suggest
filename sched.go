package main

type ArgsAndResult struct {
	res      chan obtainedData
	args     string
	download Downloader
}

const maxChanelsSched = 60

/*
**A Task Worker is a Goroutine that is listening to a channel waiting for
**a new task to do, in this case the Task would be to download, canculate
**and send to the thread creator of the channel
**the minimum, maximum, total count and the sum of all proces
**
 */

func taskWorker(ch chan ArgsAndResult, morir chan bool, workerId int) {
	for true {
		select {
		case resi := <-ch:
			GetObtainedData(resi.args, resi.res, resi.download)
		default:
			// time.Sleep(time.Second / 2)
		}
	}
}
func matar(morir chan bool, workers int) {
	for i := 0; i < workers; i++ {
		morir <- true
	}
}
func startWorkers(workers int) (chan ArgsAndResult, chan bool) {
	ch := make(chan ArgsAndResult)
	morir := make(chan bool)
	for i := 0; i < workers; i++ {
		go taskWorker(ch, morir, i)
	}
	return ch, morir
}
