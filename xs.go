package xs

import (
	"bytes"
	"github.com/tealeg/xlsx"
	"reflect"
	"strconv"
	"strings"
)

func UnmarshalFromFile(filePath string, ss ...interface{}) error {
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		return &OpenFileError{filePath, err}
	}

	return unmarshal(xlFile, ss...)
}

func Unmarshal(body []byte, ss ...interface{}) error {
	xlFile, err := xlsx.OpenBinary(body)
	if err != nil {
		return &OpenFileError{"body", err}
	}

	return unmarshal(xlFile, ss...)
}

func Marshal(ss ...interface{}) ([]byte, error) {
	var buf bytes.Buffer

	xlFile, err := marshal(ss...)
	if err != nil {
		return nil, err
	}
	err = xlFile.Write(&buf)
	return buf.Bytes(), err
}

func unmarshal(xlFile *xlsx.File, ss ...interface{}) error {
	//检查sheet数和结构体的数目
	//check if the num of sheet equals to the num of slice
	if len(xlFile.Sheets) < len(ss) {
		return &SheetAndSLiceMismatched{len(xlFile.Sheets), len(ss)}
	}
	//空文件不处理
	//do not process empty file
	if len(xlFile.Sheets) == 0 {
		return nil
	}

	for sheetIndex, s := range ss {
		sheet := xlFile.Sheets[sheetIndex]
		tagInfo := GetTagInfo(s)
		headerIndexMap, err := genHeaderIndexMap(tagInfo, sheet.Rows[0])
		if err != nil {
			return err
		}

		//s应该是一个array或者slice的指针
		//s should be array or slice of ptr
		sValues := reflect.ValueOf(s)
		if sValues.Kind() != reflect.Ptr || sValues.IsNil() {
			return &InvalidUnmarshalError{Type: reflect.TypeOf(s)}
		}
		if sValues.Type().Elem().Kind() != reflect.Slice {
			return &InvalidUnmarshalError{Type: sValues.Type()}
		}

		//逐行读xlsx文件，并转化成结构体
		//preprocess excel
		mList := make([]map[string]*xlsx.Cell, 0)
		for i, row := range sheet.Rows {
			if i == 0 {
				continue
			}

			//the first empty row is the last row of the sheet
			var lastRow bool

			m := make(map[string]*xlsx.Cell, len(tagInfo.Headers))
			for _, tag := range tagInfo.Headers {

				//check if NonOmitempty tag could get the value
				if !tagInfo.M[tag].Omitempty && row.Cells[headerIndexMap[tag]].String() == "" {
					for _, cell := range row.Cells {
						//not empty row
						if cell.String() != "" {
							return &LackColError{i, tag}
						}
					}
					//empty row
					lastRow = true
					break
				} else {
					m[tag] = row.Cells[headerIndexMap[tag]]
				}
			}
			if lastRow {
				break
			}

			mList = append(mList, m)
		}
		err = decode(mList, tagInfo, s)
		if err != nil {
			return err
		}
	}

	return nil
}

func marshal(ss ...interface{}) (*xlsx.File, error) {
	excel := xlsx.NewFile()

	for index, s := range ss {
		var sheet, _ = excel.AddSheet("sheet" + strconv.Itoa(index))
		var sValues = reflect.ValueOf(s)
		var tagInfo = GetTagInfo(s)

		if sValues.Type().Kind() != reflect.Slice {
			return nil, &InvalidMarshalError{Type: sValues.Type()}
		}

		encode(sheet, tagInfo, sValues)
	}

	return excel, nil
}

func genHeaderIndexMap(tagInfoMap TagInfoMap, xlsxHeader *xlsx.Row) (map[string]int, error) {
	var headerIndexMap = make(map[string]int)

	for header, tagInfo := range tagInfoMap.M {
		exist := true
		if !tagInfo.Omitempty {
			exist = false
		}
		for i, cell := range xlsxHeader.Cells {
			if strings.TrimSpace(cell.String()) == header {
				headerIndexMap[header] = i
				exist = true
			}
		}

		if !exist {
			return headerIndexMap, &LackHeaderError{header}
		}
	}

	return headerIndexMap, nil
}
