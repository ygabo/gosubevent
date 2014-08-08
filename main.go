package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/go-martini/martini"
)

var (
	JSON_FILE   = "data/subevent.json"
	blankEvents []Chapter
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

func updateCache() {
	dat, err := ioutil.ReadFile(JSON_FILE)
	check(err)

	err = json.Unmarshal(dat, &blankEvents)
	check(err)
}

func main() {
	updateCache()
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Get("/events.json", func() string {
		kk, _ := json.Marshal(blankEvents)
		return string(kk)
	})
	m.Run()
}
