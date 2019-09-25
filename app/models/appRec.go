/*
@Time : 2019/8/28 19:22
@Author : zxr
@File : transModel
@Software: GoLand
*/
package models

import (
	"github.com/astaxie/beego/orm"
)

type AppRec struct {
	Id         int   `orm:"column(id);auto"`
	PoetryId   int   `orm:"column(poetry_id)"`
	ApprecId   int   `orm:"column(apprec_id)"`
	NotesId    int64 `orm:"column(notes_id)"`
	Sort       int   `orm:"column(sort)"`
	AddDate    int64 `orm:"column(add_date)"`
	UpdateDate int64 `orm:"column(update_date)"`
}

func init() {
	orm.RegisterModel(new(AppRec))
}

func (a *AppRec) TableName() string {
	return RecTable
}

func NewAppRec() *AppRec {
	return new(AppRec)
}

//根据诗词ID和赏析ID查询文本ID
func (a *AppRec) FindNotesIdByRecId(poetryId, recId int) (data AppRec, err error) {
	_, err = orm.NewOrm().QueryTable(RecTable).Filter("poetry_id", poetryId).Filter("apprec_id", recId).All(&data, "id", "notes_id")
	return
}

//根据诗词ID查询文本ID
func (a *AppRec) FindNotesIdByPoetryId(poetryId int) (data []AppRec, err error) {
	_, err = orm.NewOrm().QueryTable(RecTable).Filter("poetry_id", poetryId).All(&data, "id", "notes_id")
	return
}

func (a *AppRec) GetOrm() orm.Ormer {
	return orm.NewOrm()
}
