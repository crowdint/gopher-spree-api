package json

import (
	"errors"
)

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

func NewResponseInteractor(contentInteractor ContentInteractor) *ResponseInteractor {
	return &ResponseInteractor{
		ContentInteractor: contentInteractor,
	}
}

func (this *ResponseInteractor) GetResponse(currentPage, perPage int) (map[string]interface{}, error) {
	paginator := new(Paginator)

	err := paginator.Calculate(this.ContentInteractor, currentPage, perPage)
	if err != nil {
		return nil, err
	}

	content, err := this.getContent(paginator)
	if err != nil {
		return nil, err
	}

	response := this.getResponse(paginator, content)

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
