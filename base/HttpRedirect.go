package base

import (
	"net/http"
	"html/template"
)

//http 404 系统重定向
func Page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.tpl").ParseFiles("views/404.tpl")
	data := make(map[string]interface{})
	//data["content"] = "page not found"
	t.Execute(rw, data)
}

//http 401 系统重定向
func Page_note_permission(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("401.tpl").ParseFiles("views/401.tpl")
	data := make(map[string]interface{})
	//data["content"] = "你没有权限访问此页面，请联系超级管理员。或去<a href='/'>开启我的OPMS之旅</a>。"
	t.Execute(rw, data)
}
