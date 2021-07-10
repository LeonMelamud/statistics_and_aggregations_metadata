
#  Library that performs statistics and aggregations for file metadata
## Usage 

 #### STAGE 1

**CREATE A UNIT TEST FOR THE LIBRARY**
```bash
$ cd cmd //inside the root project folder
$ go test
```

 #### STAGE 2

**cd to root project folder**
```bash
$go run aquaStatistic // show the commands
```
**addFile with path,size,is_binary**
```bash
$go run aquaStatistic AddFile '{"path":"../README.md","size":2343,"is_binary":false}'
```
**GetStats() in JSON format when input ends.**
```bash
$ go run aquaStatistic GetStats
//to remove the file in the end of the work add flag --rm true
$go run aquaStatistic GetStats --rm true 
```
**to create executable**
```bash
//go to root project directory
$go build -o ./bin/aquastatistic .

//test the executable
$./bin/aquastatistic AddFile '{"path":"../README.md","size":2343,"is_binary":false}'
$./bin/aquastatistic GetStats
//to remove the file in the end of the work add flag --rm true
$./bin/aquastatistic GetStats --rm true 
```

#### Bonus: create a Dockerfile for the executable.
```bash
//see .Dockerfile, build from root folder
//check that the file is exsist ./bin/aquastatistic 
$docker build --file ./Dockerfile -t aquastatistic_prod .

//Tested on Mac OS , change the /tmp directory to your directory
$docker run -it -v /tmp:/root aquastatistic_prod AddFile '{"path":"../README.md","size":2343,"is_binary":false}'
$docker run -it -v /tmp:/root aquastatistic_prod GetStats

//to remove the file in the end of the work add flag --rm true
$docker run -it -v /tmp:/root aquastatistic_prod GetStats --rm true 

```
# Required Functions

**AddFile(metadata FileMetadata) error**

The function receives a structure containing the metadata of one file. This file should be taken into account
when calculating statistics. The function can return an error if the input is invalid or processing of the file fails.

 **GetStats() FileStats**

This function returns statistics for all files added until that point. The following statistics should be returned:

Number of files received
Largest file received (including name and size)
Average file size
Most frequent file extension (including number of occurences)
Percentage of text files of all files received
List of latest 10 file paths received
# Required Types/Structures
##Stage 1:
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
## Guidelines:

The library MAY define a structure/type/object on which the two required functions are to be called.

There is no guarantee that calls to AddFile() will happen sequentially.

Create a unit test for the library

##Stage 2:
Create a command line utility to uses the library. The utility reads a list of file metadata from standard input,
where each line is a JSON representation of the metadata. The utility calls AddFile() for each line, and finally
prints the output of GetStats() in JSON format when input ends.
Bonus: create a Dockerfile for the executable.
