package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const fileType string = ".go"

func main() {
	srcFiles, err := GetSrcInDir()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, srcFile := range srcFiles {
			fmt.Println(srcFile)
		}
	}
}

// GetSrcInDir - Get all files with needed file type in current dir
func GetSrcInDir() ([]string, error) {
	srcFiles := []string{}
	filesInDir, err := ioutil.ReadDir(".")
	if err != nil {
		return nil, err
	}
	for _, fileInDir := range filesInDir {
		if strings.HasSuffix(fileInDir.Name(), fileType) && !fileInDir.IsDir() {
			srcFiles = append(srcFiles, fileInDir.Name())
		}
	}

	return srcFiles, nil
}
