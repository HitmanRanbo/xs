package xs

import (
	"reflect"
	"strings"
)

const (
	XS_TAG                = "xs"
	NULL_TAG              = ""
	IGNORE_TAG            = "-"
	OMITEMPTY             = "omitempty"
	TAG_SEPERATER         = ","
	TAG_FORMATE_SEPERATER = ":"
	HYPERLINK_FORMATE     = "hyperlink"
)

func getSTypes(s interface{}) reflect.Type {
	//处理指针和非结构体这两种情况
	if reflect.ValueOf(s).Kind() == reflect.Ptr {
		return reflect.TypeOf(s).Elem().Elem()
	} else {
		return reflect.TypeOf(s).Elem()
	}
}

type TagInfo struct {
	Index       int
	Format      string
	IsHyperlink bool
	Omitempty   bool
}

type TagInfoMap struct {
	Headers []string
	M       map[string]TagInfo
}

func getHeaderAndFormat(headerWithFormat string) (string, string, bool) {
	data := strings.Split(headerWithFormat, TAG_FORMATE_SEPERATER)
	if len(data) == 1 {
		return data[0], "", false
	}
	if data[1] == HYPERLINK_FORMATE {
		return data[0], "", true
	}
	return data[0], data[1], false
}

func GetTagInfo(s interface{}) TagInfoMap {
	var sTypes = getSTypes(s)
	var m = make(map[string]TagInfo)
	var headers = make([]string, 0)

	for i := 0; i < sTypes.NumField(); i++ {
		tag := strings.Split(sTypes.Field(i).Tag.Get(XS_TAG), TAG_SEPERATER)
		if len(tag) == 1 && (tag[0] == NULL_TAG || tag[0] == IGNORE_TAG) {
			continue
		}
		header, format, isHyperLink := getHeaderAndFormat(tag[0])
		if len(tag) > 1 && tag[1] == OMITEMPTY {
			headers = append(headers, header)
			m[header] = TagInfo{Index: i, Format: format, IsHyperlink: isHyperLink, Omitempty: true}
		} else {
			headers = append(headers, header)
			m[header] = TagInfo{Index: i, Format: format, IsHyperlink: isHyperLink, Omitempty: false}
		}
	}
	return TagInfoMap{Headers: headers, M: m}
}

//parse tags. return map[string]bool, which key is tag and value is if the tag is necessary
func GetTags(s interface{}) map[string]bool {
	var headers = make(map[string]bool)
	var tagInfo = GetTagInfo(s)

	for _, header := range tagInfo.Headers {
		headers[header] = !tagInfo.M[header].Omitempty
	}

	return headers
}
