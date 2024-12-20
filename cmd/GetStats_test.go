package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func Test_GetStats(t *testing.T) {
	//if file exsist remove it to make clean test
	if _, err := os.Stat(METADATA_FILE_PATH); !os.IsNotExist(err) {
		e := os.Remove(METADATA_FILE_PATH)
		if e != nil {
			t.Fatal(e)
		}
	}
	//need to populate the file first
	args := []string{
		//testingArgs will be tested in AddFile and they need to be the samre as we test the index 0
		string(testingArgs),
		string(`{"path":"2.txt","size":2222,"is_binary":true}`),
		string(`{"path":"3.md","size":3333,"is_binary":true}`),
		string(`{"path":"4.md","size":4444,"is_binary":true}`),
		string(`{"path":"5.env","size":5555,"is_binary":true}`),
		string(`{"path":"6.md","size":6666,"is_binary":false}`),
		string(`{"path":"7.md5","size":7777,"is_binary":false}`),
		string(`{"path":"8.md5","size":8888,"is_binary":false}`),
		string(`{"path":"9.cer","size":9999,"is_binary":false}`),
		string(`{"path":"10.cer","size":101010,"is_binary":true}`),
		string(`{"path":"11.env","size":111111,"is_binary":true}`),
		string(`{"path":".12.json","size":121212,"is_binary":false}`),
	}
	//we passing no to not remove the file and append all
	addFile(t, "no", args)

	cmd := GetStatsCmd()
	b := bytes.NewBufferString("")

	cmd.SetOut(b)
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	fileStats := FileStats{}
	if err := json.Unmarshal(out, &fileStats); err != nil {
		fmt.Println(fmt.Errorf("AddFileCmd , error unmarshaling data '%v': %v", fileStats, err.Error()))
		t.Fatal(err)
	}
	fmt.Println("cmd result: ", string(out))

	//if one of them is not the same as result it fail
	if fileStats.NumFiles != 12 || fileStats.LargestFile.
		Path != ".12.json" || fileStats.LargestFile.Size != 121212 || fileStats.AverageFileSize != 31944 || fileStats.MostFrequentExt.
		NumOccurrences != 3 || fileStats.TextPercentage != 0 || len(fileStats.MostRecentPaths) != 10 || fileStats.MostFrequentExt.
		Extension != ".md" {

		t.Fatal("fileStats Data is not correct, got : " + string(out))
	}

	fmt.Println("Test_GetStats: Done")

}
