package elasticsearch

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"net/http"
)

var client *elasticsearch.Client

func InitES(esURL string) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{esURL},
		Username:  "elastic", // 配置账号
		Password:  "123456",  // 配置密码
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // 跳过证书验证
			},
		},
	}

	var err error
	client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	// 测试连接
	res, err := client.Ping()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.IsError() {
		return nil, fmt.Errorf("Error pinging Elasticsearch: %s", res.Status())
	}
	log.Println("Elasticsearch connection successful!")
	return client, nil
}

// CreateIndex 创建索引
func CreateIndex(index string) error {
	res, err := client.Indices.Create(index)
	if err != nil {
		return fmt.Errorf("cannot create index: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}
	log.Printf("Index %s created successfully\n", index)
	return nil
}

// SearchDocument 执行查询并返回结果
func SearchDocument(index string, query map[string]interface{}) (interface{}, error) {
	// 将查询转换为 JSON
	queryJSON, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("error marshaling query: %v", err)
	}

	queryReader := bytes.NewReader(queryJSON)

	// 执行搜索请求
	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(index),
		client.Search.WithBody(queryReader),
		client.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, fmt.Errorf("error executing search: %v", err)
	}
	defer res.Body.Close()

	// 解析响应结果
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error parsing response body: %v", err)
	}

	return result, nil
}
