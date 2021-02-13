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
