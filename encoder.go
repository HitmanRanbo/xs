package xs

import (
	"database/sql"
	"fmt"
	"github.com/tealeg/xlsx"
	"reflect"
	"time"
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
		sheetRow.Cells = make([]*xlsx.Cell, len(tagInfo.Headers))

		row := sValues.Index(line)
		for index, v := range tagInfo.Headers {
			cell := &xlsx.Cell{}
			data := row.Field(tagInfo.M[v].Index)
			switch t := data.Interface().(type) {
			case time.Time:
				cell.SetValue(t)
			case fmt.Stringer: // check Stringer first
				cell.SetString(t.String())
			case sql.NullString: // check null sql types nulls = ''
				if cell.SetString(``); t.Valid {
					cell.SetValue(t.String)
				}
			case sql.NullBool:
				if cell.SetString(``); t.Valid {
					cell.SetBool(t.Bool)
				}
			case sql.NullInt64:
				if cell.SetString(``); t.Valid {
					cell.SetValue(t.Int64)
				}
			case sql.NullFloat64:
				if cell.SetString(``); t.Valid {
					cell.SetValue(t.Float64)
				}
			default:
				switch data.Kind() {
				case reflect.String, reflect.Int, reflect.Int8,
					reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float64, reflect.Float32:
					cell.SetValue(data.Interface())
					if tagInfo.M[v].IsHyperlink && data.Interface().(string) != "" {
						cell.SetStringFormula(fmt.Sprintf(`=HYPERLINK("%s", "%s")`, data.Interface().(string), data.Interface().(string)))
						style := cell.GetStyle()
						style.Font.Underline = true   //加下划线
						style.Font.Color = "FF0000FF" //设置字体颜色为蓝色
						cell.SetStyle(style)
					}
				case reflect.Bool:
					cell.SetBool(t.(bool))
				default:
					cell.SetString(fmt.Sprintf("%v", data))
				}
			}
			if !tagInfo.M[v].IsHyperlink && tagInfo.M[v].Format != "" {
				cell.SetFormat(tagInfo.M[v].Format)
			}
			sheetRow.Cells[index] = cell
		}
	}
}
