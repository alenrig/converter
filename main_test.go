package main

import "testing"

func TestGetSrcInDir(t *testing.T) {
	expected := []string{"main.go", "main_test.go"}

	result, _ := GetSrcInDir()

	for i, v := range result {
		if v != expected[i] {
			t.Error("Fuck!")
		}
	}
}

func TestOpenSrcFile(t *testing.T) {
	expected := []string{"test"}
	result, _ := OpenSrcFile("./test/openfile.asc")

	if result != expected {
		t.Error("Fuck!")
	}
}
