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
		&DummyResponseParams{},
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
		&DummyResponseParams{},
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
