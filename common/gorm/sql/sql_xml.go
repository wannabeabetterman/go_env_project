package sql

import "encoding/xml"

type SqlXml struct {
	Mapper xml.Name `xml:"mapper"` //读取xml节点
	Sql    []Sql    `xml:"sql"`    //读取sql标签下到内容
}

type Sql struct {
	Id     string `xml:"id,attr"`   //读取id属性
	Script string `xml:",innerxml"` //读取 <![CDATA[ xxx ]]> 数据
}
