package gophpserialize

import "testing"
import "fmt"

func TestMarshal(t *testing.T) {
	result := `a:3:{s:5:"apple";i:1;s:6:"orange";i:2;s:5:"grape";i:3;}`
	data := make(map[string]interface{})
	data["apple"] = 1
	data["orange"] = 2
	data["grape"] = 3

	obj, err := MarshalNormal(data)

	if err != nil {
		t.Error(err)
	}

	if string(obj) != result {
		fmt.Println("error string: ", string(obj))
		fmt.Println("corrent string: ", result)
		t.Error("Marshal failed")
	}
}

func TestJsonToPhp(t *testing.T) {
	data := `a:3:{s:5:"apple";i:1;s:6:"orange";i:2;s:5:"grape";i:3;}`

	jsonStr := `{"apple":1,"orange":2,"grape":3}`
	obj, err := JsonToPhp([]byte(jsonStr))

	if err != nil {
		t.Error(err)
	}

	if string(obj) != data {
		fmt.Println("error string: ", string(obj))
		fmt.Println("corrent string: ", data)
		t.Error("convert to php error")
	}
}
