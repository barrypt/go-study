package main

import (
	"JX/gjson"
	"JX/gob"
	"JX/mapa"
	"JX/msgpack"
	"JX/sjson"
	"JX/structa"
	"JX/xml"
)

func main() {
	structa.TestStruct()
	mapa.TestMap()
	xml.XmlUM()
	xml.ProdXML()
	msgpack.MsgPack()
	gob.Gob()
	gjson.GjsonTest()
	sjson.SjsonTest()
}
