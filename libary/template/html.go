/*
@Time : 2019/9/23 19:06
@Author : zxr
@File : html
@Software: GoLand
*/
package template

import (
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

type Html struct {
	Writer        http.ResponseWriter
	Request       *http.Request
	templateFiles []string
	ViewPath      string
}

func NewHtml(writer http.ResponseWriter, request *http.Request) *Html {
	return &Html{
		Writer:   writer,
		Request:  request,
		ViewPath: "app",
	}
}

//显示错误页面
func (h *Html) DisplayErrorPage(err error) {
	h.displaySinglePage("error.html", err)
	return
}

//显示404页面
func (h *Html) Display404() {
	h.displaySinglePage("404.html", nil)
	return
}

//追加模板
func (h *Html) AddTemplate(template string) {
	h.templateFiles = append(h.templateFiles, template)
}

//显示模板页
func (h *Html) Display(page string, data interface{}) {
	layout := []string{
		h.ViewPath + "/views/public/header.html",
		h.ViewPath + "/views/public/naver.html",
		h.ViewPath + "/views/public/footer.html",
	}
	h.templateFiles = append(h.templateFiles, layout...)
	h.templateFiles = append(h.templateFiles, h.ViewPath+"/views/"+page)
	tpl := template.Must(template.ParseFiles(
		h.templateFiles...,
	))
	if err := tpl.ExecuteTemplate(h.Writer, page, data); err != nil {
		logrus.Infoln("ExecuteTemplate error:", err)
	}
	return
}

//显示单页面
func (h *Html) displaySinglePage(page string, data interface{}) {
	file := h.ViewPath + "/views/" + page
	must := template.Must(template.ParseFiles(file))
	if err := must.ExecuteTemplate(h.Writer, page, data); err != nil {
		logrus.Infoln("ExecuteTemplate error:", err)
	}
	return
}
