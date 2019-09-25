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

type Trans struct {
	Id         int   `orm:"column(id);auto"`
	PoetryId   int   `orm:"column(poetry_id)"`
	TransId    int   `orm:"column(trans_id)"`
	NotesId    int64 `orm:"column(notes_id)"`
	Sort       int   `orm:"column(sort)"`
	AddDate    int64 `orm:"column(add_date)"`
	UpdateDate int64 `orm:"column(update_date)"`
}

func init() {
	orm.RegisterModel(new(Trans))
}

func (t *Trans) TableName() string {
	return TransTable
}

func NewTrans() *Trans {
	return new(Trans)
}

//根据诗词ID和翻译ID查询文本ID
func (t *Trans) FindNotesIdByTransId(poetryId, transId int) (data Trans, err error) {
	_, err = orm.NewOrm().QueryTable(TransTable).Filter("poetry_id", poetryId).Filter("trans_id", transId).All(&data, "id", "notes_id", "sort")
	return
}

//根据诗词ID查询文本ID
func (t *Trans) FindNotesIdByPoetryId(poetryId int) (data []Trans, err error) {
	_, err = orm.NewOrm().QueryTable(TransTable).Filter("poetry_id", poetryId).All(&data, "id", "poetry_id", "notes_id")
	return
}

func (t *Trans) GetOrm() orm.Ormer {
	return orm.NewOrm()
}
