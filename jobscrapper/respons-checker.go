package jobscrapper

import "net/http"

func isFail(response *http.Response, err error) bool {
	if response != nil {
		return response.StatusCode >= 400 || err != nil
	}
	return err != nil
}
