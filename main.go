package main

import (
	"parserProj/parsePages"
	"parserProj/utils"
)

func main() {
	var address string
	var products []parsePages.Product
	URL := "https://www.globus.ru"

	categories := []string{"Бытовая техника, электроника", "Зоотовары", "Чай, кофе, какао"}
	parsePages.InitMain(URL, categories, &address, &products)

	for i := range categories {
		parsePages.InitProducts(i, &products, URL, categories)
	}

	utils.Export(products, address)
}
