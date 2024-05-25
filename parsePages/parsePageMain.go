package parsePages

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

type Product struct {
	Id                                                                                                int
	Category, ProductUrl, ProductName, MiniImgUrlProduct, ImgUrlProduct, CategoryUrl, OldPrice, Price string
}

func InitMain(URL string, categories []string, address *string, products *[]Product) {
	c := colly.NewCollector(colly.AllowURLRevisit(),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"))

	c.OnHTML("div.js-p-m-m__item.p-m-m__item", func(e *colly.HTMLElement) {
		categoryUrl := e.ChildAttr("a.p-m-m__item-link", "href")
		category := e.DOM.Find("a.p-m-m__item-link").Text()
		category = category[1 : len(category)-1]
		categoryId := 0

		if categoryUrl != "" && category != "" {
			if categories[0] != category && categories[1] != category && categories[2] != category {
				return
			}

			switch category {
			case categories[0]:
				categoryId = 0
			case categories[1]:
				categoryId = 1
			case categories[2]:
				categoryId = 2
			}

			product := Product{
				Id:          categoryId,
				CategoryUrl: URL + categoryUrl,
				Category:    category,
			}
			println(product.CategoryUrl)
			*products = append(*products, product)
		}
	})

	c.OnHTML("div#cur-city-info", func(e *colly.HTMLElement) {
		element := e.DOM.Find("div.left").Eq(1)
		*address = element.Text()
		println(address)
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "*/*")
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("Got: ", r.Request.Headers)
		log.Println("Got: ", r.Request.Body)
		log.Println("Got: ", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, e error) {
		log.Println("Error: ", r.Request.URL, e, r.Headers)
		fmt.Println("Got this error:", e, r.StatusCode)
	})

	c.Visit(URL)
}
