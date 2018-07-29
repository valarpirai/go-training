package main

import (
	"fmt"
	"sync"
	"time"

	"./mytest"
	"github.com/jochasinga/requests"
)

func main() {

	var wg sync.WaitGroup

	fmt.Println("Hello World!")
	fmt.Println("My First Go Program!")

	go printNum()
	go printNum()

	mytest.Test()

	wg.Add(1)
	go func() {
		url := "https://gist.githubusercontent.com/valarpirai/473305f09f8433f1d338634ed42c437d/raw/8a09237e123bd832999b1968c47cf371cad6e8ec/live-radio.json"

		fmt.Println(url)
		timeout := func(r *requests.Request) {

			// Set Timeout on *Request instead of *http.Client
			r.Timeout = time.Duration(5) * time.Second
		}

		res, err := requests.Get(url, timeout)

		if err != nil {
			panic(err)
		}

		defer wg.Done()

		// Helper method
		htmlStr := res.String()
		// fmt.Println(htmlStr)

		printResponse(htmlStr)
	}()

	wg.Wait()
}

func printNum() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(10 * time.Millisecond)
	}
}

func printResponse(jsonStr string) {
	// fmt.Println(jsonStr)

}
