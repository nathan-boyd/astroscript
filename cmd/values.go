package cmd

import "strings"

// a list of subDirectories which the application is allowed to delete jpg in
var subDirectories = [...]string{
	"Light",
	"Dark",
	"Bias",
	"Flat",
}

func stringContainsSlice(incString string, incList []string) bool {

	for _, value := range incList {
		if strings.Contains(incString, value) {
			return true
		}
	}

	return false
}

func stringInSlice(incString string, incList []string) bool {
	for _, b := range incList {
		if b == incString {
			return true
		}
	}
	return false
}

func sliceInSlice(sliceOne []string, sliceTwo []string) bool {
	for _, v1 := range sliceOne {
		if stringInSlice(v1, sliceTwo) {
			return true
		}
	}
	return false
}
