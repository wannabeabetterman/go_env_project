package recover

import (
	"alger/common/global"
	"net/http"
)

type HttpRecover struct {
}

func NewHttpRecover() *HttpRecover {
	return &HttpRecover{}
}

func (m *HttpRecover) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				global.Logger.Error("recover err", r)
			}
		}()
		next.ServeHTTP(w, r)
	}
}
