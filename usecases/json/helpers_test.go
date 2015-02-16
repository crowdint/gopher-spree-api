package json

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

func (this *DummyResponseParams) GetGransakParams() (string, []interface{}, error) {
	return "", []interface{}{}, nil
}

func newDummyResponseParams(currentPage, perPage int, gransakQuery string) *DummyResponseParams {
	return &DummyResponseParams{
		currentPage:  currentPage,
		perPage:      perPage,
		gransakQuery: gransakQuery,
	}
}
