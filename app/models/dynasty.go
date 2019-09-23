package models

import (
	"github.com/astaxie/beego/orm"
)

type Dynasty struct {
	Id          int    `orm:"column(id);auto"`
	DynastyName string `orm:"column(dynasty_name)"`
	AddDate     int64  `orm:"column(add_date)"`
}

func init() {
	orm.RegisterModel(new(Dynasty))
}

func (d *Dynasty) TableName() string {
	return DynastyTable
}

func NewDynasty() *Dynasty {
	return new(Dynasty)
}

//根据朝代ID查询朝代
func (d *Dynasty) GetDataById(id int) (data Dynasty, err error) {
	_, err = orm.NewOrm().QueryTable(DynastyTable).Filter("id", id).All(&data, "dynasty_name")
	return
}

//根据朝代ID slice 查询朝代
func (d *Dynasty) GetDataByIdArr(ids []int) (data []Dynasty, err error) {
	_, err = orm.NewOrm().QueryTable(DynastyTable).Filter("id__in", ids).All(&data, "id", "dynasty_name")
	return
}
