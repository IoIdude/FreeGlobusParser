package utils

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
	"log"
	"parserProj/parsePages"
)

func Export(products []parsePages.Product, address string) {
	f := excelize.NewFile()
	sheetName := "products"
	f.SetSheetName("Sheet1", sheetName)

	headers := []string{
		"Категория товара",
		"Url товара ",
		"Наименование товара",
		"Url изображение товара (маленькая)",
		"Url изображение товара",
		"Текущая цена",
		"Старая цена",
	}
	for i, header := range headers {
		column := string('A' + i)
		cell := fmt.Sprintf("%s1", column)

		f.SetColWidth(sheetName, column, column, 30.00)
		f.SetCellValue(sheetName, cell, header)
	}

	for i, product := range products {
		row := i + 2
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), product.Category)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), product.ProductUrl)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), product.ProductName)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), product.MiniImgUrlProduct)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), product.ImgUrlProduct)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), product.Price)
		if product.OldPrice == "" {
			product.OldPrice = "-"
		}
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), product.OldPrice)
	}
	f.SetCellValue(sheetName, fmt.Sprintf("A%d", len(products)+3), "Адрес - "+address)

	fileName := fmt.Sprintf("./exportedFiles/products_%s.xlsx", uuid.New().String())
	if err := f.SaveAs(fileName); err != nil {
		log.Fatalf("Ошибка при сохранении файла: %v", err)
	}
}
