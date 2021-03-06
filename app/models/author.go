/*
@Time : 2019/9/20 17:56
@Author : zxr
@File : author
@Software: GoLand
*/
package models

import (
	"github.com/astaxie/beego/orm"
)

type Author struct {
	Id            int64  `orm:"column(id);auto"`
	Author        string `orm:"column(author)"`
	SourceUrl     string `orm:"column(source_url)"`
	WorksUrl      string `orm:"column(works_url)"`
	DynastyId     int    `orm:"column(dynasty_id)"`
	AuthorsId     int    `orm:"column(authors_id)"`
	PhotoUrl      string `orm:"column(photo_url)"`
	PhotoFileName string `orm:"column(photo_file_name)"`
	AuthorIntro   string `orm:"column(author_intro)"`
	PoetryCount   int    `orm:"column(poetry_count)"`
	IsRecommend   int    `orm:"column(is_recommend)"`
	Pinyin        string `orm:"column(pinyin)"`
	Acronym       string `orm:"column(acronym)"`
	AuthorTitle   string `orm:"column(author_title)"`
	AddDate       int64  `orm:"column(add_date)"`
	UpdateDate    int64  `orm:"column(update_date)"`
}

func init() {
	orm.RegisterModel(new(Author))
}

func (a *Author) TableName() string {
	return AuthorTable
}

func NewAuthor() *Author {
	return new(Author)
}

//根据id查询作者信息
func (a *Author) GetAuthorInfoByIds(idList []int64) (data []Author, err error) {
	fields := []string{"id", "author", "dynasty_id", "pinyin", "acronym", "poetry_count", "author_intro", "photo_file_name", "photo_url"}
	_, err = orm.NewOrm().QueryTable(AuthorTable).Filter("id__in", idList).All(&data, fields...)
	return
}

//根据作者名字获取作者资料
func (a *Author) GetAuthorInfoByName(name string) (data Author, err error) {
	fields := []string{"id", "author", "dynasty_id", "pinyin", "acronym", "poetry_count", "author_intro", "photo_file_name", "photo_url"}
	_, err = orm.NewOrm().QueryTable(AuthorTable).Filter("author", name).All(&data, fields...)
	return
}

//根据诗词总数倒序查询作者列表
func (a *Author) GetListByOrderCountDesc(offset, limit int) (data []Author, err error) {
	fields := []string{"id", "author", "poetry_count", "author_intro", "photo_file_name", "photo_url"}
	_, err = orm.NewOrm().QueryTable(AuthorTable).OrderBy("-poetry_count").Limit(limit, offset).All(&data, fields...)
	return
}

//根据朝代ID查询作者列表
func (a *Author) GetListByDynastyId(dynastyId int, offset, limit int) (data []Author, err error) {
	fields := []string{"id", "author", "poetry_count", "author_intro", "photo_file_name", "photo_url"}
	if dynastyId > 0 {
		_, err = orm.NewOrm().QueryTable(AuthorTable).Filter("dynasty_id", dynastyId).OrderBy("-poetry_count").Limit(limit, offset).All(&data, fields...)
	} else {
		_, err = orm.NewOrm().QueryTable(AuthorTable).OrderBy("-poetry_count").Limit(limit, offset).All(&data, fields...)
	}
	return
}

//根据朝代ID查询作者总数
func (a *Author) GetCountByDynastyId(dynastyId int) (count int64, err error) {
	if dynastyId > 0 {
		count, err = orm.NewOrm().QueryTable(AuthorTable).Filter("dynasty_id", dynastyId).Count()
	} else {
		count, err = orm.NewOrm().QueryTable(AuthorTable).Count()
	}
	return
}
