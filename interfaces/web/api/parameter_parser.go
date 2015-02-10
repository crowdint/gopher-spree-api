package api

import (
	"net/http"
	"net/url"
	"strconv"

	. "github.com/crowdint/gransak/filter"
)

type ParameterParser struct {
	currentPage  int
	perPage      int
	gransakQuery string
}

func (this *ParameterParser) Parse(r *http.Request) error {
	params := r.URL.Query()

	currentPage, err := getIntParameter(params, "page")
	if err != nil {
		return err
	}

	this.currentPage = currentPage

	perPage, err := getIntParameter(params, "per_page")
	if err != nil {
		return err
	}

	this.perPage = perPage

	this.gransakQuery = Gransak.FromRequest(r)

	return nil
}

func getIntParameter(params url.Values, key string) (int, error) {
	str := params.Get(key)

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
