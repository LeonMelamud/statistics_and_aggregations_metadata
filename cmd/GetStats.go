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
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// GetStatsCmd represents the GetStats command
var GetStatsCmd = &cobra.Command{
	Use:   "GetStats",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GetStats called")
	},
}

func init() {
	rootCmd.AddCommand(GetStatsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// GetStatsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// GetStatsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type FileStats struct {
	NumFiles        int64           `json:"num_files"`
	LargestFile     LargestFileInfo `json:"largest_file"`
	AverageFileSize float64         `json:"avg_file_size"`
	MostFrequentExt ExtInfo         `json:"most_frequent_ext"`
	TextPercentage  float32         `json:"text_percentage"`
	MostRecentPaths []string        `json:"most_recent_paths"`
}
type LargestFileInfo struct {
	Path string `json:"path"`
	Size int64  `json:"size"`
}
type ExtInfo struct {
	Extension      string `json:"extension"`
	NumOccurrences int64  `json:"num_occurrences"`
}

// Number of files received
// Largest file received (including name and size)
// Average file size
// Most frequent file extension (including number of occurences)
// Percentage of text files of all files received
// List of latest 10 file paths received
func GetStats() {

	if _, err := os.Stat("./FilesMetadata.json"); os.IsNotExist(err) {
		check(err)
	}

	fmt.Println("in GetStats")

}
