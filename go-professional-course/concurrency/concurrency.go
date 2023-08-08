package concurrency

import (
	"fmt"
	"net/http"
	"time"
)

func Conc() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
	}

	canal := make(chan string)

	for _, api := range apis {
		go checkApi(api, canal)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Println(<-canal)
	}

	time.Sleep(1 * time.Second)

	elapsed := time.Since(start)

	fmt.Println("Elapsed time: ", elapsed)
}

func checkApi(url string, ch chan string) {
	if _, err := http.Get(url); err != nil {
		ch <- fmt.Sprintf("ERROR: %s There was a problem with the API\n", url)
		return
	}
	ch <- fmt.Sprintf("SUCCESS: %s is working normally\n", url)
}
