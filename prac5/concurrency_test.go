package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://xx.com",
		"waat://yy.com",
	}
	actualResults := CheckWebsites(mockWebsiteChecker, websites)
	expectedResults := map[string]bool{
		"http://google.com": true,
		"http://xx.com":     true,
		"waat://yy.com":     false,
	}

	if !reflect.DeepEqual(expectedResults, actualResults) {
		t.Fatalf("Wanted %v, got %v", expectedResults, actualResults)
	}
}
