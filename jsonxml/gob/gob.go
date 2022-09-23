package gob

import (
	 "bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
)

type Fish struct {
	Name string
	Age int
}

func Gob() {
	// 编码
	var fishes=[]Fish{{"金鱼",1},{"鲸鱼",100}}
	var buf =&bytes.Buffer{}
	//var buf =new(bytes.Buffer)
	encoder:=gob.NewEncoder(buf)
	encoder.Encode(fishes)
	fmt.Println(buf)
	// 保存到文件中，buf格式转成字节切片
	ioutil.WriteFile("./gob/fish.gob",buf.Bytes(),0666)

	// 解码
	var fished []Fish
	file,_ := os.Open("./gob/fish.gob")
	decoder:=gob.NewDecoder(file)
	decoder.Decode(&fished)
	fmt.Println(fishes)

}
