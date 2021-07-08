package cmd

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_ExecuteCommand(t *testing.T) {
	s := string(`{"path": "../README.md", "size": "2323232",is_binary:false}`)

	fmt.Printf("useDarToTest content is %s\n", s)
	data := FileMetadata{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("Operation: %s , %i , %b", data.Path, data.Size, data.IsBinary)
	getFile(data)

}
