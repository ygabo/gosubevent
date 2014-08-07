package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/go-martini/martini"
)

var (
	JSON_FILE       = "subevent.json"
	eventsByChapter []Chapter
)

type Subevent struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Missable bool   `json:"missable"`
	Location string `json:"location"`
	Period   string `json:"period"`
	Info     string `json:"info"`
	Reward   string `json:"reward"`
}

type Chapter struct {
	Chapter   int         `json:"chapter"`
	Title     string      `json:"title"`
	Location  string      `json:"location"`
	Subevents *[]Subevent `json:"subevents"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func updateJsonCache() {
	dat, err := ioutil.ReadFile(JSON_FILE)
	check(err)

	err = json.Unmarshal(dat, &eventsByChapter)
	check(err)
}

func main() {
	updateJsonCache()
	m := martini.Classic()
	m.Get("/", func() string {
		kk, _ := json.Marshal(eventsByChapter)
		return string(kk)
	})
	m.Run()
}
