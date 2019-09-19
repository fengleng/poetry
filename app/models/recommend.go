/*
@Time : 2019/9/19 17:04
@Author : zxr
@File : Recommend
@Software: GoLand
*/
package models

import (
	"github.com/astaxie/beego/orm"
)

//poetry_recommend 诗文推荐表
type Recommend struct {
	Id          int64 `orm:"column(id);auto"`
	PoetryId    int64 `orm:"column(poetry_id)"`
	Sort        int   `orm:"column(sort)"`
	Status      int   `orm:"column(status)"`
	RecommeTime int64 `orm:"column(recomme_time)"`
	AddDate     int64 `orm:"column(add_date)"`
}

func init() {
	orm.RegisterModel(new(Recommend))
}

func (r *Recommend) TableName() string {
	return RecommendTable
}

func NewRecommendModel() *Recommend {
	return new(Recommend)
}

//获取当天的推荐数据
func (r *Recommend) GetSameDayData(offset, limit int) (data []Recommend, err error) {
	_, err = orm.NewOrm().QueryTable(RecommendTable).Filter("status", 1).OrderBy("-recomme_time", "id").Limit(limit, offset).All(&data, "poetry_id", "sort")
	return
}
