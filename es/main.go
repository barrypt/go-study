package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	_ "log"
	_ "time"

	"github.com/elastic/go-elasticsearch/v7"
)

func Loges(msg string, dateIndex string) {

	addresses := []string{"http://es-dev.test.betawm.com:9201/"}

	config := elasticsearch.Config{

		Addresses: addresses,
	}

	// new client

	es, err := elasticsearch.NewClient(config)
	if err == nil {
		fmt.Println(err, "Error creating the client")
	}
	res, err := es.Info()

	if err == nil {

		fmt.Println(err, "Error getting response")

	}
	// fmt.Println(res.String())

	var buf bytes.Buffer

	query := map[string]interface{}{

		"query": map[string]interface{}{

			"match": map[string]interface{}{

				"companyNameIK": "中银",
			},
		},

		"highlight": map[string]interface{}{

			"pre_tags": []string{"<font color='red'>"},

			"post_tags": []string{"</font>"},

			"fields": map[string]interface{}{

				"companyNameIK": map[string]interface{}{},
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {

		fmt.Println(err, "Error encoding query")

	}

	// Perform the search request.

	res, err = es.Search(

		es.Search.WithContext(context.Background()),

		es.Search.WithIndex("fundfiltering"),

		es.Search.WithBody(&buf),

		es.Search.WithTrackTotalHits(true),

		es.Search.WithPretty(),
	)

	if err != nil {

		fmt.Println(err, "Error getting response")

	}

	defer res.Body.Close()

	// return res.String()

	fmt.Println(res.String())

}

var (
	msg string

	dateIndex string
)

func init() {

	flag.StringVar(&msg, "msg", "检验是否满足xx条件", "输入要匹配的内容")

	flag.StringVar(&dateIndex, "dateIndex", "product", "输入要搜索的索引")

}

func main() {

	flag.Parse()

	Loges(msg, dateIndex)

}

// func main() {
// 	// 创建 Elasticsearch 配置
// 	cfg := elasticsearch.Config{
// 		Addresses: []string{"http://es-dev.test.betawm.com:9201/"}, // Elasticsearch 节点的地址
// 	}

// 	// 创建 Elasticsearch 客户端
// 	typedClient, err := elasticsearch.NewTypedClient(cfg)
// 	if err != nil {
// 		log.Fatalf("Error creating the client: %s", err)
// 	}
// 	inf, err := typedClient.Info().Do(context.Background())
// 	if err == nil {
// 		fmt.Println("info", inf.Version)
// 	}
// 	// settings := map[string]interface{}{
// 	// 	"index.number_of_replicas": 2,
// 	// 	"index.refresh_interval":   "30s",
// 	// 	// 添加其他要更改的设置
// 	// }
// 	// 将 map 转换成 JSON 字符串
// 	// jsonData, err := json.Marshal(settings)
// 	// if err != nil {
// 	// 	fmt.Println("Error marshaling data to JSON:", err)
// 	// 	return
// 	// }

// 	// 执行 Elasticsearch 查询
// 	//  _ = bytes.NewReader(jsonData)

// 	insL := &types.IndexSettings{
// 		NumberOfReplicas: "2",
// 		NumberOfShards:   "2",
// 	}

// 	res, err := typedClient.Indices.PutSettings().Index("45666").Request(insL).Do(context.Background())

// 	if err != nil {
// 		log.Fatalf("Error executing the request: %s", err)
// 	}

// 	// 处理查询结果
// 	fmt.Println(res)
// }
