package main

import (
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
)

func main()  {
	run(1, 5 * time.Second,"https://www.youtube.com/watch?v=f7MzFfuNOtY")

	wait := make(chan bool)

	<-wait
}

func run(concurrency int, interval time.Duration, url string) {
	for i:=0; i<concurrency;i++  {
		go func() {
			delay := time.Duration(int64(math.Round(interval.Seconds() * rand.Float64()))) * time.Second

			time.Sleep(delay)

			for true {
				rsp, err := http.Get(url)
				if err != nil {
					log.Print(err)
				}

				log.Print(rsp.Header)

				time.Sleep(interval)
			}
		}()
	}
}
