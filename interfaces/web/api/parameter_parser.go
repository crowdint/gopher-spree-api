package api

import (
	"net/http"
	"net/url"
	"strconv"

	. "github.com/crowdint/gransak/filter"
)

func NewRequestParameters(r *http.Request) *RequestParameters {
	return &RequestParameters{
		request:  r,
		queryMap: r.URL.Query(),
	}
}

type RequestParameters struct {
	currentPage  int
	perPage      int
	gransakQuery string
	request      *http.Request
	queryMap     url.Values
}

func (this *RequestParameters) GetCurrentPage() (int, error) {
	return getIntParameter(this.queryMap, "page")
}

func (this *RequestParameters) GetPerPage() (int, error) {
	return getIntParameter(this.queryMap, "per_page")
}

func (this *RequestParameters) GetGransakQuery() (string, error) {
	return Gransak.FromRequest(this.request), nil
}

func getIntParameter(queryMap url.Values, key string) (int, error) {
	str := queryMap.Get(key)

	if str == "" {
		return 0, nil
	}

	number, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	if number < 0 {
		return 0, nil
	}

	return number, nil
}
