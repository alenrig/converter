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

func TestCutHeader(t *testing.T) {
	expected := []string{"133Cs 27Al", "133Cs 69Ga", "133Cs 75As"}

	file, _ := OpenSrcFile(&testPath, "DLT001_Al_10_5.dp_rpc_asc")
	result, _ := CutDatapoints(file)

	for i, v := range result {
		if v != expected[i] {
			t.Error(v, expected[i])
		}
	}
}

func TestCutDatapoints(t *testing.T) {
	expectedFirstLine := "2.40000E-002		1.26501E+002	5.46667E-002		8.71941E+003	8.40000E-002		6.51140E+003"
	expectedLastLine := "1.69520E+001		1.47737E+001	1.69827E+001		1.77897E+005	1.70200E+001		3.06937E+004"

	file, _ := OpenSrcFile(&testPath, "DLT001_Al_10_5.dp_rpc_asc")
	_, result := CutDatapoints(file)

	if result[0] != expectedFirstLine {
		t.Error(result[0], expectedFirstLine)
	}

	if result[len(result)-1] != expectedLastLine {
		t.Error(result[len(result)-1], expectedLastLine)
	}
}

func TestParseHead(t *testing.T) {
	expected := []string{"time", "133Cs 27Al", "133Cs 69Ga", "133Cs 75As"}

	file, _ := OpenSrcFile(&testPath, "DLT001_Al_10_5.dp_rpc_asc")
	header, _ := CutDatapoints(file)

	result := ParseHeader(header)

	for i, v := range expected {
		if v != result[i] {
			t.Error(v, result[i])
		}
	}
}

func TestParseDatapoints(t *testing.T) {
	expectedFirstLine := "2.40000E-002,1.26501E+002,8.71941E+003,6.51140E+003"
	expectedLastLine := "1.69520E+001,1.47737E+001,1.77897E+005,3.06937E+004"

	file, _ := OpenSrcFile(&testPath, "DLT001_Al_10_5.dp_rpc_asc")
	_, result := CutDatapoints(file)

	if result[0] != expectedFirstLine {
		t.Error(result[0], expectedFirstLine)
	}

	if result[len(result)-1] != expectedLastLine {
		t.Error(result[len(result)-1], expectedFirstLine)
	}
}

func TestFindIndexesToDelete(t *testing.T) {
	array := []string{"2.40000E-002", "1.26501E+002", "5.46667E-002", "8.71941E+003", "8.40000E-002", "6.51140E+003"}
	expected := []int{2, 4}

	result := FindIndexesToDelete(array)

	for i, v := range expected {
		if v != result[i] {
			t.Error(v, result[i])
		}
	}
}
