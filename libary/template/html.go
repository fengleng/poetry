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

//解析字符串为html
func unescaped(x string) template.HTML {
	return template.HTML(x)
}

//显示模板页
func (h *Html) Display(page string, data interface{}) {
	htmlPath := []string{
		h.ViewPath + "/views/public/header.html",
		h.ViewPath + "/views/public/naver.html",
		h.ViewPath + "/views/public/footer.html",
		h.ViewPath + "/views/" + page,
	}
	if len(h.templateFiles) > 0 {
		h.templateFiles = append(h.templateFiles, htmlPath...)
	} else {
		h.templateFiles = htmlPath
	}
	tpl := template.New(page)
	//切记：加的自定义函数在Parse之前
	tpl = tpl.Funcs(template.FuncMap{"unescaped": unescaped})
	tpl = template.Must(tpl.ParseFiles(
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
	must := template.Must(template.New(page).Funcs(template.FuncMap{"unescaped": unescaped}).ParseFiles(file))
	if err := must.ExecuteTemplate(h.Writer, page, data); err != nil {
		logrus.Infoln("ExecuteTemplate error:", err)
	}
	return
}
