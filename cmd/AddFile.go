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

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

//var userinput string
var in string

// AddFileCmd represents the AddFile command
func AddFileCmd() *cobra.Command {
	addFileCmd := &cobra.Command{
		Use:   "AddFile",
		Short: "The function receives a structure containing the metadata of one file",
		Long: `The function receives a structure containing the metadata of one file. This file should be taken into account
	when calculating statistics. The function can return an error if the input is invalid or processing of the file fails.
	"{path": "../README.md", "size": "2323232",is_binary:false}`,
		RunE: func(cmd *cobra.Command, args []string) error {
			data := FileMetadata{}
			if err := json.Unmarshal([]byte(args[0]), &data); err != nil {
				fmt.Println(fmt.Errorf("AddFileCmd , error unmarshaling data '%v': %v", data, err.Error()))
				log.Fatal(err)
			}

			err := getFile(data)
			if err != nil {
				log.Fatal(err)
			}
			color.Blue.Printf("set json file to directory : " + "\n")
			color.LightMagenta.Printf(METADATA_FILE_PATH + "\n")
			return nil
		},
	}
	addFileCmd.Flags().StringVar(&in, "FileMetadata", "", "FileMetadata input.")
	return addFileCmd
}

func init() {
	addFileCmd := AddFileCmd()
	rootCmd.AddCommand(addFileCmd)
}

var fileMutex sync.Mutex

func getFile(metadata FileMetadata) error {
	fileMutex.Lock() // Use a single mutex to serialize all access to file
	if _, err := os.Stat(METADATA_FILE_PATH); os.IsNotExist(err) {
		f, err := os.Create(METADATA_FILE_PATH)
		if err != nil {
			return errors.New(err.Error())
		}
		defer f.Close()
	}

	byteValue, err := ioutil.ReadFile(METADATA_FILE_PATH)
	if err != nil {
		return errors.New(err.Error())
	}

	filesMetadata := []FileMetadata{}
	if len(byteValue) > 0 {
		if err := json.Unmarshal(byteValue, &filesMetadata); err != nil {
			return errors.New("error occured during Unmarshal. error")
		}
	}
	if metadata.Path == "" && metadata.Size == 0 {
		return errors.New("metadata is not correct")
	}

	filesMetadata = append(filesMetadata, metadata)
	j, err := json.Marshal(filesMetadata)
	if err != nil {
		return errors.New("error occured during marshaling. error")
	}
	// define slice of FileMetadata
	var jsonStruct []FileMetadata
	// Unmarshall it
	if err := json.Unmarshal(j, &jsonStruct); err != nil {
		log.Println(err)
		return errors.New("error occured during Unmarshal. error")
	}
	color.Cyan.Printf("getFile, json unmarshell to struct  : ")
	color.Tag("info").Println(jsonStruct)
	file, _ := json.MarshalIndent(jsonStruct, "", " ")

	err = ioutil.WriteFile(METADATA_FILE_PATH, file, 0775)
	if err != nil {
		return errors.New(err.Error())
	}

	defer fileMutex.Unlock() //unlock even if there an error
	return nil
}
