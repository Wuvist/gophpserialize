# gophpserialize

GO module to parse php serialised obj and help to convert to json and serialize obj to php serialised obj and help to convert from json.

Currently not supported php object.

## Example

	import (
		"github.com/linluxiang/gophpserialize"
	)

	func main() {
		data := `a:3:{s:5:"apple";i:1;s:6:"orange";i:2;s:5:"grape";i:3;}`
		obj := gophpserialize.Unmarshal([]byte(data))

        jsonData, err := gophpserialize.PhpToJson([]byte(data))
		...
		phpobj, err := gophpserialize.MarshalNormal(obj)
        // phpobc == data

        newphpobj, err := gophpserialize.JsonToPhp(jsonData)
        // newphpobj == data
	}

Need more test case & review on API.

##ChangeLog:
* version 0.1: support JsonToPhp and Marshal
