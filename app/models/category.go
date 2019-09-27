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
	_, err = orm.NewOrm().QueryTable(CategoryTable).Filter("id__in", id).All(&data, fields...)
	return
}

//根据show_position查询所有分类
func (c *Category) GetCateByPositionLimit(showPosition, offset, limit int) (data []Category, err error) {
	fields := []string{"id", "cat_name", "source_url", "source_url_crc32"}
	_, err = orm.NewOrm().QueryTable(CategoryTable).Filter("show_position", showPosition).Filter("pid", 0).Limit(limit, offset).All(&data, fields...)
	return
}
