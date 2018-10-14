package main

import "time"

func main() {
	lastUsed := time.Now()
	idle := time.Second * 3
	tickr := time.NewTicker(idle)
	i := 0
	for t := range tickr.C {
		lastUsed = lastUsed.Add(time.Second * 4)
		println("before", t.Before(lastUsed), "after", t.After(lastUsed))
		if i > 3 {
			tickr.Stop()
			return
		}
		i++
	}
	println("finished")
}
