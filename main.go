package main

import (
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

	for _, srcFile := range srcFiles {
		slicedContent, err := OpenSrcFile(srcFile)
		if err != nil {
			log.Fatal(err)
		}
		name := GetHeader(slicedContent)
		header, datapoints := CutDatapoints(slicedContent)
		ions := parseHeader(header)
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

func GetHeader(slicedContent []string) string {
	slicedNameString := strings.Split(slicedContent[2], "\t")
	fullName := slicedContent[len(slicedNameString)-1]
	name := strings.Split(fullName, ".")[0]

	return name
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

// CutDatapoints cuts datapoints from src file.
// Strings ***<SOMETHING>*** are benchmarks - there are always a datapoints inside this range.
func CutDatapoints(slicedContent []string) (string, []string) {
	startLine := findIndexByContent(slicedContent, "*** DATA START ***") + 3
	endLine := findIndexByContent(slicedContent, "*** DATA END ***") - 1

	header := slicedContent[startLine]
	datapoints := slicedContent[startLine+2 : endLine]

	return header, datapoints
}

func findIndexByContent(slicedContent []string, contentToFind string) int {
	var result int
	for i, v := range slicedContent {
		if v == contentToFind {
			result = i
		}
	}

	return result
}

func parseHeader(rawString string) []string {
	slicedHeader := strings.Split(rawString, "\t")
	ions := deleteEmpty(slicedHeader)

	return ions
}
