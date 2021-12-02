package ghtrend

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var baseurl = "https://github.com/trending"

type Repo struct {
	Url         string
	Name        string
	Description string
	Language    string
}

type Repos []Repo

func Scraping(language string) (Repos, error) {

	doc, err := goquery.NewDocument(baseurl + "/" + language)

	if err != nil {
		return nil, err
	}

	var res Repos

	doc.Find("article").Each(func(_ int, s *goquery.Selection) {

		t, _ := s.Find("h1 a").First().Attr("href")

		u := "https://github.com" + t
		n := strings.Replace(t, "/", "", 1)

		des := strings.TrimSpace(s.Find("p").First().Text())

		la := strings.TrimSpace(s.Find("div.f6 span").First().Text())
		if la == "Built by" {
			la = ""
		}
		res = append(res, Repo{u, n, des, la})

	})

	return res, nil
}
