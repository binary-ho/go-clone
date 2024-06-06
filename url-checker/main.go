package url_checker

import (
	"clone-project/url-checker/urls"
	"fmt"
)

func Run() {
	channel := make(chan string)
	resultSize := urls.HitUrl(channel)
	for _ = range resultSize {
		fmt.Println(<-channel)
	}
}
