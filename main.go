package gophpserialize

import (
	"fmt"
	"github.com/linluxiang/gophpserialize"
)

func main() {
	data := `a:3:{s:5:"apple";i:1;s:6:"orange";i:2;s:5:"grape";i:3;}`
	obj := gophpserialize.Unmarshal([]byte(data))

	phpobj, nil := gophpserialize.MarshalNormal(obj)
	// phpobc == data
	fmt.Println(phpobj)

	newphpobj, nil := gophpserialize.MarshalJson(`{"apple":1, "orange":2, "grape": 3}`)
	// newphpobj == data
	fmt.Println(newphpobj)
}
