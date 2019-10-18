package models

import (
	"github.com/astaxie/beego/orm"
)

type AncientAuthor struct {
	Id         int    `orm:"column(id);auto"`
	AuthorName string `orm:"column(author_name)"`
	SourceUrl  string `orm:"column(source_url)"`
	AddDate    int64  `orm:"column(add_date)"`
}

func init() {
	orm.RegisterModel(new(AncientAuthor))
}

func (a *AncientAuthor) TableName() string {
	return AncientAuthorTable
}

func NewAncientAuthor() *AncientAuthor {
	return new(AncientAuthor)
}

//根据作者ID查询作者信息
func (a *AncientAuthor) GetAuthorById(id int) (author AncientAuthor, err error) {
	_, err = orm.NewOrm().QueryTable(AncientAuthorTable).Filter("id", id).All(&author, "author_name")
	return
}
