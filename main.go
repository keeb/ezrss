package main

import (
	"fmt"
	"os"

	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	f, _ := os.Open("test/ezrss.xml")
	defer f.Close()

	feed, _ := fp.Parse(f)
	fmt.Println(feed.Updated)

	for _, link := range feed.Items {
		fmt.Println(link.Title)
	}
}