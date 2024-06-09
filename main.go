package main

import (
	"clone-project/jobscrapper"
)

func main() {
	//banking.Run()
	//url_checker.Run()
	jobScrapper := jobscrapper.GetJobScrapper()
	//fmt.Println(jobScrapper.getPages())
	pageCount := jobScrapper.GetPageCount()
	jobsChannel := make(chan []jobscrapper.Job)
	for page := range pageCount {
		go jobScrapper.GetPage(page, jobsChannel)
	}

	var result []jobscrapper.Job
	for range pageCount {
		jobs := <-jobsChannel
		result = append(result, jobs...)
	}
}
