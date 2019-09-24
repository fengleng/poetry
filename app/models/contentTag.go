/*
@Time : 2019/9/24 16:45
@Author : zxr
@File : contentTag
@Software: GoLand
*/
package models

import "github.com/astaxie/beego/orm"

type ContentTag struct {
	Id         int   `orm:"column(id);auto"`
	PoetryId   int   `orm:"column(poetry_id)"`
	CategoryId int   `orm:"column(category_id)"`
	UpdateTime int64 `orm:"column(update_time)"`
}

func init() {
	orm.RegisterModel(new(ContentTag))
}

func (d *ContentTag) TableName() string {
	return ContentTagTable
}

func NewContentTag() *ContentTag {
	return new(ContentTag)
}

//根据诗词ID集合查询数据
func (d *ContentTag) GetDataByPoetryId(poetryIds []int64) (data []ContentTag, err error) {
	_, err = orm.NewOrm().QueryTable(ContentTagTable).Filter("poetry_id__in", poetryIds).All(&data, "poetry_id", "category_id")
	return
}
