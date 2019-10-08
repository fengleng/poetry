/*
@Time : 2019/9/20 17:56
@Author : zxr
@File : author
@Software: GoLand
*/
package models

import (
	"github.com/astaxie/beego/orm"
)

//作者资料信息
type AuthorData struct {
	Id         int   `orm:"column(id);auto"`
	AuthorId   int64 `orm:"column(author_id)"`
	DataId     int   `orm:"column(data_id)"`
	NotesId    int   `orm:"column(notes_id)"`
	Sort       int   `orm:"column(sort)"`
	AddDate    int64 `orm:"column(add_date)"`
	UpdateDate int64 `orm:"column(update_date)"`
}

func init() {
	orm.RegisterModel(new(AuthorData))
}

func (a *AuthorData) TableName() string {
	return AuthorDataTable
}

func NewAuthorData() *AuthorData {
	return new(AuthorData)
}

//根据作者id查询作者详情资料信息
func (a *AuthorData) GetAuthorDetailDataById(id int) (data []AuthorData, err error) {
	fields := []string{"author_id", "data_id", "notes_id", "sort"}
	_, err = orm.NewOrm().QueryTable(AuthorDataTable).Filter("author_id", id).OrderBy("id").All(&data, fields...)
	return
}
