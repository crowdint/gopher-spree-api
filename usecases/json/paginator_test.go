package json

import (
	"strconv"
	"testing"

	"github.com/crowdint/gopher-spree-api/configs"
)

func TestPaginator_Calculate_User_Values(t *testing.T) {
	paginator := new(Paginator)

	err := paginator.CalculatePaginationData(
		new(FakeContentInteractor),
		2,
		10,
		&FakeResponseParameters{},
	)

	if err != nil {
		t.Error("An error has ocurred: ", err.Error())
	}

	if paginator.TotalCount != 16 {
		t.Error("Paginator error: Invalid value for TotalCount")
	}

	if paginator.CurrentPage != 2 {
		t.Error("Paginator error: Invalid value for CurrentPage")
	}

	if paginator.PerPage != 10 {
		t.Errorf("Paginator error: Invalid value for PerPage, Got: %d Want: %d", paginator.PerPage, 10)
	}

	if paginator.Pages != 2 {
		t.Error("Paginator error: Invalid value for Pages")
	}
}

func TestPaginator_Calculate_Default_Values(t *testing.T) {
	paginator := new(Paginator)

	err := paginator.CalculatePaginationData(
		new(FakeContentInteractor),
		0,
		0,
		&FakeResponseParameters{},
	)

	if err != nil {
		t.Error("An error has ocurred: ", err.Error())
	}

	if paginator.TotalCount != 16 {
		t.Error("Paginator error: Invalid value for TotalCount")
	}

	if paginator.CurrentPage != 1 {
		t.Error("Paginator error: Invalid value for CurrentPage")
	}

	if strconv.Itoa(paginator.PerPage) != configs.Get(configs.PER_PAGE) {
		t.Errorf("Paginator error: Invalid value for PerPage, Got: %d Want: %s", paginator.PerPage, configs.PER_PAGE)
	}

	if paginator.Pages < 1 {
		t.Error("Paginator error: Invalid value for Pages")
	}
}

type FakeContentInteractor struct{}

func (this *FakeContentInteractor) GetTotalCount(params ResponseParameters) (int64, error) {
	return 16, nil
}

func (this *FakeContentInteractor) GetResponse(a, b int, params ResponseParameters) (ContentResponse, error) {
	return &ProductResponse{}, nil
}

func (this *FakeContentInteractor) GetShowResponse(a ResponseParameters) (interface{}, error) {
	return struct{}{}, nil
}

type FakeResponseParameters struct{}

func (this *FakeResponseParameters) GetIntParam(a string) (int, error) {
	return 0, nil
}

func (this *FakeResponseParameters) GetStrParam(a string) (string, error) {
	return "", nil
}

func (this *FakeResponseParameters) GetGransakParams() (string, []interface{}, error) {
	return "", []interface{}{}, nil
}
