package i18n

import (
	"alger/common/context"
	"alger/common/global"
	"alger/common/i18n"

	"alger/common/response"
	"encoding/json"

	"net/http"
)

const (
	Lang        = "accept-language"
	LangZhCn    = "zh-CN" //中文简体
	LangEnUs    = "en-US" //英文
	LangDefault = "zh-CN"
)

func I18n(next http.Handler, serverCtx *context.ApiServiceContext) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mrw := response.NewMyResponseWriter(w)
		next.ServeHTTP(mrw, r)

		var resp response.Response
		err := json.Unmarshal(mrw.Buf.Bytes(), &resp)
		if err != nil {
			return
		}
		lang := r.Header.Get(Lang)
		if lang == "" {
			lang = LangDefault
		}
		msg := global.Bundle.Translate(&i18n.TranslateParam{
			Lang: lang,
			//Accept:      "zh",
			MessageID: resp.Message,
			//PluralCount: 1,
			//TemplateData: map[string]interface{}{
			//    "PluralCount": 1,
			//},
		})
		resp.Message = msg
		bs, _ := json.Marshal(resp)
		mrw.Writer.Write(bs)
	})
}
