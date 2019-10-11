/*
@Time : 2019/9/24 16:45
@Author : zxr
@File : contentTag
@Software: GoLand
*/
package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strconv"
)

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

//根据分类ID查询诗词总数
func (d *ContentTag) GetCountByCategoryId(categoryId int) (count int, err error) {
	var (
		maps  []orm.Params
		mpVal interface{}
		ok    bool
	)
	sql := fmt.Sprintf(`SELECT COUNT(DISTINCT(poetry_id)) AS C FROM %s WHERE category_id=%d `, ContentTagTable, categoryId)
	_, err = orm.NewOrm().Raw(sql).Values(&maps)
	if len(maps) > 0 {
		if mpVal, ok = maps[0]["C"]; !ok {
			return
		}
		if reflect.TypeOf(mpVal).String() == reflect.String.String() {
			c := reflect.ValueOf(mpVal).String()
			count, err = strconv.Atoi(c)
		}
	}
	return
}
