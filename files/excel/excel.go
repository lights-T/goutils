package excel

import (
	"github.com/xuri/excelize/v2"
)

func ReadeExcel(address string) ([]string, []map[string]string, error) {
	cols := make([]string, 0, 100)
	ret := make([]map[string]string, 0, 100)
	f, err := excelize.OpenFile(address)
	if err != nil {
		return cols, ret, err
	}
	defer f.Close()
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return cols, ret, err
	}
	for i, row := range rows {
		if i == 0 { //取得第一行的所有数据---execel表头
			for _, colCell := range row {
				cols = append(cols, colCell)
			}
		} else {
			theRow := map[string]string{}
			for j, colCell := range row {
				k := cols[j]
				theRow[k] = colCell
			}
			ret = append(ret, theRow)
		}
	}
	return cols, ret, nil
}
