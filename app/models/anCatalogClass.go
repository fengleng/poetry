/*
@Time : 2019/9/4 18:40
@Author : zxr
@File : catalogCategoryModel
@Software: GoLand
*/
package models

import "github.com/astaxie/beego/orm"

//古籍-书名目录分类
type AnCatalogClass struct {
	Id      int64  `orm:"column(id);auto"`
	BookId  int64  `orm:"column(book_id)"`
	CatName string `orm:"column(cat_name)"`
	Sort    int    `orm:"column(sort)"`
	AddDate int64  `orm:"column(add_date)"`
}

func init() {
	orm.RegisterModel(new(AnCatalogClass))
}

func (a *AnCatalogClass) TableName() string {
	return AnCataClassTable
}

func NewAnCatalogClass() *AnCatalogClass {
	return new(AnCatalogClass)
}

//根据书ID查询目录分类列表
func (a *AnCatalogClass) GetClassListByBookId(bookId int) (data []AnCatalogClass, err error) {
	_, err = orm.NewOrm().QueryTable(AnCataClassTable).Filter("book_id", bookId).OrderBy("sort").All(&data, "id", "cat_name")
	return
}

//根据ID查询目录分类信息
func (a *AnCatalogClass) GetClassDataById(id int) (data AnCatalogClass, err error) {
	_, err = orm.NewOrm().QueryTable(AnCataClassTable).Filter("id", id).All(&data, "id", "cat_name")
	return
}
