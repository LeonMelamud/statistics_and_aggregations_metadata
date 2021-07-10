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
	"github.com/spf13/cobra"
)

type FileMetadata struct {
	Path     string `json:"path"`      // the file's absolute path
	Size     int64  `json:"size"`      // the file size in bytes
	IsBinary bool   `json:"is_binary"` // whether the file is a binary file or a simple text file
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

var cfgFile string
var METADATA_FILE_PATH = UserHomeDir() + "/FilesMetadata.json"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aquaStatistic",
	Short: "library that performs statistics and aggregations for file metadata",
	Long:  ``,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

}
