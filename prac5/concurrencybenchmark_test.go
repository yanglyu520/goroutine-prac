package concurrency

import (
	"testing"
	"time"
)

// dependency injection function for benchmark test
func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func make100Urls() []string {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	return urls
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make100Urls()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
