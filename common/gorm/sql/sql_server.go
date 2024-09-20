package sql

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
)

func Initialize(dir string) (sql map[string]string, err error) {
	sqlXml, err := readDirSqlXml(dir)
	if err != nil {
		return
	}
	sql = make(map[string]string)
	for _, b := range sqlXml.Sql {
		sql[b.Id] = convStrForXml(b.Script)
	}
	return
}

// 遍历xml文件夹下面所有xml
func readDirSqlXml(dir string) (sqlXml SqlXml, err error) {
	fileList, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Sprintf("读取目录失败，错误:%s", err)
		return
	}
	//遍历根目录下的所有sqlxml
	for _, c := range fileList {
		f, readErr := ioutil.ReadFile(dir + c.Name())
		if readErr != nil {
			fmt.Sprintf("Readload sql xml file has error:%s", readErr)
			return
		}

		readErr = xml.Unmarshal(f, &sqlXml)
		if readErr != nil {
			fmt.Sprintf("reader sql xml has :%s", readErr)
			return
		}
	}
	return
}

func convStrForXml(script string) string {
	script = strings.ReplaceAll(script, "&lt;", "<")
	script = strings.ReplaceAll(script, "&gt;", ">")
	return script
}
