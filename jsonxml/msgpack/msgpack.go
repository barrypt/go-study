package msgpack

import (
	"fmt"
	"github.com/vmihailenco/msgpack"
)

type Dog struct {
	Name string
	Age int
	Type string
}

func MsgPack() {
	var dogs []Dog
	dog:=Dog{"小奶狗",1,"田园犬"}
	dog2:=Dog{"小陆",1,"贵宾"}
	dog3:=Dog{"小杨",1,"二哈"}
	dogs=append(dogs,dog,dog2,dog3)
	//fmt.Println(dogs)
	// 序列化
	res,_:=msgpack.Marshal(&dogs)
	fmt.Println(string(res))

	// 反序列化
	var dogs1 []Dog
	msgpack.Unmarshal(res,&dogs1)
	fmt.Println(dogs1)

}