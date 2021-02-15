package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const fileType string = ".dp_rpc_asc"

func main() {
	pathFlag := flag.String("p", ".", "working directory")
	flag.Parse()

	srcFiles, err := GetSrcInDir(pathFlag)
	if err != nil {
		log.Fatal(err)
	}

	for _, srcFile := range srcFiles {
		slicedContent, err := OpenSrcFile(pathFlag, srcFile)
		if err != nil {
			log.Fatal(err)
		}

		name := GetName(slicedContent)
		header, datapoints := CutDatapoints(slicedContent)
		//ions := parseHeader(header)
		fmt.Println(name, header, datapoints)
	}
}

// GetSrcInDir - Get all files with needed file type in current dir
func GetSrcInDir(path *string) ([]string, error) {
	srcFiles := []string{}
	filesInDir, err := ioutil.ReadDir(*path)
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
func OpenSrcFile(path *string, srcFile string) ([]string, error) {
	bytesContent, err := ioutil.ReadFile(*path + "/" + srcFile)
	if err != nil {
		return nil, err
	}
	oneStringContent := string(bytesContent)
	slicedContent := strings.Split(oneStringContent, "\n")
	slicedContent = deleteEmpty(slicedContent)

	return slicedContent, nil
}

// GetName gets original filename.
func GetName(slicedContent []string) string {
	slicedNameString := strings.Split(slicedContent[2], "\t")
	fullName := slicedNameString[len(slicedNameString)-1]
	name := strings.Split(fullName, ".")[0]

	return name
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		i := strings.TrimRight(str, "\t\r\n")
		if i != "" {
			r = append(r, i)
		}
	}
	return r
}

// CutDatapoints cuts datapoints from src file.
// Strings ***<SOMETHING>*** are benchmarks - there are always a datapoints inside this range.
func CutDatapoints(slicedContent []string) ([]string, []string) {
	startLine := findIndexByContent(slicedContent, "*** DATA START ***") + 2
	endLine := findIndexByContent(slicedContent, "*** DATA END ***")

	rawHeader := slicedContent[startLine]
	header := strings.Split(rawHeader, "\t")
	header = deleteEmpty(header)
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
	ions = append([]string{"time"}, ions...)

	return ions
}
