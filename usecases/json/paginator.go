package json

import (
	"errors"
	"math"
	"strconv"

	"github.com/crowdint/gopher-spree-api/configs"
)

type Paginator struct {
	TotalCount        int
	CurrentPage       int
	PerPage           int
	Pages             int
	ContentInteractor ContentInteractor
	responseParams    ResponseParameters
}

func (this *Paginator) CalculatePaginationData(contentInteractor ContentInteractor, currentPage, perPage int, params ResponseParameters) error {
	this.ContentInteractor = contentInteractor

	this.responseParams = params

	this.calculateCurrentPage(currentPage)

	calculatedPerPage, err := this.calculatePerPage(perPage)
	if err != nil {
		return err
	}
	calculatedTotalCount, err := this.calculateTotalCount()
	if err != nil {
		return err
	}

	this.calculatePages(calculatedTotalCount, calculatedPerPage)

	return err
}

func (this *Paginator) calculateCurrentPage(currentPage int) {
	if currentPage == 0 {
		this.CurrentPage = 1
	} else {
		this.CurrentPage = currentPage
	}
}

func (this *Paginator) calculatePerPage(perPage int) (int, error) {
	var err error

	if perPage == 0 {
		perPage, err = this.getPerPageDefault(10)
		if err != nil {
			return 0, this.getError(err)
		}
	}

	this.PerPage = perPage

	return perPage, nil
}

func (this *Paginator) getPerPageDefault(def int) (int, error) {
	perPageStr := configs.Get(configs.PER_PAGE)

	if perPageStr == "" {
		return def, nil
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		return 0, err
	}

	return perPage, nil
}

func (this *Paginator) calculateTotalCount() (int, error) {
	totalCount, err := this.ContentInteractor.GetTotalCount(this.responseParams)
	if err != nil {
		return 0, this.getError(err)
	}

	intTotalCount := int(totalCount)

	this.TotalCount = intTotalCount

	return intTotalCount, nil
}

func (this *Paginator) calculatePages(totalCount, perPage int) {
	this.Pages = int(math.Ceil(float64(totalCount) / float64(perPage)))
}

func (this *Paginator) getError(err error) error {
	return errors.New("Paginator error: " + err.Error())
}
