package gin

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
)

type IResponse interface {
	IJson(obj interface{}) IResponse
	IJsonp(obj interface{}) IResponse
	IXml(obj interface{}) IResponse
	IHtml(file string, obj interface{}) IResponse
	IText(format string, values ...interface{}) IResponse
	IRedirect(path string) IResponse
	ISetHeader(key string, val string) IResponse
	ISetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse
	ISetStatus(code int) IResponse
	ISetOkStatus() IResponse
}

func (ctx *Context) IJson(obj interface{}) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		return ctx.ISetStatus(http.StatusInternalServerError)
	}
	ctx.ISetHeader("Content-Type", "application/json")
	ctx.Writer.Write(byt)
	return ctx
}

func (ctx *Context) IJsonp(obj interface{}) IResponse {
	callbackFunc, _ := ctx.DefaultQueryString("callback", "callback_function")
	ctx.ISetHeader("Content-Type", "application/javascript")
	// 输出到前端页面的时候，需要进行字符过滤，防止xss
	callback := template.JSEscapeString(callbackFunc)

	jsonByte, err := json.Marshal(obj)
	if err != nil {
		return ctx
	}

	var ret bytes.Buffer
	ret.WriteString(callback)
	ret.WriteString("(")
	ret.Write(jsonByte)
	ret.WriteString(")")

	_, _ = ctx.Writer.Write(ret.Bytes())
	return ctx
}

func (ctx *Context) IXml(obj interface{}) IResponse {
	byt, err := xml.Marshal(obj)
	if err != nil {
		return ctx.ISetStatus(http.StatusInternalServerError)
	}
	ctx.ISetHeader("Content-Type", "application/xml")
	_, _ = ctx.Writer.Write(byt)
	return ctx
}

func (ctx *Context) IHtml(file string, obj interface{}) IResponse {
	t, err := template.New("output").ParseFiles(file)
	if err != nil {
		return ctx
	}
	ctx.ISetHeader("Content-Type", "application/html")
	_ = t.Execute(ctx.Writer, obj)
	return ctx
}

func (ctx *Context) IText(format string, values ...interface{}) IResponse {
	out := fmt.Sprintf(format, values...)
	ctx.ISetHeader("Content-Type", "application/text")
	ctx.Writer.Write([]byte(out))
	return ctx
}

func (ctx *Context) IRedirect(path string) IResponse {
	http.Redirect(ctx.Writer, ctx.Request, path, http.StatusMovedPermanently)
	return ctx
}

func (ctx *Context) ISetHeader(key string, value string) IResponse {
	ctx.Writer.Header().Add(key, value)
	return ctx
}

func (ctx *Context) ISetCookie(key string, val string, maxAge int, path string, domain string, secure bool, httpOnly bool) IResponse {
	if path == "" {
		path = "/"
	}
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     key,
		Value:    val,
		Path:     path,
		Domain:   domain,
		MaxAge:   maxAge,
		Secure:   secure,
		HttpOnly: httpOnly,
		SameSite: 1,
	})
	return ctx
}

func (ctx *Context) ISetStatus(code int) IResponse {
	ctx.Writer.WriteHeader(code)
	return ctx
}

func (ctx *Context) ISetOkStatus() IResponse {
	ctx.Writer.WriteHeader(http.StatusOK)
	return ctx
}
