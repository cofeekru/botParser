package parser

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/gocolly/colly"
)

const (
	hostBigGeek = "https://biggeek.ru"
)

type unit struct {
	Name  string `json:"name"`
	Price string `json:"price"`
	URL   string `json:"url"`
}

func parser(url *url.URL) ([]byte, int) {
	collector := colly.NewCollector()
	collector.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	var items []unit
	collector.OnHTML("div.catalog-card", func(html *colly.HTMLElement) {
		item := unit{
			Name:  html.ChildText("a.catalog-card__title"),
			Price: html.ChildText("b.cart-modal-count"),
			URL:   hostBigGeek + html.ChildAttr("a.catalog-card__title", "href"),
		}
		items = append(items, item)
	})

	collector.OnHTML("a.prod-pagination__item-next", func(html *colly.HTMLElement) {
		request := html.Request.AbsoluteURL(html.Attr("href"))
		collector.Visit(request)

	})

	err := collector.Visit(hostBigGeek + url.String())
	if err != nil {
		log.Println("Request is error.")
		return nil, http.StatusNotFound
	}

	json, err := json.Marshal(items)

	if err != nil {
		log.Println("JSON converting is error.")
		return nil, http.StatusInternalServerError
	}

	return json, http.StatusOK
}
