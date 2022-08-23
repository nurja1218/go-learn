package util

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// func JsonPrettyPrint(in string) string {
// 	var out bytes.Buffer
// 	err := json.Indent(&out, []byte(in), "", "\t")
// 	if err != nil {
// 		return in
// 	}
// 	return out.String()
// }

func PrintJson(in any) {
	var out bytes.Buffer
	byteArr, _ := json.Marshal(in)
	err := json.Indent(&out, []byte(byteArr), "", "\t")
	if err != nil {
		fmt.Println(string(byteArr))
	}
	fmt.Println(out.String())
}
