package api

import (
	"strconv"
	"strings"

	"github.com/crowdint/gopher-spree-api/configs"
	rsp "github.com/crowdint/gopher-spree-api/usecases/json"
	"github.com/crowdint/gopher-spree-api/utils"
	. "github.com/crowdint/gransak"

	"github.com/gin-gonic/gin"
)

func NewRequestParameters(context *gin.Context, queryType int) *RequestParameters {
	context.Request.ParseForm()

	return &RequestParameters{
		context:   context,
		queryType: queryType,
	}
}

type RequestParameters struct {
	context   *gin.Context
	queryType int
	Index     string
	Type      string
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
	var query string
	var params []interface{}

	if this.queryType == rsp.GRANSAK {
		query, params = Gransak.FromRequest(this.context.Request)
	} else {
		index := configs.Get(configs.ES_INDEX)
		ptype := configs.Get(configs.ES_PRODUCT_TYPE)

		ids, err := esfetcher.GetProducIds(index, ptype, this.context.Request)
		if err != nil {
			return nil, err
		}

		query, params = this.getParamsFromES(ids)
	}

	reqQuery := &rsp.RequestQuery{
		Type:   this.queryType,
		Query:  query,
		Params: params,
	}

	return reqQuery, nil
}

func (this *RequestParameters) getParamsFromES(ids []int64) (string, []interface{}) {
	var params []interface{}
	var paramPlaceHolders []string
	placeHolder := "?"

	for i, id := range ids {
		params = append(params, id)

		if configs.Get(configs.DB_ENGINE) == "postgres" {
			placeHolder = "$" + strconv.Itoa(i+1)
		}

		paramPlaceHolders = append(paramPlaceHolders, placeHolder)
	}

	return "id IN (" + strings.Join(paramPlaceHolders, ",") + ")", params
}

func (this *RequestParameters) BindPermittedParams(key string, obj interface{}) bool {
	return this.context.Bind(obj)
}

func getInt(str string) (int, error) {
	if str == "" {
		return 0, nil
	}

	number, err := strconv.Atoi(str)
	if err != nil {
		utils.LogrusError("GetInt", err)
		return 0, err
	}

	if number < 0 {
		return 0, nil
	}

	return number, nil
}
