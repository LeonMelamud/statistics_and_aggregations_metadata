/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	//"sync"
	"encoding/json"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var userinput string

// AddFileCmd represents the AddFile command
var AddFileCmd = &cobra.Command{
	Use:   "AddFile",
	Short: "The function receives a structure containing the metadata of one file",
	Long: `The function receives a structure containing the metadata of one file. This file should be taken into account
	when calculating statistics. The function can return an error if the input is invalid or processing of the file fails.
	"{path": "../README.md", "size": "2323232",is_binary:false}`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("userinput content is %s\n", rootCmd.Flag("userinput").Value)
		// define slice of FileMetadata
		data := FileMetadata{}
		json.Unmarshal([]byte(userinput), &data)
		err := getFile(data)
		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(AddFileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// AddFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// AddFileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&userinput, "userinput", "u", "", "userinput json configuration ")

}

type FileMetadata struct {
	Path     string `json:"path"`      // the file's absolute path
	Size     int64  `json:"size"`      // the file size in bytes
	IsBinary bool   `json:"is_binary"` // whether the file is a binary file or a simple text file
}

var fileMutex sync.Mutex

func getFile(metadata FileMetadata) error {
	fileMutex.Lock() // Use a single mutex to serialize all access to file
	if _, err := os.Stat("./FilesMetadata.json"); os.IsNotExist(err) {
		f, err := os.Create("./FilesMetadata.json")
		if err != nil {
			return errors.New(err.Error())
		}
		defer f.Close()
	}

	byteValue, err := ioutil.ReadFile("./FilesMetadata.json")
	if err != nil {
		return errors.New(err.Error())
	}

	fmt.Println("Successfully Opened FilesMetadata.json")

	filesMetadata := []FileMetadata{}
	json.Unmarshal(byteValue, &filesMetadata)

	for i := 0; i < len(filesMetadata); i++ {
		fmt.Println("Path: ", filesMetadata[i].Path)
		fmt.Println("Size: ", filesMetadata[i].Size)
		fmt.Println("IsBinary: ", filesMetadata[i].IsBinary)

	}
	//her we will get the file from path
	file1 := FileMetadata{}
	file1.IsBinary = true
	file1.Size = 3434
	file1.Path = "/root/check"
	file2 := FileMetadata{}
	file2.IsBinary = true
	file2.Size = 3434
	file2.Path = "/root/check"

	//TODO : append the real metadata
	if metadata.Path != "" && metadata.Size != 0 {
		return errors.New("metadata is not correct")
	}

	filesMetadata = append(filesMetadata, file1, file2)
	fmt.Println("file array : ", filesMetadata)
	j, err := json.Marshal(filesMetadata)
	if err != nil {
		return errors.New("error occured during marshaling. error")
	}
	// define slice of FileMetadata
	var jsonStruct []FileMetadata
	// Unmarshall it
	if err := json.Unmarshal(j, &jsonStruct); err != nil {
		log.Println(err)
	}
	fmt.Println("json unmarshell to struct  : ", jsonStruct)

	file, _ := json.MarshalIndent(jsonStruct, "", " ")

	err = ioutil.WriteFile("FilesMetadata.json", file, 0775)
	if err != nil {
		return errors.New(err.Error())
	}

	defer fileMutex.Unlock() //unlock even if there an error
	return nil
}
