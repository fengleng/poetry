/*
@Time : 2019/9/2 15:51
@Author : zxr
@File : ancientCategoryModel
@Software: GoLand
*/
package models

import "github.com/astaxie/beego/orm"

//古籍-栏目分类表
type AncientClassify struct {
	Id      int    `orm:"column(id);auto"`
	CatName string `orm:"column(cat_name)"`
	SrcUrl  string `orm:"column(src_url)"`
	Pid     int64  `orm:"column(pid)"`
	Sort    int    `orm:"column(sort)"`
}

var ancientFields = []string{"id", "cat_name", "pid", "sort"}

func init() {
	orm.RegisterModel(new(AncientClassify))
}

func (a *AncientClassify) TableName() string {
	return AncientClassifyTable
}

func NewAncientClassify() *AncientClassify {
	return new(AncientClassify)
}

//根据PID查询分类数据并分页
func (a *AncientClassify) GetDataLimitByPid(pid int, offset, limit int) (data []AncientClassify, err error) {
	_, err = orm.NewOrm().QueryTable(AncientClassifyTable).Filter("pid", pid).OrderBy("sort").Limit(limit, offset).All(&data, ancientFields...)
	return
}

//根据分类名称查询数据
func (a *AncientClassify) GetCategoryDataByName(catName string) (data AncientClassify, err error) {
	_, err = orm.NewOrm().QueryTable(AncientClassifyTable).Filter("cat_name", catName).All(&data, "id", "pid", "sort")
	return
}
