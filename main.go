package main

import (
	"fmt"
    "os"
    "io/fs"
    "time"
    "sort"
)

func main() {

    start := time.Now()
    listFiles(os.Args[1])

    elapsed := time.Since(start)

    fmt.Println("Time taken:", elapsed)
}

func listFiles(path string) {
    var files []string

    fileMap := make(map[string]int64)

    fileSystem:= os.DirFS(path) // set path to the directory you want to list

    fs.WalkDir(fileSystem, path, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if d.IsDir() {
            return nil
        }
        files = append(files, path)
        return nil
    })

    for _, fileName := range files {

        fileInfo, err := os.Stat(fileName)
        if err != nil {
            fmt.Println(err)
            continue
        }

        fileMap[fileName] = fileInfo.Size()

}

//sort the map by value 

keys := make([]string, 0, len(fileMap))
for k := range fileMap {
    keys = append(keys, k)
}

sort.Slice(keys, func(i, j int) bool {
    return fileMap[keys[i]] < fileMap[keys[j]]
})

for _, fileName := range keys {
    fmt.Printf("%-20s %s\n", getfileSize(fileMap[fileName]), fileName)
}

    fmt.Println()

}

//get the file size in human readable format
func getfileSize(fileSize int64) string {

        var size float64
        var result string

        if fileSize < 1024 {
            size = float64(fileSize)
            result = fmt.Sprint(size)+ " B"

        } else if fileSize > 1024 && fileSize < 1048576 {
            size = float64(fileSize/1024)
            result = fmt.Sprint(size)+ " KB"

        } else if fileSize > 1048576 && fileSize < 1073741824 {
            size = float64(fileSize/1048576)
            result = fmt.Sprint(size)+ " MB"

        } else {    
            size = float64(fileSize/1073741824)
            result = fmt.Sprint(size)+ " GB"
        }
        return result
}

