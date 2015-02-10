package json

type DummyResponseParams struct {
	currentPage  int
	perPage      int
	gransakQuery string
}

func (this *DummyResponseParams) GetCurrentPage() (int, error) {
	return this.currentPage, nil
}

func (this *DummyResponseParams) GetPerPage() (int, error) {
	return this.perPage, nil
}

func (this *DummyResponseParams) GetGransakQuery() (string, error) {
	return this.gransakQuery, nil
}

func newDummyResponseParams(currentPage, perPage int, gransakQuery string) *DummyResponseParams {
	return &DummyResponseParams{
		currentPage:  currentPage,
		perPage:      perPage,
		gransakQuery: gransakQuery,
	}
}
