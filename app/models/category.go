package models

import (
	"github.com/astaxie/beego/orm"
)

//诗词分类Model
type Category struct {
	Id             int    `orm:"column(id);auto"`
	CatName        string `orm:"column(cat_name)"`
	SourceUrl      string `orm:"column(source_url)"`
	SourceUrlCrc32 uint32 `orm:"column(source_url_crc32)"`
	ShowPosition   int    `orm:"column(show_position)"`
	Status         int    `orm:"column(status)"`
	Sort           int    `orm:"column(sort)"`
	Pid            int    `orm:"column(pid)"`
}

func init() {
	orm.RegisterModel(new(Category))
}

func (c *Category) TableName() string {
	return CategoryTable
}

func NewCategory() *Category {
	return &Category{}
}

//根据id数组查询数据
func (c *Category) GetDataByIds(id []int) (data []Category, err error) {
	fields := []string{"id", "cat_name", "source_url", "source_url_crc32", "show_position", "pid"}
	_, err = orm.NewOrm().QueryTable(CategoryTable).Filter("id__in", id).Filter("status", 1).All(&data, fields...)
	return
}

//根据show_position查询所有分类
func (c *Category) GetCateByPositionLimit(showPosition, offset, limit int) (data []Category, err error) {
	fields := []string{"id", "cat_name", "source_url", "source_url_crc32"}
	_, err = orm.NewOrm().QueryTable(CategoryTable).Filter("show_position", showPosition).Filter("pid", 0).Filter("status", 1).OrderBy("-sort").Limit(limit, offset).All(&data, fields...)
	return
}

//根据分类名字查询分类信息
func (c *Category) GetCategoryInfoByCateName(catName string, showPosition int) (data Category, err error) {
	fields := []string{"id", "cat_name", "source_url", "source_url_crc32"}
	_, err = orm.NewOrm().QueryTable(CategoryTable).Filter("cat_name", catName).Filter("show_position", showPosition).All(&data, fields...)
	return
}

//查询子分类
func (c *Category) GetAllSubCateData(pid int, showPosition int, offset, limit int) (data []Category, err error) {
	fields := []string{"id", "cat_name", "pid", "source_url", "source_url_crc32"}
	if pid > 0 {
		_, err = orm.NewOrm().QueryTable(CategoryTable).Filter("pid", pid).Filter("show_position", showPosition).Limit(limit, offset).All(&data, fields...)
	} else {
		_, err = orm.NewOrm().QueryTable(CategoryTable).Filter("pid__gt", 0).Filter("show_position", showPosition).Limit(limit, offset).All(&data, fields...)
	}
	return
}

//根据分类名字和PID查询分类信息
func (c *Category) GetCateInfoByNameAndPid(pid int, cateName string) (data Category, err error) {
	fields := []string{"id", "cat_name", "source_url_crc32"}
	_, err = orm.NewOrm().QueryTable(CategoryTable).Filter("pid", pid).Filter("cat_name", cateName).All(&data, fields...)
	return
}
