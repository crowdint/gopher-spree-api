package json

import (
	"errors"
)

var (
	responsePaginator    *Paginator
	SpreeResponseFetcher *ResponseInteractor
)

func init() {
	if responsePaginator == nil {
		responsePaginator = new(Paginator)
	}

	if SpreeResponseFetcher == nil {
		SpreeResponseFetcher = new(ResponseInteractor)
	}
}

type ResponseParameters interface {
	GetCurrentPage() (int, error)
	GetPerPage() (int, error)
	GetGransakQuery() (string, error)
}

type ContentInteractor interface {
	GetTotalCount() (int64, error)
	GetResponse(int, int, string) (ContentResponse, error)
}

type ContentResponse interface {
	GetCount() int
	GetData() interface{}
	GetTag() string
}

type ResponseInteractor struct {
	ContentInteractor ContentInteractor
}

func (this *ResponseInteractor) GetResponse(contentInteractor ContentInteractor, params ResponseParameters) (map[string]interface{}, error) {
	this.ContentInteractor = contentInteractor

	currentPage, err := params.GetCurrentPage()
	if err != nil {
		return nil, err
	}

	perPage, err := params.GetPerPage()
	if err != nil {
		return nil, err
	}

	err = responsePaginator.CalculatePaginationData(this.ContentInteractor, currentPage, perPage)
	if err != nil {
		return nil, err
	}

	query, err := params.GetGransakQuery()
	if err != nil {
		return nil, err
	}

	content, err := this.getContent(responsePaginator, query)
	if err != nil {
		return nil, err
	}

	response := this.getResponse(responsePaginator, content)

	return response, nil
}

func (this *ResponseInteractor) getContent(paginator *Paginator, query string) (ContentResponse, error) {
	content, err := this.ContentInteractor.GetResponse(
		paginator.CurrentPage,
		paginator.PerPage,
		query,
	)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (this *ResponseInteractor) getResponse(paginator *Paginator, contentResponse ContentResponse) map[string]interface{} {

	responseMap := map[string]interface{}{
		"count":                  contentResponse.GetCount(),
		"total_count":            paginator.TotalCount,
		"current_page":           paginator.CurrentPage,
		"per_page":               paginator.PerPage,
		"pages":                  paginator.Pages,
		contentResponse.GetTag(): contentResponse.GetData(),
	}

	return responseMap
}

func (this *ResponseInteractor) getError(err error) error {
	return errors.New("Response error: " + err.Error())
}
