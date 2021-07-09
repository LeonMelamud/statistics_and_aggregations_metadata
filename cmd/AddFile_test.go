package cmd

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func addFile(t *testing.T, rm string, args []string) {
	cmd := AddFileCmd()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	if len(args) == 0 {
		s := string(`{"path":"../README.md","size":2343,"is_binary":false}`)
		cmd.SetArgs([]string{s})
		cmd.Execute()
	} else {
		for _, v := range args {
			cmd.SetArgs([]string{v})
			cmd.Execute()
		}

	}

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if rm == "" {
		e := os.Remove(METADATA_FILE_PATH)

		if e != nil {
			t.Fatal(e)
		}
	}
	result := string("done adding file to: " + METADATA_FILE_PATH)
	if !strings.Contains(string(out), METADATA_FILE_PATH) {
		t.Fatalf("expected \"%s\" got \"%s\"", result, string(out))
	}

}
func Test_AddFileCmd(t *testing.T) {
	var rm string
	var args []string
	addFile(t, rm, args)

}
