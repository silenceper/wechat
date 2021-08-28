package context

import (
	"bytes"
	"encoding/xml"
	"net/http"
	"io"
)

var xmlContentType = []string{"application/xml; charset=utf-8"}
var plainContentType = []string{"text/plain; charset=utf-8"}

//Render render from bytes
func (ctx *Context) Render(bytes1 []byte) {
	//debug
	//fmt.Println("response msg = ", string(bytes))
	bytesBuffer := bytes.NewBuffer(bytes1)
    io.Copy(ctx.Writer, bytesBuffer)
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
