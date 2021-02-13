package main

import "fmt"

func main() {
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
