package json

import (
	"errors"
	"math"
	"strconv"

	"github.com/crowdint/gopher-spree-api/configs"
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

func (this *ResponseInteractor) GetResponse(currentPage, perPage int) (*json.ProductResponse, error) {
	if currentPage == 0 {
		currentPage = 1
	}

	if perPage == 0 {
		tmp, err := this.getPerPageDefault(10)
		if err != nil {
			return nil, this.getError(err)
		}

		perPage = int(tmp)
	}

	totalCount, err := this.ContentInteractor.GetTotalCount()
	if err != nil {
		return nil, this.getError(err)
	}

	content, err := this.ContentInteractor.GetResponse(currentPage, int(perPage))
	if err != nil {
		return nil, this.getError(err)
	}

	pages := math.Ceil(float64(totalCount) / float64(perPage))

	response := &json.ProductResponse{
		TotalCount:  int(totalCount),
		CurrentPage: currentPage,
		PerPage:     perPage,
		Pages:       int(pages),
		Products:    content,
		Count:       len(content),
	}

	return response, nil
}

func (this *ResponseInteractor) getPerPageDefault(def int64) (int64, error) {
	perPageStr := configs.Get(configs.PER_PAGE)

	if perPageStr == "" {
		return def, nil
	}

	var perPage int64

	temp, err := strconv.Atoi(perPageStr)
	perPage = int64(temp)
	if err != nil {
		return 0, err
	}

	return perPage, nil
}

func (this *ResponseInteractor) getError(err error) error {
	return errors.New("Response error: " + err.Error())
}
