package xs

import (
	"github.com/tealeg/xlsx"
	"reflect"
	"strings"
)

func decode(mList []map[string]*xlsx.Cell, tagInfo TagInfoMap, s interface{}) error {
	var sValue = reflect.ValueOf(s).Elem()

	slice := reflect.MakeSlice(sValue.Type(), 0, len(mList))
	for _, m := range mList {
		//根据结构体对应字段的类型对该字段赋值，如果类型不匹配会报错
		elem := reflect.New(sValue.Type().Elem())
		for k, v := range m {
			switch elem.Elem().Field(tagInfo.M[k].Index).Kind() {
			case reflect.Bool:
				value := v.Bool()
				elem.Elem().Field(tagInfo.M[k].Index).Set(reflect.ValueOf(value))
				break
			case reflect.Int:
				value, err := v.Int()
				if err != nil {
					return err
				}
				elem.Elem().Field(tagInfo.M[k].Index).Set(reflect.ValueOf(value))
				break
			case reflect.Int64:
				value, err := v.Int64()
				if err != nil {
					return err
				}
				elem.Elem().Field(tagInfo.M[k].Index).Set(reflect.ValueOf(value))
				break
			case reflect.Float64:
				value, err := v.Float()
				if err != nil {
					return err
				}
				elem.Elem().Field(tagInfo.M[k].Index).Set(reflect.ValueOf(value))
				break
			default:
				value := strings.TrimRight(v.String(), " ")
				elem.Elem().Field(tagInfo.M[k].Index).Set(reflect.ValueOf(value))
			}
		}
		slice = reflect.Append(slice, reflect.Indirect(elem))
	}

	sValue.Set(slice.Slice3(0, len(mList), len(mList)))
	return nil
}
