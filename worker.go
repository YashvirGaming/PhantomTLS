package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var total int
var success int
var failed int

func startWorkerMode() {
	var wg sync.WaitGroup

	workers := 20
	target := ""

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < 10; j++ {
				total++
				resp, err := http.Get(target)
				if err != nil {
					failed++
				} else {
					success++
					resp.Body.Close()
				}
			}
		}()
	}

func startProxyListMode() {
	println("Load proxies from proxies.txt (ip:port:user:pass)")
}

	go liveCounter()

	wg.Wait()
}

func liveCounter() {
	for {
		fmt.Printf("\rRequests: %d | Success: %d | Fail: %d",
			total, success, failed)
		time.Sleep(1 * time.Second)
	}
}
