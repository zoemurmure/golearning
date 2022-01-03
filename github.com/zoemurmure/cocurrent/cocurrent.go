package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string, 10)

	apis := []string{
		"https://www.baidu.com",
		"https://www.163.com",
		"https://www.huorong.cn/",
		"https://im.qq.com/",
		"https://api.somewhereintheinternet.com/",
		"http://iqiyi.com/",
	}

	for _, api := range apis {
		go checkAPI(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

func checkAPI(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}
	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}
