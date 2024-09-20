package i18n

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"io/ioutil"
	"text/template"
)

// 配置
var defaultLanguage = language.Chinese //默认语种中文

type Bundle struct {
	bundle *i18n.Bundle
}

type translate interface {
	Translate(param *TranslateParam) string
}

func Initialize(dir string) (*Bundle, error) {
	//默认语种中文
	bundle := i18n.NewBundle(defaultLanguage)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	//不同语种配置文件
	fileList, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	//遍历根目录下的所有toml
	for _, c := range fileList {
		bundle.MustLoadMessageFile(dir + c.Name())
	}
	return &Bundle{bundle: bundle}, nil
}

type TranslateParam struct {
	Lang           string           //转换语种
	Accept         string           //浏览器请求语种 Accept-Language
	MessageID      string           //配置文件 MessageID
	TemplateData   interface{}      //参数模板
	PluralCount    interface{}      //单复数需配置
	DefaultMessage *i18n.Message    //默认消息
	Func           template.FuncMap //参数模板函数
}

func (b *Bundle) Translate(param *TranslateParam) string {
	lang := param.Lang
	accept := param.Accept
	localize := i18n.NewLocalizer(b.bundle, lang, accept)
	s, _ := localize.Localize(&i18n.LocalizeConfig{
		MessageID:      param.MessageID,
		TemplateData:   param.TemplateData,
		PluralCount:    param.PluralCount,
		DefaultMessage: param.DefaultMessage,
		Funcs:          param.Func,
	})
	return s
}
