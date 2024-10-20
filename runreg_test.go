package main

import (
	"testing"
)

func TestRunReg(t *testing.T) {

	// define testdata
	testdata := [][]string{
		{"xvar1", "xvar2", "xvar3", "xvar4", "yvar1", "yvar2", "yvar3", "yvar4"},
		{"1", "2", "4", "19", "17", "1", "1", "3"},
		{"14", "7", "5", "14", "14", "8", "9", "5"},
		{"12", "13", "10", "11", "11", "19", "7", "6"},
		{"5", "1", "16", "14", "12", "9", "13", "8"},
		{"6", "11", "0", "1", "1", "0", "11", "3"},
		{"8", "15", "10", "3", "13", "3", "10", "9"},
		{"3", "12", "0", "6", "2", "6", "13", "15"},
		{"8", "10", "4", "5", "17", "1", "1", "13"},
	}

	r1, r2, r3, r4 := RunReg(testdata)

	correct_ans_reg1 := "Predicted = 8.8843 + xvar1*0.2794"
	correct_ans_reg2 := "Predicted = 4.7163 + xvar2*0.1306"
	correct_ans_reg3 := "Predicted = 7.1791 + xvar3*0.1544"
	correct_ans_reg4 := "Predicted = 10.2286 + xvar4*-0.2716"

	// compare parsed data to the data it should be parsed to
	if r1.Formula != correct_ans_reg1 {
		t.Errorf("parsed incorrectly")
	}

	if r2.Formula != correct_ans_reg2 {
		t.Errorf("parsed incorrectly")
	}

	if r3.Formula != correct_ans_reg3 {
		t.Errorf("parsed incorrectly")
	}

	if r4.Formula != correct_ans_reg4 {
		t.Errorf("parsed incorrectly")
	}

}

func BenchmarkRunReg(b *testing.B) {
	// define testdata
	testdata := [][]string{
		{"xvar1", "xvar2", "xvar3", "xvar4", "yvar1", "yvar2", "yvar3", "yvar4"},
		{"1", "2", "4", "19", "17", "1", "1", "3"},
		{"14", "7", "5", "14", "14", "8", "9", "5"},
		{"12", "13", "10", "11", "11", "19", "7", "6"},
		{"5", "1", "16", "14", "12", "9", "13", "8"},
		{"6", "11", "0", "1", "1", "0", "11", "3"},
		{"8", "15", "10", "3", "13", "3", "10", "9"},
		{"3", "12", "0", "6", "2", "6", "13", "15"},
		{"8", "10", "4", "5", "17", "1", "1", "13"},
	}

	RunReg(testdata)

}
