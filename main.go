package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const fileType string = ".go"

func main() {
	srcFiles, err := GetSrcInDir()
	if err != nil {
		log.Fatal(err)
	}
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

// OpenSrcFile opens given filepath as slice of strings.
func OpenSrcFile(srcFile string) ([]string, error) {
	bytesContent, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return nil, err
	}
	oneStringContent := string(bytesContent)
	slicedContent := strings.Split(oneStringContent, "\n")
	slicedContent = deleteEmpty(slicedContent)

	return slicedContent, nil
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
