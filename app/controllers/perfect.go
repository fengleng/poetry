/*
@Time : 2019/10/22 16:07
@Author : zxr
@File : perfect
@Software: GoLand
*/
package controllers

import (
	"net/http"
	"poetry/app/bootstrap"
	"poetry/app/models"
	"poetry/config/define"
	"poetry/libary/template"
	"poetry/tools"
	"strings"
	"time"
)

//完善页面
func Correction(w http.ResponseWriter, r *http.Request) {
	var (
		assign map[string]interface{}
	)
	assign = make(map[string]interface{})
	assign["version"] = define.StaticVersion
	assign["cdnDomain"] = bootstrap.G_Conf.CdnStaticDomain
	assign["webDomain"] = bootstrap.G_Conf.WebDomain
	assign["urlPath"] = ""
	template.NewHtml(w).Display("perfect/correction.html", assign)
}

//完善页面提交处理
func CorrSubmit(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("txttitle")
	content := r.PostFormValue("preview")
	if len(email) == 0 || len(content) == 0 || len(email) > 30 {
		tools.OutputString(w, "<script>alert('非法请求')</script>")
		return
	}
	content = strings.TrimLeft(content, "<br />")
	data := &models.Perfect{
		Email:   email,
		Content: content,
		AddDate: time.Now().Unix(),
	}
	if id, err := models.NewPerfect().Save(data); err != nil || id == 0 {
		tools.OutputString(w, "<script>alert('提交失败');window.location.href='/perfect';</script>")
		return
	}
	tools.OutputString(w, "<script>alert('提交成功');window.location.href='/';</script>")
	return
}
