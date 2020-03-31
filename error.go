package xs

import (
	"fmt"
	"reflect"
)

type SheetAndSLiceMismatched struct {
	SheetNum int
	SliceNum int
}

func (e *SheetAndSLiceMismatched) Error() string {
	return fmt.Sprintf("xs: xlsx contain %d sheet. and get %d slice", e.SheetNum, e.SliceNum)
}

// the parameter s is in a invalid type
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "xs: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return fmt.Sprintf("xs: Unmarshal(non-pointer %s)", e.Type.String())
	}

	return "xs: Unmarshal(non-slice pointer)"
}

// the parameter s is in a invalid type
type InvalidMarshalError struct {
	Type reflect.Type
}

func (e *InvalidMarshalError) Error() string {
	if e.Type == nil {
		return "xs: Marshal(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return fmt.Sprintf("xs: Marshal(non-pointer %s)", e.Type.String())
	}

	return "xs: Marshal(non-slice pointer)"
}

//xlsx file can't be opened without error
type OpenFileError struct {
	FilePath string
	Err      error
}

func (e *OpenFileError) Error() string {
	return fmt.Sprintf("parse_ad_agent_excel: failed to open file %s: %s", e.FilePath, e.Err.Error())
}

//In some rows, the value of some required fields is null
type LackColError struct {
	Row    int
	Header string
}

func (e *LackColError) Error() string {
	return fmt.Sprintf("%d row %s col is nil", e.Row, e.Header)
}

//the sheet need to contains more than one row
type EmptySheetError struct {
}

func (e *EmptySheetError) Error() string {
	return fmt.Sprintf("empty sheet")
}

//lack of required header
type LackHeaderError struct {
	Header string
}

func (e *LackHeaderError) Error() string {
	return fmt.Sprintf("check headers failed. missing header: %s", e.Header)
}

//the type of data mismatched the required type
type TypeMismatchedError struct {
	Data         string
	RequiredType reflect.Type
	GivenType    reflect.Type
}

func (e *TypeMismatchedError) Error() string {
	return fmt.Sprintf("type mismatched。the data is：%s. the required type is: %s. the given type is： %s", e.Data, e.RequiredType, e.GivenType)
}
