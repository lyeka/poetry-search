package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"

	"poetry_search/internal/model"
)

const (
	IndexName = "poetry_v2"
)

const mappingTpl = `{
	"mapping": {
		"properties": {
			"id": {"type": "long"},
			"title": {"type": "keyword"},
			"author": {"type": "keyword"},
			"dynasty": {"type": "keyword"},
			"content": {"type": "text"}
		}
	}
}`

func main() {
	es, _ := newESClient()

	initIndex(es)
	err := dealSinglePoetryCSVData("./poetry_data/宋_1.csv", es)
	if err != nil {
		log.Println(err)
	}
}

func newESClient() (*elasticsearch.Client, error) {
	return elasticsearch.NewClient(
		elasticsearch.Config{
			Addresses: []string{"http://localhost:9200"},
		},
	)
}

func initIndex(es *elasticsearch.Client) {
	res, err := es.Indices.Get([]string{IndexName})
	if err != nil {
		log.Println("query index failed", err)
		return
	}
	if res.StatusCode == 404 {
		_, err = es.Indices.Create(IndexName, es.Indices.Create.WithBody(strings.NewReader(mappingTpl)))
		if err != nil {
			log.Println("create index failed", err)
			return
		}
		log.Println("create index succeeded")
	}

}

func dealSinglePoetryCSVData(fileName string, esClient *elasticsearch.Client) error {
	if esClient == nil {
		return errors.New("nil es client")
	}
	cntb, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	ss, _ := r2.ReadAll()
	sz := len(ss)

	for i := 1; i < sz; i++ {
		poetry := model.Poetry{
			Title:   ss[i][0],
			Dynasty: ss[i][1],
			Author:  ss[i][2],
			Content: ss[i][3],
		}
		fmt.Println(poetry)

		b, err := json.Marshal(poetry)
		if err != nil {
			return err
		}

		req := esapi.IndexRequest{
			Index:   IndexName,
			Body:    strings.NewReader(string(b)),
			Refresh: "true",
		}

		res, err := req.Do(context.Background(), esClient)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.IsError() {
			return fmt.Errorf("status:%s", res.Status())
		} else {
			fmt.Println("success, status:", res.Status())
		}
	}

	return nil
}
