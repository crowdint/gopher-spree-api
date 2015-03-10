package api

import (
	"fmt"
	"net/http"
	"strings"
)

func NewESQueryGenerator(serverUrl string) *ESQueryGenerator {
	return &ESQueryGenerator{
		serverUrl:        serverUrl,
		outgoingTemplate: "%s/%s/%s/_search?q=%s&fields=id",
	}
}

type ESQueryGenerator struct {
	outgoingTemplate string
	params           []string
	serverUrl        string
}

func (this *ESQueryGenerator) Parse(index, itype string, r *http.Request) string {
	this.params = []string{}

	paramMap := r.URL.Query()

	for field, values := range paramMap {
		this.addToParams(field, values)
	}

	strParams := strings.Join(this.params, ",")

	return fmt.Sprintf(this.outgoingTemplate, this.serverUrl, index, itype, strParams)
}

func (this *ESQueryGenerator) addToParams(field string, values []string) {
	if this.isNotValid(field) {
		return
	}

	if len(values) == 1 {
		value := strings.Replace(values[0], " ", "+", -1)
		this.params = append(this.params, field+":"+value)
	} else {
		for _, value := range values {
			value = strings.Replace(value, " ", "+", -1)
			this.params = append(this.params, field+":"+value)
		}
	}
}

func (this *ESQueryGenerator) isNotValid(value string) bool {
	if value == "token" || value == "page" || value == "per_page" {
		return true
	}
	return false
}
