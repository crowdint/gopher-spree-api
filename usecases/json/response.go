package json

import (
	"errors"
)

const (
	PAGE_PARAM          = "page"
	PER_PAGE_PARAM      = "per_page"
	GRANSAK_QUERY_PARAM = "gransak"
	ID_PARAM            = "id"
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
	GetIntParam(string) (int, error)
	GetStrParam(string) (string, error)
}

type ContentInteractor interface {
	GetTotalCount(string) (int64, error)
	GetResponse(int, int, string) (ContentResponse, error)
	GetShowResponse(interface{}) (interface{}, error)
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

	currentPage, err := params.GetIntParam(PAGE_PARAM)
	if err != nil {
		return nil, err
	}

	perPage, err := params.GetIntParam(PER_PAGE_PARAM)
	if err != nil {
		return nil, err
	}

	query, err := params.GetStrParam(GRANSAK_QUERY_PARAM)
	if err != nil {
		return nil, err
	}

	err = responsePaginator.CalculatePaginationData(this.ContentInteractor, currentPage, perPage, query)
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

func (this *ResponseInteractor) GetShowResponse(contentInteractor ContentInteractor, params ResponseParameters) (interface{}, error) {
	id, err := params.GetIntParam(ID_PARAM)
	if err != nil {
		return struct{}{}, err
	}

	int64Id := int64(id)

	return this.ContentInteractor.GetShowResponse(int64Id)
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
