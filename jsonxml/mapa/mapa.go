package mapa

import (
	"encoding/json"
	"fmt"
)

func TestMap() {
	//定义一个map
	a := make(map[string]interface{})

	//使用map之前 必须make一下
	a["name"] = "小崽子"
	a["age"] = 8
	a["address"] = "上海市浦东新区"

	// 将a map结构体序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err is %v", err)
	}
	//输出序列化结果
	fmt.Printf("map序列化后 = %v", string(data))
	//反序列化
	var a1 map[string]interface{}
	json.Unmarshal(data, &a1)
	fmt.Println(a1)
}
