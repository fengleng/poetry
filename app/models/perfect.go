/*
@Time : 2019/10/22 17:07
@Author : zxr
@File : perfect
@Software: GoLand
*/
package models

import "github.com/astaxie/beego/orm"

//资料完善Model
type Perfect struct {
	Id      int64  `orm:"column(id);auto"`
	Email   string `orm:"column(email)"`
	Content string `orm:"column(content)"`
	AddDate int64  `orm:"column(add_date)"`
}

func init() {
	orm.RegisterModel(new(Perfect))
}

func (p *Perfect) TableName() string {
	return PerfectTable
}

func NewPerfect() *Perfect {
	return new(Perfect)
}

//保存数据
func (p *Perfect) Save(data *Perfect) (id int64, err error) {
	id, err = orm.NewOrm().Insert(data)
	return
}
