package server

import "net/http"

var textContentType = []string{"text/plain; charset=utf-8"}

//Set http response Content-Type
func setContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

// Query 查询 URL query string
func (srv *Server) Query(key string) string {
	req := srv.Request
	return req.URL.Query().Get(key)
}

// SetResponseWrite 设置回调返回值
func (srv *Server) SetResponseWrite(str string) {
	setContentType(srv.Write, textContentType)
	srv.Write.WriteHeader(http.StatusOK)
	_, err := srv.Write.Write([]byte(str))
	if err != nil {
		panic(err)
	}
}
