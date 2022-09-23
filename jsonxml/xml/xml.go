package xml

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Device struct {
	XMLName xml.Name `xml:"devices"`
	Version string   `xml:"version,attr"`
	Host    []Host   `xml:"host"`
	Desc    string   `xml:",innerxml"`
	Comment    string   `xml:",comment"` // 序列化的时候用，生成注释<!--这是注释-->
	Data    string   `xml:",chardata"` // 序列化用，和innerxml相反，

}
type Host struct {
	XMLName  xml.Name `xml:"host"`
	HostName string   `xml:"hostName"`
	HostCode string   `xml:"hostCode"`
	HostDate string   `xml:"hostDate"`
	ID       int      `xml:"id,attr"`
}

func XmlUM() {
	var d Device
	data, _ := ioutil.ReadFile(".\\xml\\device.xml")
	xml.Unmarshal(data, &d)
	fmt.Println(d)
}

func ProdXML(){

	var d Device=Device{Comment: "这是注释",Data:"<host><Name>测试</Name></host>"}
	d.Host = append(d.Host,Host{HostName: "订单服务",HostCode: "1001",HostDate: "2024-09-08",ID: 1} )
	d.Host = append(d.Host,Host{HostName: "数据服务",HostCode: "1002",HostDate: "2024-09-08"} )
	d.Host = append(d.Host,Host{HostName: "商品服务",HostCode: "1003",HostDate: "2024-09-08"} )
	data,err:=xml.Marshal(d)
	if err!=nil {
		fmt.Println("出错：",err)
		return
	}
	fmt.Println(string(data))
	// 1 不带xml头的
	//ioutil.WriteFile("./device2.xml",data,0666)
	// 2 带xml头的
	headByte:=[]byte(xml.Header) // 把xml头转成byte切片
	headByte=append(headByte,data...) // 把数据拼到后面
	ioutil.WriteFile("./xml/device2.xml",headByte,0666)
}