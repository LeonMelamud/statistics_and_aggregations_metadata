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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rm string

// GetStatsCmd represents the GetStats command
func GetStatsCmd() *cobra.Command {
	getStatsCmd := &cobra.Command{
		Use:   "GetStats",
		Short: "This function returns statistics for all files added until that point",
		Long: `This function returns statistics for all files added until that point. The following statistics should be returned:
	Number of files received
	Largest file received (including name and size)
	Average file size
	Most frequent file extension (including number of occurences)
	Percentage of text files of all files received
	List of latest 10 file paths received.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("GetStats called")

			//fmt.Printf("remove file flag is %s\n", cmd.Flag("rm").Value)

			fileStats := GetStats()

			cmd.Println("JSON format : \n", fileStats)
			//return jsonStruct
			return nil
		},
	}
	getStatsCmd.Flags().StringVar(&rm, "rm", "", "getStatsCmd clear MetaData file")
	return getStatsCmd
}

func init() {
	getStatsCmd := GetStatsCmd()
	rootCmd.AddCommand(getStatsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// GetStatsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// GetStatsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func GetStats() string {
	if _, err := os.Stat(METADATA_FILE_PATH); os.IsNotExist(err) {
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	byteValue, err := ioutil.ReadFile(METADATA_FILE_PATH)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("GetStats", "Successfully Opened FilesMetadata.json")

	filesMetadata := []FileMetadata{}

	json.Unmarshal(byteValue, &filesMetadata)

	fmt.Println("GetStats ,filesMetadata structs : ", filesMetadata)

	var sumOfSizes float64
	//Number of files received
	largestFileSize := filesMetadata[0].Size
	//
	largestFilePath := filesMetadata[0].Path

	mapExt := make(map[string]int)
	numOfFiles := len(filesMetadata)
	latest10Pates := make([]string, 10)

	for i := 0; i < numOfFiles; i++ {
		fileSize := filesMetadata[i].Size
		if largestFileSize < fileSize {
			largestFileSize = fileSize
			largestFilePath = filesMetadata[i].Path
		}
		fileExtension := filepath.Ext(filesMetadata[i].Path)
		if val, ok := mapExt[fileExtension]; ok {
			mapExt[fileExtension] = val + 1
		} else {
			mapExt[fileExtension] = 1
		}
		if i < 10 {
			latest10Pates[i] = filesMetadata[i].Path
		}
		sumOfSizes += float64(fileSize)
	}

	fileStats := FileStats{}
	//Number of files received
	fileStats.NumFiles = int64(numOfFiles)
	//Average file size
	fileStats.AverageFileSize = sumOfSizes / float64(numOfFiles)
	//Largest file received (including name and size)
	largestFileInfo := LargestFileInfo{}
	largestFileInfo.Path = largestFilePath
	largestFileInfo.Size = largestFileSize
	fileStats.LargestFile = largestFileInfo
	//Most frequent file extension (including number of occurences)
	fileStats.MostFrequentExt = mostFrequentExt(mapExt)
	//Percentage of text files of all files received
	numberOfTxtFiles := mapExt[".txt"]
	fileStats.TextPercentage = float32(numberOfTxtFiles/numOfFiles) * 100
	//List of latest 10 file paths received
	if numOfFiles < 10 {
		fileStats.MostRecentPaths = latest10Pates[:numOfFiles]
	} else {
		fileStats.MostRecentPaths = latest10Pates
	}
	j, err := json.Marshal(fileStats)
	if err != nil {
		log.Fatal("json Marshal error")
	}
	return string(j)
}
