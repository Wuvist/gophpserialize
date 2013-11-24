package gophpserialize

import (
	"encoding/json"
	"fmt"
	"strings"
)

type MarshalError struct {
	Reason string
}

func (this *MarshalError) Error() string {
	return this.Reason
}

func marshalSimpleType(value interface{}, isFromJson bool) (result string, err error) {
	switch value.(type) {
	case nil:
		return "N;", nil
	case int:
		return fmt.Sprintf("i:%v;", value), nil
	case string:
		return fmt.Sprintf(`s:%d:"%s";`, len(value.(string)), value.(string)), nil
	case float32, float64:
		// for all the interface{} passed to json.Unmarshal, numbers are treat as float no matter what it is.
		if isFromJson {
			svalue := fmt.Sprintf("%v", value)
			if strings.Contains(svalue, ".") {
				return fmt.Sprintf("d:%v;", value), nil
			} else {
				return fmt.Sprintf("i:%v;", value), nil
			}
		} else {
			return fmt.Sprintf("d:%v;", value), nil
		}
	case bool:
		if value.(bool) {
			return "b:1;", nil
		} else {
			return "b:0;", nil
		}
	default:
		return "", &MarshalError{fmt.Sprintf("Not a simple type: %v", value)}
	}
}

func marshalComplexInner(index, value interface{}, isFromJson bool) (result string, err error) {
	switch value.(type) {
	case []interface{}, map[string]interface{}:
		indexvalue, err := marshalSimpleType(index, isFromJson)
		if err != nil {
			return "", err
		}
		subvalue, err1 := marshalComplex(value, isFromJson)
		if err1 != nil {
			return "", err1
		}
		return fmt.Sprintf("%s%s", indexvalue, subvalue), nil
	default:
		indexvalue, err := marshalSimpleType(index, isFromJson)
		if err != nil {
			return "", err
		}
		subvalue, err1 := marshalSimpleType(value, isFromJson)
		if err1 != nil {
			return "", err1
		}
		return fmt.Sprintf("%s%s", indexvalue, subvalue), nil
	}
}

func marshalComplex(value interface{}, isFromJson bool) (result string, err error) {
	/* 
		Here because the dark side of go language. Such as len() and range cannot take a interface{} as parameter.
		I have to repeat the for statement and range statement.
	*/
	switch value.(type) {
	case []interface{}:
		typeList := []string{fmt.Sprintf("a:%d:{", len(value.([]interface{})))} // I really hate this
		for index, value := range value.([]interface{}) {
			subvalue, err1 := marshalComplexInner(index, value, isFromJson)
			if err1 != nil {
				return "", nil
			}
			typeList = append(typeList, subvalue)
		}
		typeList = append(typeList, "}")
		return strings.Join(typeList, ""), nil
	case map[string]interface{}:
		typeList := []string{fmt.Sprintf("a:%d:{", len(value.(map[string]interface{})))}
		for index, value := range value.(map[string]interface{}) {
			subvalue, err1 := marshalComplexInner(index, value, isFromJson)
			if err1 != nil {
				return "", nil
			}
			typeList = append(typeList, subvalue)
		}
		typeList = append(typeList, "}")
		return strings.Join(typeList, ""), nil
	default:
		return "", &MarshalError{fmt.Sprintf("Not a complex type: %v", value)}
	}
}

func marshal(jsonData interface{}, isFromJson bool) (phpData []byte, err error) {

	switch jsonData.(type) {
	case []interface{}, map[string]interface{}:
		subvalue, err1 := marshalComplex(jsonData, isFromJson)
		if err1 != nil {
			return []byte(``), err1
		}
		return []byte(subvalue), nil
	default:
		subvalue, err := marshalSimpleType(jsonData, isFromJson)
		if err != nil {
			return []byte(``), err
		}
		return []byte(subvalue), nil
	}
}

func MarshalNormal(data interface{}) (result []byte, err error) {
	return marshal(data, false)
}

func MarshalJson(data interface{}) (result []byte, err error) {
	return marshal(data, true)
}

func JsonToPhp(jsonRaw []byte) (phpData []byte, err error) {
	var jsonData interface{}
	err = json.Unmarshal(jsonRaw, &jsonData)
	if err != nil {
		fmt.Println(err)
	}
	return MarshalJson(jsonData)
}
