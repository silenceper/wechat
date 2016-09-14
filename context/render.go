package context

import (
	"encoding/xml"
	"net/http"
)

var xmlContentType = []string{"application/xml; charset=utf-8"}
var plainContentType = []string{"text/plain; charset=utf-8"}

//Render render from bytes
func (ctx *Context) Render(bytes []byte) {
	//debug
	//fmt.Println("response msg = ", string(bytes))
	ctx.Writer.WriteHeader(200)
	_, err := ctx.Writer.Write(bytes)
	if err != nil {
		panic(err)
	}
}

//String render from string
func (ctx *Context) String(str string) {
	writeContextType(ctx.Writer, plainContentType)
	ctx.Render([]byte(str))
}

//XML render to xml
func (ctx *Context) XML(obj interface{}) {
	writeContextType(ctx.Writer, xmlContentType)
	bytes, err := xml.Marshal(obj)
	if err != nil {
		panic(err)
	}
	ctx.Render(bytes)
}

func writeContextType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
