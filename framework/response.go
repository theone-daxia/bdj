package framework

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
)

// 方法返回接口本身，可让使用者进行链式调用

type IResponse interface {
	Json(obj interface{}) IResponse
	Jsonp(obj interface{}) IResponse
	Xml(obj interface{}) IResponse
	Html(file string, obj interface{}) IResponse
	Text(format string, values ...interface{}) IResponse
	Redirect(path string) IResponse
	SetHeader(key string, val string) IResponse
	SetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse
	SetStatus(code int) IResponse
	SetOkStatus() IResponse
}

func (ctx *Context) Json(obj interface{}) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}
	ctx.SetHeader("Content-Type", "application/json")
	ctx.responseWriter.Write(byt)
	return ctx
}

func (ctx *Context) Jsonp(obj interface{}) IResponse {
	callbackFunc, _ := ctx.QueryString("callback", "callback_function")
	ctx.SetHeader("Content-Type", "application/javascript")
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

	_, _ = ctx.responseWriter.Write(ret.Bytes())
	return ctx
}

func (ctx *Context) Xml(obj interface{}) IResponse {
	byt, err := xml.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}
	ctx.SetHeader("Content-Type", "application/xml")
	_, _ = ctx.responseWriter.Write(byt)
	return ctx
}

func (ctx *Context) Html(file string, obj interface{}) IResponse {
	t, err := template.New("output").ParseFiles(file)
	if err != nil {
		return ctx
	}
	ctx.SetHeader("Content-Type", "application/html")
	_ = t.Execute(ctx.responseWriter, obj)
	return ctx
}

func (ctx *Context) Text(format string, values ...interface{}) IResponse {
	out := fmt.Sprintf(format, values...)
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.responseWriter.Write([]byte(out))
	return ctx
}

func (ctx *Context) Redirect(path string) IResponse {
	http.Redirect(ctx.responseWriter, ctx.request, path, http.StatusMovedPermanently)
	return ctx
}

func (ctx *Context) SetHeader(key string, value string) IResponse {
	ctx.responseWriter.Header().Add(key, value)
	return ctx
}

func (ctx *Context) SetCookie(key string, val string, maxAge int, path string, domain string, secure bool, httpOnly bool) IResponse {
	if path == "" {
		path = "/"
	}
	http.SetCookie(ctx.responseWriter, &http.Cookie{
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

func (ctx *Context) SetStatus(code int) IResponse {
	ctx.responseWriter.WriteHeader(code)
	return ctx
}

func (ctx *Context) SetOkStatus() IResponse {
	ctx.responseWriter.WriteHeader(http.StatusOK)
	return ctx
}
