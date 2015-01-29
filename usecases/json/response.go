package json

import (
	"errors"
	"github.com/crowdint/gopher-spree-api/domain/json"
)

type ResponseInteractor struct {
	ContentInteractor *ProductInteractor
}

func NewResponseInteractor() *ResponseInteractor {
	return &ResponseInteractor{
		ContentInteractor: NewProductInteractor(),
	}
}

func (this *ResponseInteractor) GetResponse() (*json.ProductResponse, error) {
	response := &json.ProductResponse{
		Count:       1,
		Pages:       1,
		CurrentPage: 1,
	}

	content, err := this.ContentInteractor.GetResponse()
	if err != nil {
		return nil, errors.New("Response error: " + err.Error())
	}

	response.Products = content

	return response, nil
}
