package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Data struct {
	Date   string
	Height string
	Figure int
}

func main() {

	f, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	datas := make([]Data, 0, 50)
	// Get value from cell by given worksheet name and axis.
	// cell, err := f.GetCellValue("Sheet1", "B2")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	for ri, row := range rows {
		var date string
		if ri == 0 {
			continue
		}
		for ci, colCell := range row {
			if ci == 0 {
				date = colCell
				continue
			}
			s := strings.Split(colCell, "/")
			num, _ := strconv.Atoi(s[2])
			data := Data{timeFormat(date) + s[0], s[1], num}
			datas = append(datas, data)
			fmt.Print(data, "\n")
		}

	}

	createExcel(datas)
}

func createExcel(datas []Data) {
	f := excelize.NewFile()
	categories := map[string]string{"A1": "날짜", "B1": "조위", "C1": "고/저"}
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	for i, data := range datas {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa((i+2)), data.Date)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa((i+2)), data.Figure)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa((i+2)), data.Height)
	}
	// Save xlsx file by the given path.
	if err := f.SaveAs("formated.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func timeFormat(s1 string) string {
	d := strings.Split(s1, "-")
	year := "20" + d[2]
	month := d[0]
	day := d[1]
	return year + "-" + month + "-" + day
}
