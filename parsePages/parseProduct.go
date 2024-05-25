package parsePages

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
)

func ParseProduct(product Product, stockUrl string) Product {
	c := colly.NewCollector(colly.AllowURLRevisit(),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"))

	c.OnHTML("div.js-content-and-footer.content-and-footer", func(e *colly.HTMLElement) {
		imgUrl := e.ChildAttr("img.js-catalog-detail__header-image-img.catalog-detail__header-image-img", "src")
		product.ImgUrlProduct = stockUrl + imgUrl
	})

	c.OnHTML("div.catalog-detail__header-prices", func(e *colly.HTMLElement) {
		imgUrl := e.DOM.Find("div.catalog-detail__item-price-actual.catalog-detail__item-price-actual--discount-color")

		priceMainEl := e.DOM.Find("span.catalog-detail__item-price-actual-main")
		priceSubEl := e.DOM.Find("span.catalog-detail__item-price-actual-sub")

		product.Price = strings.ReplaceAll(priceMainEl.Text()+"."+priceSubEl.Text()+"руб.", " ", "")

		if imgUrl.Text() != "" {
			priceOldMainEl := e.DOM.Find("span.catalog-detail__item-price-old-main")
			priceOldSubEl := e.DOM.Find("span.catalog-detail__item-price-old-sub")

			product.OldPrice = strings.ReplaceAll(priceOldMainEl.Text()+"."+priceOldSubEl.Text()+"руб.", " ", "")
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

	c.Visit(product.ProductUrl)

	return product
}
