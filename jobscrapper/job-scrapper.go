package jobscrapper

import (
	"clone-project/jobscrapper/urls"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type jobScrapper struct {
	urls urls.Urls
}

var jobScrapperInstance = &jobScrapper{urls: urls.GetInstance()}

const (
	PaginationClassName   = ".pagination"
	ATag                  = "a"
	SearchCardsClassName  = ".jobsearch-SerpJobCard"
	SearchCardIdClassName = "data-jk"
)

func GetJobScrapper() jobScrapper {
	return *jobScrapperInstance
}

func (jabScrapper *jobScrapper) GetPage(page int, jobsChannel chan<- []Job) {
	channel := make(chan *Job)
	pageUrl := jabScrapper.urls.GetPageUrl(page)
	response, err := http.Get(pageUrl)
	if isFail(response, err) {
		jobsChannel <- []Job{}
		return
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if isFail(nil, err) {
		jobsChannel <- []Job{}
		return
	}

	var jobs []Job
	searchCards := document.Find(SearchCardsClassName)
	searchCards.Each(func(i int, selection *goquery.Selection) {
		go CreateJob(selection, channel)
		//job := CreateJob(selection)
		//jobs = append(jobs, *job)
	})

	for range searchCards.Length() {
		job := <-channel
		jobs = append(jobs, *job)
	}
	jobsChannel <- jobs
}

func (jabScrapper *jobScrapper) GetPageCount() int {
	response, err := http.Get(jabScrapper.urls.GetBaseUrl())
	if isFail(response, err) {
		return -1
	}

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if isFail(nil, err) {
		return -1
	}

	var result int
	document.Find(PaginationClassName).Each(func(i int, selection *goquery.Selection) {
		result = extractPageCount(selection)
	})
	return result
}

func extractPageCount(selection *goquery.Selection) int {
	return selection.Find(ATag).Length()
}
