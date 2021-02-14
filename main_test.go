package main

import (
	"testing"
)

func TestGetSrcInDir(t *testing.T) {
	expected := []string{"main.go", "main_test.go"}
	var path = "."

	result, _ := GetSrcInDir(&path)

	for i, v := range result {
		if v != expected[i] {
			t.Error(v, expected[i])
		}
	}
}

func TestOpenSrcFile(t *testing.T) {
	expected := []string{"test", "test"}
	result, _ := OpenSrcFile("./test/openfile.asc")

	for i, v := range result {
		if v != expected[i] {
			t.Error(v, expected[i])
		}
	}
}
