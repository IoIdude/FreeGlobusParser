package parsePages

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"slices"
)

func InitProducts(categoryId int, products *[]Product, stockURL string, categories []string) {
	productId := slices.IndexFunc(*products, func(c Product) bool { return c.Id == categoryId })
	categoryUrl := (*products)[productId].CategoryUrl

	c := colly.NewCollector(colly.AllowURLRevisit(),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"))

	c.OnHTML("div.pim-list__items.js-pim-list__items", func(e *colly.HTMLElement) {
		var productUrl string

		productHrefEl := e.ChildAttrs("a", "href")
		for index, data := range productHrefEl {
			if index < 5 {
				productUrl = stockURL + data
				productId := slices.IndexFunc(*products, func(c Product) bool { return c.Id == categoryId && c.ProductUrl == "" })

				if productId != -1 {
					(*products)[productId].ProductUrl = productUrl
					(*products)[productId] = ParseProduct((*products)[productId], stockURL)
				} else {
					product := Product{
						Id:          categoryId,
						ProductUrl:  productUrl,
						CategoryUrl: categoryUrl,
						Category:    categories[categoryId],
					}

					*products = append(*products, ParseProduct(product, stockURL))
				}
			}
		}

		miniImgUrlEl := e.ChildAttrs("img.pim-list__item-img", "data-src")
		for index, data := range miniImgUrlEl {
			if index < 5 {
				miniImgUrlProduct := stockURL + data
				productId := slices.IndexFunc(*products, func(c Product) bool { return c.Id == categoryId && c.ProductUrl != "" && c.MiniImgUrlProduct == "" })
				(*products)[productId].MiniImgUrlProduct = miniImgUrlProduct
			}
		}

		productNameEl := e.ChildAttrs("div.pim-list__item-title.js-crop-text", "data-full-text")
		for index, data := range productNameEl {
			if index < 5 {
				productName := data
				productId := slices.IndexFunc(*products, func(c Product) bool {
					return c.Id == categoryId && c.ProductUrl != "" && c.MiniImgUrlProduct != "" && c.ProductName == ""
				})
				(*products)[productId].ProductName = productName
			}
		}
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

	c.Visit(categoryUrl)
}
