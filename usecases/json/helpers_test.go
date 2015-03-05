package json

import "github.com/crowdint/gopher-spree-api/interfaces/repositories"

type DummyResponseParams struct {
	currentPage  int
	perPage      int
	gransakQuery string
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

func (this *DummyResponseParams) GetPermittedParams(key string, obj interface{}) bool {
	return false
}

func newDummyResponseParams(currentPage, perPage int, gransakQuery string) *DummyResponseParams {
	return &DummyResponseParams{
		currentPage:  currentPage,
		perPage:      perPage,
		gransakQuery: gransakQuery,
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
