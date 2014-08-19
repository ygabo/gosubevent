package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

var (
	JSON_FILE       = "data/subevent.json"
	eventsByChapter []Chapter
)

type Subevent struct {
	Id       string `json:"id"`
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
	Period    string      `json:"period"`
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

	err = json.Unmarshal(dat, &eventsByChapter)
	check(err)
}

func getChapter(chapter int) Chapter {
	for _, current := range eventsByChapter {
		if current.Chapter == chapter {
			return current
		}
	}
	return Chapter{}
}

func indexHandler(r render.Render) {
	r.HTML(200, "index", nil)
}
func ttHandler(r render.Render) {
	r.HTML(200, "tt", nil)
}
func main() {
	updateCache()
	m := martini.Classic()

	templateOptions := render.Options{}
	templateOptions.Delims.Left = "#{"
	templateOptions.Delims.Right = "}#"
	m.Use(render.Renderer(templateOptions))

	m.Get("/", indexHandler)
	m.Get("/tt", ttHandler)

	m.Get("/chapter", func() string {
		kk, _ := json.Marshal(eventsByChapter)
		return string(kk)
	})
	m.Get("/chapter/:id", func(params martini.Params) string {
		var chap Chapter
		if i, err := strconv.Atoi(params["id"]); err == nil {
			chap = getChapter(i)
		} else {
			chap = Chapter{}
		}
		kk, _ := json.Marshal(chap)
		return string(kk)
	})
	m.Get("/subevent", func() string {
		kk, _ := json.Marshal(eventsByChapter)
		return string(kk)
	})
	m.Get("/reload", func(r render.Render) {
		updateCache()
		r.HTML(200, "index", nil)
	})
	m.Run()
}
