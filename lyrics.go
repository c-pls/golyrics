package golyrics

import (
	"log"
	"strings"
)

import (
	"github.com/gocolly/colly/v2"
)

var countryCode = map[string]string{
	"en": "english",
	"vi": "Vietnamese",
	"jp": "japanese-romaji",
	"es": "Spanish",
}

var baseURL = "https://www.musixmatch.com/lyrics/"

func GetLyrics(song, artist string) string {
	var res string
	url := baseURL + formatURL(artist) + "/" + formatURL(song)

	c := colly.NewCollector()
	c.OnHTML(".mxm-lyrics__content", func(e *colly.HTMLElement) {
		res += e.Text
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	return res
}

func GetLyricsWithTranslate(song, artist, codeLang string) map[string]string {
	m := make(map[string]string)
	url := baseURL + formatURL(artist) + "/" + formatURL(song) + "/translation" + "/" + countryCode[codeLang]
	c := colly.NewCollector()
	var translation string

	c.OnHTML(".mxm-translatable-line-readonly", func(e *colly.HTMLElement) {
		sel := e.DOM
		translation += strings.TrimSpace(sel.Find("div.row div.col-xs-6.col-sm-6.col-md-6.col-ml-6.col-lg-6 div div").Text()) + "\n"
	})
	err := c.Visit(url)
	if err != nil {
		return nil
	}
	m["original"] = GetLyrics(song, artist)
	m["translation"] = translation

	return m
}

func formatURL(s string) string {
	return strings.Replace(s, " ", "-", -1)
}
