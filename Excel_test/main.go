package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func main() {
	// Create a new Excel file
	file := xlsx.NewFile()

	// Add a new sheet to the Excel file
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a new row and add cells to it
	row := sheet.AddRow()
	cell1 := row.AddCell()
	cell2 := row.AddCell()
	cell3 := row.AddCell()

	// Set values for the cells
	cell1.SetString("Name")
	cell2.SetString("Age")
	cell3.SetString("Country")

	// Add data rows
	data := [][]string{
		{"John", "30", "USA"},
		{"Alice", "25", "Canada"},
		{"Bob", "40", "UK"},
	}

	for _, rowData := range data {
		row = sheet.AddRow()
		for _, value := range rowData {
			cell := row.AddCell()
			cell.SetString(value)
		}
	}

	// Save the Excel file
	err = file.Save("example.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Excel file 'example.xlsx' created successfully.")
}
