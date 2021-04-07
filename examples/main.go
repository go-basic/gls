package main

import (
	"fmt"
	"github.com/go-basic/gls"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			defer gls.Clean()

			defer func() {
				fmt.Printf("%d: number = %d\n", idx, gls.Get("number"))
			}()
			gls.Set("number", idx+100)
		}(i)
	}
	wg.Wait()
}
