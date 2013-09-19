# gophpserialize

GO module to parse php serialised obj and help to convert to json.

Currently only support unserialze php array.

## Example

	import (
		"github.com/Wuvist/gophpserialize"
	)

	func main() {
		data := `a:3:{s:5:"apple";i:1;s:6:"orange";i:2;s:5:"grape";i:3;}`
		obj := gophpserialize.Unmarshal([]byte(data))
		...
	}

Need more test case & review on API.
