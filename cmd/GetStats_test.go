package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_GetStats(t *testing.T) {
	//need to populate the file first
	args := []string{
		string(`{"path":"1.txt","size":1111,"is_binary":false}`),
		string(`{"path":"2.txt","size":2222,"is_binary":true}`),
		string(`{"path":"3.md","size":3333,"is_binary":true}`),
		string(`{"path":"4.md","size":4444,"is_binary":true}`),
		string(`{"path":"5.env","size":5555,"is_binary":true}`),
		string(`{"path":"6.md","size":6666,"is_binary":false}`),
		string(`{"path":"7.md5","size":7777,"is_binary":false}`),
		string(`{"path":"8.md5","size":8888,"is_binary":false}`),
		string(`{"path":"9.cer","size":9999,"is_binary":false}`),
		string(`{"path":"10.cer","size":101010,"is_binary":true}`),
		string(`{"path":"11.cer","size":111111,"is_binary":true}`),
		string(`{"path":".12.json","size":121212,"is_binary":false}`),
	}
	addFile(t, "yes", args)

	cmd := GetStatsCmd()
	b := bytes.NewBufferString("")

	cmd.SetOut(b)
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Test_GetStats result: ", string(out))

}
