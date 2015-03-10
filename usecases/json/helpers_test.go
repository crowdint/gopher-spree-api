package json

import (
	"encoding/json"

	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

type DummyResponseParams struct {
	currentPage          int
	perPage              int
	gransakQuery         string
	permittedParamsBytes []byte
}

func (this *DummyResponseParams) GetIntParam(key string) (int, error) {
	if key == PAGE_PARAM {
		return this.currentPage, nil
	}
	return this.perPage, nil
}

func (this *DummyResponseParams) GetStrParam(key string) (string, error) {
	if key == GRANSAK_QUERY_PARAM {
		return this.gransakQuery, nil
	}
	return "", nil
}

func (this *DummyResponseParams) GetQuery() (*RequestQuery, error) {
	return &RequestQuery{}, nil
}

func (this *DummyResponseParams) BindPermittedParams(key string, obj interface{}) bool {
	return json.Unmarshal(this.permittedParamsBytes, obj) == nil
}

func newDummyResponseParams(currentPage, perPage int, gransakQuery string, permittedParamsBytes []byte) *DummyResponseParams {
	return &DummyResponseParams{
		currentPage:          currentPage,
		perPage:              perPage,
		gransakQuery:         gransakQuery,
		permittedParamsBytes: permittedParamsBytes,
	}
}

func ResetDB() {
	repositories.Spree_db.Rollback()
	repositories.Spree_db.Close()
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

func (this *FakeContentInteractor) GetCreateResponse(a ResponseParameters) (interface{}, interface{}, error) {
	return struct{}{}, nil, nil
}
