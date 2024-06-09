package jobscrapper

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type Job struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

const (
	IdAttribute   = "data-jk"
	TitleClass    = ".title>a"
	LocationClass = ".sjcl"
	SalaryClass   = ".salaryText"
	SummaryClass  = ".summary"
)

func CreateJob(selection *goquery.Selection, channel chan<- *Job) {
	id, _ := selection.Attr(IdAttribute)
	title := getCleanString(selection.Find(TitleClass).Text())
	location := getCleanString(selection.Find(LocationClass).Text())
	salary := getCleanString(selection.Find(SalaryClass).Text())
	summary := getCleanString(selection.Find(SummaryClass).Text())
	channel <- &Job{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}
}

func getCleanString(target string) string {
	trimmed := strings.TrimSpace(target)
	fields := strings.Fields(trimmed)
	return strings.Join(fields, " ")
}
