package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestParameterParser(t *testing.T) {
	url := &url.URL{
		RawQuery: "q[name_eq]=name&q[last_name_eq]=lastName",
	}

	request := &http.Request{
		URL: url,
	}

	context := &gin.Context{
		Request: request,
	}

	params := NewRequestParameters(context, 0)

	reqQuery, err := params.GetQuery()
	if err != nil {
		t.Error("An error has ocurred:", err)
	}

	query := reqQuery.Query

	gparams := reqQuery.Params

	expected := "name "

	if !strings.Contains(query, expected) {
		t.Errorf("Mismatch, string: %s does not contain: %s", query, expected)
	}

	expected = "last_name "

	if !strings.Contains(query, expected) {
		t.Errorf("Mismatch, string: %s does not contain: %s", query, expected)
	}

	gparamsStr := fmt.Sprintf("%v", gparams)
	expected = "name"

	if !strings.Contains(gparamsStr, expected) {
		t.Errorf("Mismatch, string: %s does not contain: %s", gparamsStr, expected)
	}

	expected = "lastName"

	if !strings.Contains(gparamsStr, expected) {
		t.Errorf("Mismatch, string: %s does not contain: %s", gparamsStr, expected)
	}
}
