package xs

import (
	"bytes"
	"github.com/tealeg/xlsx"
	"reflect"
	"strconv"
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
	if len(xlFile.Sheets) != len(ss) {
		return &SheetAndSLiceMismatched{len(xlFile.Sheets), len(ss)}
	}
	//空文件不处理
	//do not process empty file
	if len(xlFile.Sheets) == 0 {
		return nil
	}

	for sheetIndex, sheet := range xlFile.Sheets {
		var s = ss[sheetIndex]
		var tagInfo = GetTagInfo(s)
		var sValues = reflect.ValueOf(s)

		//s应该是一个array或者slice的指针
		//s should be array or slice of ptr
		if sValues.Kind() != reflect.Ptr || sValues.IsNil() {
			return &InvalidUnmarshalError{Type: reflect.TypeOf(s)}
		}
		if sValues.Type().Elem().Kind() != reflect.Slice {
			return &InvalidUnmarshalError{Type: sValues.Type()}
		}

		//逐行读xlsx文件，并转化成结构体
		//preprocess excel
		mList := make([]map[string]*xlsx.Cell, sheet.MaxRow-1, sheet.MaxRow-1)
		for i, row := range sheet.Rows {
			if i == 0 {
				continue
			}

			m := make(map[string]*xlsx.Cell, len(tagInfo.Headers))
			for _, tag := range tagInfo.Headers {
				//非空字段不可为空
				if !tagInfo.M[tag].Omitempty && row.Cells[tagInfo.M[tag].Index].String() == "" {
					return &LackColError{i, tag}
				}
				m[tag] = row.Cells[tagInfo.M[tag].Index]
			}
			mList[i-1] = m
		}
		err := decode(mList, tagInfo, s)
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
