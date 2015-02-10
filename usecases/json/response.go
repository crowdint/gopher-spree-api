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
	GetCurrentPage() int
	GetPerPage() int
	GetGransakQuery() string
	GetInteractor() ContentInteractor
}

type ContentInteractor interface {
	GetTotalCount() (int64, error)
	GetResponse(int, int) (ContentResponse, error)
}

type ContentResponse interface {
	GetCount() int
	GetData() interface{}
	GetTag() string
}

type ResponseInteractor struct {
	ContentInteractor ContentInteractor
}

func (this *ResponseInteractor) GetResponse(contentInteractor ContentInteractor, currentPage, perPage int) (map[string]interface{}, error) {
	this.ContentInteractor = contentInteractor

	err := responsePaginator.CalculatePaginationData(this.ContentInteractor, currentPage, perPage)
	if err != nil {
		return nil, err
	}

	content, err := this.getContent(responsePaginator)
	if err != nil {
		return nil, err
	}

	response := this.getResponse(responsePaginator, content)

	return response, nil
}

func (this *ResponseInteractor) getContent(paginator *Paginator) (ContentResponse, error) {
	content, err := this.ContentInteractor.GetResponse(
		paginator.CurrentPage,
		paginator.PerPage,
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
