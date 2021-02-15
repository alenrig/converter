package main

import (
	"testing"
)

var testPath string = "test"

func TestGetSrcInDir(t *testing.T) {
	expected := []string{"DLT001_Al_10_5.dp_rpc_asc"}

	result, _ := GetSrcInDir(&testPath)

	for i, v := range result {
		if v != expected[i] {
			t.Error(v, expected[i])
		}
	}
}

func TestOpenSrcFile(t *testing.T) {
	expected := []string{"test", "test"}
	result, _ := OpenSrcFile(&testPath, "openfile.asc")

	for i, v := range result {
		if v != expected[i] {
			t.Error(v, expected[i])
		}
	}
}

func TestGetName(t *testing.T) {
	expected := "DLT001_Al_10_5"

	file, _ := OpenSrcFile(&testPath, "DLT001_Al_10_5.dp_rpc_asc")
	result := GetName(file)

	if expected != result {
		t.Error(result, expected)
	}
}

func TestCutDatapoints(t *testing.T) {
	expected := []string{"133Cs 27Al", "133Cs 69Ga", "133Cs 75As"}

	file, _ := OpenSrcFile(&testPath, "DLT001_Al_10_5.dp_rpc_asc")
	result, _ := CutDatapoints(file)

	for i, v := range result {
		if v != expected[i] {
			t.Error(v, expected[i])
		}
	}
}
