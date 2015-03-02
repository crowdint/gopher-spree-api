package api

import (
	"net/http"
	"net/http/httptest"

	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func PerformRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
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
