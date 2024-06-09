package urls

import "strconv"

type Urls []string

var baseUrlsInstance = Urls{
	"https://kr.indeed.com/jobs?q=go&limit=50",
}

const (
	PageQuery = "&start="
	PageSize  = 50
)

func GetInstance() Urls {
	return baseUrlsInstance
}

func (urls *Urls) GetBaseUrl() string {
	return (*urls)[0]
}

func (urls *Urls) GetPageUrl(page int) string {
	baseUrl := urls.GetBaseUrl()
	return baseUrl + PageQuery + strconv.Itoa(page*PageSize)
}
