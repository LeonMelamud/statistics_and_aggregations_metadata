package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var testingArgs = `{"path":"1.txt","size":1111,"is_binary":false}`

func addFile(t *testing.T, rm string, args []string) {
	cmd := AddFileCmd()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	if len(args) == 0 {

		//the shold be the same
		s := string(testingArgs)
		cmd.SetArgs([]string{s})
		cmd.Execute()
	} else {
		for _, v := range args {
			cmd.SetArgs([]string{v})
			cmd.Execute()
		}

	}
	//test that the file exsist
	if _, err := os.Stat(METADATA_FILE_PATH); os.IsNotExist(err) {
		if err != nil {
			t.Fatalf("file not exsist")
		}
	}
	//read the file created
	byteValue, err := ioutil.ReadFile(METADATA_FILE_PATH)
	if err != nil {
		t.Fatalf(err.Error())
	}
	//parse the struct from the file
	filesMetadata := []FileMetadata{}
	if len(byteValue) > 0 {
		if err := json.Unmarshal(byteValue, &filesMetadata); err != nil {
			t.Fatalf("error occured during Unmarshal. error")
		}
	}
	j, err := json.Marshal(filesMetadata)
	if err != nil {
		t.Fatalf("error occured during marshaling. error")
	}

	var jsonStruct []FileMetadata

	if err := json.Unmarshal(j, &jsonStruct); err != nil {
		t.Fatalf(err.Error())
	}
	//check that the struct from the file is same as expected
	if jsonStruct[0].IsBinary != false && jsonStruct[0].Path != "1.txt" && jsonStruct[0].Size != 1111 {
		t.Fatalf("jsonStruct is no as expected ")

	}

	if rm == "" {
		e := os.Remove(METADATA_FILE_PATH)

		if e != nil {
			t.Fatal(e)
		}

	}

}
func Test_AddFileCmd(t *testing.T) {
	var rm string
	var args []string
	addFile(t, rm, args)
	fmt.Println("Test_AddFileCmd: Done")
}
