package structa

import (
	"encoding/json"
	"fmt"
)

//定义一个简单的结构体 Person
type Person struct {
	Name     string
	Age      int
	Birthday string
	Sex      float32
	Hobby    string
}

//写一个 testStruct()结构体的序列化方法
func TestStruct() {
	person := Person{
		Name:     "小崽子",
		Age:      50,
		Birthday: "2019-09-27",
		Sex:      1000.01,
		Hobby:    "泡妞",
	}

	// 将Monster结构体序列化
	data, err := json.Marshal(&person)
	if err != nil {
		fmt.Printf("序列化错误 err is %v", err)
	}
	//输出序列化结果
	fmt.Printf("person序列化后 = %v", string(data))
    //反序列化
	person2 := Person{}
	json.Unmarshal(data,&person2)
	fmt.Println(person2)
}