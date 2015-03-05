package api

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func PerformRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	return rec
}

func EqualFromJSONString(expected *string, got string) bool {
	*expected = *expected + "\n"
	return *expected == got
}

func NotEqualFromJSONString(expected *string, got string) bool {
	return !EqualFromJSONString(expected, got)
}

func ResetDB() {
	repositories.Spree_db.Rollback()
	repositories.Spree_db.Close()
}
