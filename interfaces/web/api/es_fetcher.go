package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/crowdint/gopher-spree-api/configs"
	"github.com/crowdint/gopher-spree-api/utils"
)

var esfetcher *ESFetcher

func init() {
	if esfetcher == nil {
		esfetcher = &ESFetcher{
			qg: NewESQueryGenerator(configs.Get(configs.ELASTIC_SEARCH_SERVER_URL)),
		}
	}
}

type ESFetcher struct {
	qg *ESQueryGenerator
}

func (this *ESFetcher) GetProducIds(index, itype string, r *http.Request) ([]int64, error) {
	esquery := this.qg.Parse(index, itype, r)

	resBytes, err := this.doRequest(esquery)
	if err != nil {
		utils.LogrusError("GetProducIds", err)

		return []int64{}, err
	}

	esr, err := this.toResponse(resBytes)
	if err != nil {
		return []int64{}, err
	}

	ids := this.extractIds(esr)

	return ids, nil
}

func (this *ESFetcher) doRequest(esquery string) ([]byte, error) {
	req, err := http.NewRequest("GET", esquery, nil)
	if err != nil {
		utils.LogrusError("doRequest", err)

		return []byte{}, err
	}

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		utils.LogrusError("doRequest", err)

		return []byte{}, err
	}

	return ioutil.ReadAll(res.Body)
}

func (this *ESFetcher) toResponse(resBytes []byte) (*ESResponse, error) {
	esr := &ESResponse{}

	err := json.Unmarshal(resBytes, esr)
	if err != nil {
		utils.LogrusError("toResponse", err)

		return nil, err
	}

	return esr, nil
}

func (this *ESFetcher) extractIds(esr *ESResponse) []int64 {
	ids := []int64{}

	for _, product := range esr.Hits.Hits {
		idArray := product.Fields.Ids
		if len(idArray) > 0 {
			ids = append(ids, idArray[0])
		}
	}

	return ids
}

type ESResponse struct {
	Hits Hits `json:"hits"`
}

type Hits struct {
	Hits []Element `json:"hits"`
}

type Element struct {
	Fields Field `json:"fields"`
}

type Field struct {
	Ids []int64 `json:"id"`
}
