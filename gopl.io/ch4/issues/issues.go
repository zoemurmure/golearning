package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	var timeResult map[string][]*github.Issue = make(map[string][]*github.Issue)
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range result.Items {
		days := time.Since(item.CreatedAt).Hours() / 24
		switch {
		case days >= 365:
			timeResult["more than 1 year"] = append(timeResult["more than 1 year"], item)
		case days >= 30:
			timeResult["more than 1 month"] = append(timeResult["more than 1 month"], item)
		default:
			timeResult["less than 1 month"] = append(timeResult["less than 1 month"], item)
		}
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for k, v := range timeResult {
		fmt.Println(k)
		for _, item := range v {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}
