/*
@Time : 2019/9/17 15:24
@Author : zxr
@File : middler
@Software: GoLand
*/
package middleware

import "net/http"

type Middler interface {
	Before(writer http.ResponseWriter, request *http.Request)
	Run(writer http.ResponseWriter, request *http.Request)
	After(writer http.ResponseWriter, request *http.Request)
}

type MiddleCenter struct {
	MiddleWareList []Middler
}

func NewMiddleCenter() *MiddleCenter {
	return &MiddleCenter{}
}

func (m *MiddleCenter) RegisterMiddleware(middleware Middler) {
	m.MiddleWareList = append(m.MiddleWareList, middleware)
}

func (m *MiddleCenter) DelMiddleware(middleware Middler) {
	for i, middle := range m.MiddleWareList {
		if middle == middleware {
			m.MiddleWareList[i] = nil
		}
	}
}

func (m *MiddleCenter) AllMiddle() []Middler {
	return m.MiddleWareList
}
