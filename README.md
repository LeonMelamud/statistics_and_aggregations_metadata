
# library that performs statistics and aggregations for file metadata

# Required Functions

<p>AddFile(metadata FileMetadata) error</p>

The function receives a structure containing the metadata of one file. This file should be taken into account
when calculating statistics. The function can return an error if the input is invalid or processing of the file fails.

 <p><GetStats() FileStats</p>

This function returns statistics for all files added until that point. The following statistics should be returned:

Number of files received
Largest file received (including name and size)
Average file size
Most frequent file extension (including number of occurences)
Percentage of text files of all files received
List of latest 10 file paths received
# Required Types/Structures

The following structures should be used by the library:
```bash

type FileMetadata struct {
Path string `json:"path"` // the file's absolute path
Size int64 `json:"size"` // the file size in bytes
IsBinary bool `json:"is_binary"` // whether the file is a binary file or a simple text file
}

type FileStats struct {
NumFiles int64 `json:"num_files"`
LargestFile LargestFileInfo `json:"largest_file"`
AverageFileSize float64 `json:"avg_file_size"`
MostFrequentExt ExtInfo `json:"most_frequent_ext"`
TextPercentage float32 `json:"text_percentage"`
MostRecentPaths []string `json:"most_recent_paths"`
}

type LargestFileInfo struct {
Path string `json:"path"`
Size int64 `json:"size"`
}
type ExtInfo struct {
    Extension string `json:"extension"`
    NumOccurrences int64 `json:"num_occurrences"`
}
```
# Guidelines:

The library MAY define a structure/type/object on which the two required functions are to be called.

There is no guarantee that calls to AddFile() will happen sequentially.

Create a unit test for the library

# #Stage 2:
Create a command line utility to uses the library. The utility reads a list of file metadata from standard input,
where each line is a JSON representation of the metadata. The utility calls AddFile() for each line, and finally
prints the output of GetStats() in JSON format when input ends.
Bonus: create a Dockerfile for the executable.
