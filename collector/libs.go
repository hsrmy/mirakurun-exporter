package collector

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ApiInfo struct {
	Host   string
	Port   string
	Schema string
}
type Query struct {
	Query *string
}

func newAPI() ApiInfo {
	apiInfo := new(ApiInfo)

	// Configure Host
	h, e := os.LookupEnv("MIRAKURUN_HOST")
	if !e {
		h = "localhost"
	}
	apiInfo.Host = h

	// Configure Port
	p, e := os.LookupEnv("MIRAKURUN_PORT")
	if !e {
		p = "40772"
	}
	apiInfo.Port = p

	// Configure Schema
	s, e := os.LookupEnv("MIRAKURUN_SCHEMA")
	if !e {
		s = "http"
	}
	apiInfo.Schema = s

	return *apiInfo
}

func getApiRoot(apiInfo *ApiInfo) string {
	return fmt.Sprintf("%s://%s:%s/api/", apiInfo.Schema, apiInfo.Host, apiInfo.Port)
}

func fetch(apiInfo *ApiInfo, namespace string, query *Query) []byte {
	root := getApiRoot(apiInfo)
	url := fmt.Sprintf("%s%s", root, namespace)
	if query.Query != nil {
		url = fmt.Sprintf("%s?%s", url, *query.Query)
	}

	// APIを叩く
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}
