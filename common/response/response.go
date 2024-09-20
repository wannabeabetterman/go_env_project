package response

import (
	"alger/common/global"
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func build(code string, data interface{}, msgId string) *Response {
	return &Response{
		Code:    code,
		Message: msgId,
		Data:    data,
	}

}

// WriteJson writes v as json string into w with code.
func writeJson(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if bs, err := json.Marshal(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if n, err := w.Write(bs); err != nil {
		// http.ErrHandlerTimeout has been handled by http.TimeoutHandler,
		// so it's ignored here.
		if err != http.ErrHandlerTimeout {
			global.Logger.Errorf("write response failed, error: %s", err)
		}
	} else if n < len(bs) {
		global.Logger.Errorf("actual bytes: %d, written bytes: %v", len(bs), n)
	}
}

// WriteJson writes v as json string into w with code.
func writeData(w http.ResponseWriter, code int) {
	//w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(code)
}

const (
	ERROR       = "-1"  //异常
	JWT_INVALID = "-2"  //TOKEN异常
	SUCCESS     = "200" //正常
)

func Ok(w http.ResponseWriter) {
	writeJson(w, http.StatusOK, build(SUCCESS, map[string]interface{}{}, "OkMessage"))
}

func OkWithMessageId(w http.ResponseWriter, messageId string) {
	writeJson(w, http.StatusOK, build(SUCCESS, map[string]interface{}{}, messageId))
}

func OkWithData(w http.ResponseWriter, data interface{}) {
	writeJson(w, http.StatusOK, build(SUCCESS, data, "OkMessage"))
}
func OkWithFileData(w http.ResponseWriter) {
	writeData(w, http.StatusOK)
}
func Fail(w http.ResponseWriter) {
	writeJson(w, http.StatusOK, build(ERROR, map[string]interface{}{}, "FailMessage"))
}

func FailWithMessageId(w http.ResponseWriter, messageId string) {
	writeJson(w, http.StatusOK, build(ERROR, map[string]interface{}{}, messageId))

}

func JwtFailWithMessageId(w http.ResponseWriter, messageId string) {
	writeJson(w, http.StatusOK, build(JWT_INVALID, map[string]interface{}{}, messageId))
}
