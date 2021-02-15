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
	path := "test"
	result, _ := OpenSrcFile(&path, "openfile.asc")

	for i, v := range result {
		if v != expected[i] {
			t.Error(v, expected[i])
		}
	}
}

func TestGetName(t *testing.T) {
	expected := "DLT001_Al_10_5"

	path := "test"
	file, _ := OpenSrcFile(&path, "DLT001_Al_10_5.dp_rpc_asc")
	result := GetName(file)

	if expected != result {
		t.Error(result, expected)
	}
}
