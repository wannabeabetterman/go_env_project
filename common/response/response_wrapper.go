package response

import (
	"bytes"
	"net/http"
)

type MyResponseWriter struct {
	Writer http.ResponseWriter
	Buf    *bytes.Buffer
}

func (w *MyResponseWriter) Header() http.Header {
	return w.Writer.Header()
}
func (w *MyResponseWriter) Write(bytes []byte) (int, error) {
	return w.Buf.Write(bytes)
}
func (w *MyResponseWriter) WriteHeader(i int) {
	w.Writer.WriteHeader(i)
}

func NewMyResponseWriter(w http.ResponseWriter) *MyResponseWriter {
	return &MyResponseWriter{
		Writer: w,
		Buf:    &bytes.Buffer{},
	}
}
