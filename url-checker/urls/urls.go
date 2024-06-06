package urls

import (
	"fmt"
	"net/http"
	"strconv"
)

type urls []string

var urlsInstance = urls{
	"https://www.naver.com/",
	"https://github.com/binary-ho/",
	"https://github.com/binary-hohohohohohohoho/",
	"https://www.google.com/",
	"https://nomadcoders.co/",
	"https://www.naver.com/",
	"https://github.com/binary-ho/",
	"https://github.com/binary-hohohohohohohoho/",
	"https://www.google.com/",
	"https://nomadcoders.co/",
	"https://www.naver.com/",
	"https://github.com/binary-ho/",
	"https://github.com/binary-hohohohohohohoho/",
	"https://www.google.com/",
	"https://nomadcoders.co/",
	"https://www.naver.com/",
	"https://github.com/binary-ho/",
	"https://github.com/binary-hohohohohohohoho/",
	"https://www.google.com/",
	"https://nomadcoders.co/",
}

func getUrlsInstance() urls {
	return urlsInstance
}

func PrintUrls() {
	for _, url := range getUrlsInstance() {
		fmt.Println(url)
	}
}

func getUrlSize() int {
	return len(getUrlsInstance())
}

func HitUrl(channel chan<- string) int {
	for index, url := range getUrlsInstance() {
		go hitUrl(index, url, channel)
	}
	return getUrlSize()
}

func hitUrl(index int, url string, channel chan<- string) {
	result := isValid(url)
	channel <- strconv.Itoa(index) + " : " + url + " -> " + result
}

func isValid(url string) string {
	response, err := http.Get(url)
	if err == nil && response.StatusCode < 400 {
		return "OK"
	}
	return "NOOOOO"
}
