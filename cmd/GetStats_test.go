package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_GetStats(t *testing.T) {
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
		string(`{"path":"11.cer","size":111111,"is_binary":true}`),
		string(`{"path":".12.json","size":121212,"is_binary":false}`),
	}
	//we passing yes to remove the file so next time we will have
	addFile(t, "yes", args)

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

	if fileStats.NumFiles != 12 && fileStats.LargestFile.
		Path != ".12.json" && fileStats.LargestFile.Size != 121212 && fileStats.AverageFileSize != 31944 && fileStats.MostFrequentExt.
		Extension != ".md" && fileStats.MostFrequentExt.NumOccurrences != 3 && fileStats.TextPercentage != 0 && len(fileStats.MostRecentPaths) != 10 {

		t.Fatal("fileStats Data is not correct, got : " + string(out))
	}

	fmt.Println("Test_GetStats: Done")

}
