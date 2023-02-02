package main

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/mmcdole/gofeed"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Show struct {
	Title    string `json:"name,omitempty"`
	Magnet   string `json:"magnet,omitempty"`
	Seeds    string `json:"seeds,omitempty"`
	Peers    string `json:"peers,omitempty"`
	Verified string `json:"verified,omitempty"`
	Pubdate  string `json:"pubdate,omitempty"`
	FileName string `json:"filename,omitempty"`
}

func fromItem(item *gofeed.Item) Show {
	s := Show{}

	s.Title = item.Title
	s.Pubdate = item.Published

	for _, ext := range item.Extensions {
		s.FileName = ext["fileName"][0].Value
		s.Magnet = ext["magnetURI"][0].Value
		s.Seeds = ext["seeds"][0].Value
		s.Peers = ext["peers"][0].Value
		s.Verified = ext["verified"][0].Value

	}
	return s
}

func parseFeed(feed *gofeed.Feed) []Show {
	s := make([]Show, 0)
	for _, link := range feed.Items {
		show := fromItem(link)
		s = append(s, show)
	}
	return s
}

func showsToJSON(shows []Show) string {
	b, err := json.Marshal(shows)
	check(err)
	return string(b)
}

func main() {
	fp := gofeed.NewParser()
	rss, e := http.Get("https://eztv.ag/ezrss.xml")
	check(e)
	feed, e := fp.Parse(rss.Body)
	check(e)

	shows := parseFeed(feed)

	fmt.Println(showsToJSON(shows))
}
