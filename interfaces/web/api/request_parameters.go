package api

import (
	"strconv"

	rsp "github.com/crowdint/gopher-spree-api/usecases/json"
	. "github.com/crowdint/gransak"

	"github.com/gin-gonic/gin"
)

func NewRequestParameters(context *gin.Context) *RequestParameters {
	context.Request.ParseForm()

	return &RequestParameters{
		context: context,
	}
}

type RequestParameters struct {
	context *gin.Context
}

func (this *RequestParameters) GetIntParam(key string) (int, error) {
	var param string

	if key == rsp.ID_PARAM {
		param = this.context.Params.ByName("product_id")
	} else {
		param = this.context.Request.Form.Get(key)
	}

	return getInt(param)
}

func (this *RequestParameters) GetStrParam(key string) (string, error) {
	param := this.context.Params.ByName(key)
	if param == "" {
		return this.context.Request.Form.Get(key), nil
	}
	return param, nil
}

func (this *RequestParameters) GetQuery() (*rsp.RequestQuery, error) {
	query, params := Gransak.FromRequest(this.context.Request)

	reqQuery := &rsp.RequestQuery{
		Type:   rsp.GRANSAK,
		Query:  query,
		Params: params,
	}

	return reqQuery, nil
}

func getInt(str string) (int, error) {
	if str == "" {
		return 0, nil
	}

	number, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	if number < 0 {
		return 0, nil
	}

	return number, nil
}
