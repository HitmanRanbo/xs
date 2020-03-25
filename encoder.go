package xs

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"reflect"
)

//encode write a xlsx.sheet from slice of struct
func encode(sheet *xlsx.Sheet, tagInfo TagInfoMap, sValues reflect.Value) {
	//write header
	row := sheet.AddRow()
	for _, tag := range tagInfo.Headers {
		row.AddCell().Value = tag
	}

	//loop slice and write xlsx.sheet row by row
	for line := 0; line < sValues.Len(); line++ {
		//init row
		sheetRow := sheet.AddRow()
		sheetRow.Cells = make([]*xlsx.Cell, len(tagInfo.Headers), len(tagInfo.Headers))

		row := sValues.Index(line)
		for _, v := range tagInfo.Headers {
			cell := &xlsx.Cell{}
			data := row.Field(tagInfo.M[v].Index)
			switch data.Kind() {
			case reflect.Int:
				cell.SetInt(data.Interface().(int))
			case reflect.Int64:
				cell.SetInt64(data.Interface().(int64))
			case reflect.String:
				cell.SetString(data.Interface().(string))
				if tagInfo.M[v].IsHyperlink {
					cell.SetFormula(fmt.Sprintf(`HYPERLINK("%s","%s")`, data.Interface().(string), data.Interface().(string)))
				}
			case reflect.Float64:
				cell.SetFloat(data.Interface().(float64))
			default:
				cell.SetString(fmt.Sprintf("%v", data))
			}
			if tagInfo.M[v].Format != "" {
				cell.SetFormat(tagInfo.M[v].Format)
			}
			sheetRow.Cells[tagInfo.M[v].Index] = cell
		}
	}

}
