/*
@Time : 2019/9/3 18:52
@Author : zxr
@File : ancientBookModel
@Software: GoLand
*/
package models

import (
	"github.com/astaxie/beego/orm"
)

type AncientBook struct {
	Id               int64  `orm:"column(id);auto"`
	CatId            int    `orm:"column(cat_id)"`
	BookTitle        string `orm:"column(book_title)"`
	BookIntroduction string `orm:"column(book_introduction)"`
	LinkUrl          string `orm:"column(link_url)"`
	LinkUrlCrc32     uint32 `orm:"column(link_url_crc32)"`
	SongUrl          string `orm:"column(song_url)"`
	SongFilePath     string `orm:"column(song_file_path)"`
	SongSrcUrl       string `orm:"column(song_src_url)"`
	FamousTotal      int    `orm:"column(famous_total)"`
	CoverChart       string `orm:"column(cover_chart)"`
	CoverChartPath   string `orm:"column(cover_chart_path)"`
	Status           int    `orm:"column(status)"`
	AddDate          int64  `orm:"column(add_date)"`
}

func init() {
	orm.RegisterModel(new(AncientBook))
}

func (a *AncientBook) TableName() string {
	return AncientBookTable
}

func NewAncientBook() *AncientBook {
	return new(AncientBook)
}

//根据标题和分类查询书籍数据
func (a *AncientBook) GetBookByTitleAndCatId(title string, catId int) (data AncientBook, err error) {
	_, err = orm.NewOrm().QueryTable(AncientBookTable).Filter("cat_id", catId).Filter("book_title", title).All(&data, "id")
	return
}

//根据标题和urlcrc32值查询
func (a *AncientBook) GetBookByTitleANDUrlCrc32(title string, urlCrc uint32) (data AncientBook, err error) {
	_, err = orm.NewOrm().QueryTable(AncientBookTable).Filter("title", title).Filter("link_url_crc32", urlCrc).All(&data, "id")
	return
}

//根据偏移量查询古籍-书名列表
func (a *AncientBook) GetBookListByLimit(offset, limit int) (data []AncientBook, err error) {
	fields := []string{"id", "book_title", "link_url_crc32"}
	_, err = orm.NewOrm().QueryTable(AncientBookTable).Filter("status", 1).Limit(limit, offset).All(&data, fields...)
	return
}

//根据分类ID查询书籍列表
func (a *AncientBook) GetBookListLimitByCatId(catId []int, offset, limit int) (data []AncientBook, err error) {
	fields := []string{"id", "book_title", "cat_id", "book_introduction", "link_url_crc32", "song_url", "song_file_path", "famous_total", "cover_chart", "cover_chart_path"}
	if len(catId) > 0 {
		_, err = orm.NewOrm().QueryTable(AncientBookTable).Filter("cat_id__in", catId).Filter("status", 1).Limit(limit, offset).All(&data, fields...)
	} else {
		_, err = orm.NewOrm().QueryTable(AncientBookTable).Limit(limit, offset).Filter("status", 1).All(&data, fields...)
	}
	return
}

//根据分类ID查询书籍总数
func (a *AncientBook) GetBookCountByCatId(catId []int) (num int64, err error) {
	if len(catId) > 0 {
		num, err = orm.NewOrm().QueryTable(AncientBookTable).Filter("status", 1).Filter("cat_id__in", catId).Count()
	} else {
		num, err = orm.NewOrm().QueryTable(AncientBookTable).Filter("status", 1).Count()
	}
	return
}

//根据ID查询书名列表
func (a *AncientBook) GetBookListByIds(bookIds []int) (data []AncientBook, err error) {
	fields := []string{"id", "book_title", "cat_id", "book_introduction", "link_url_crc32", "song_url", "song_file_path", "famous_total", "cover_chart", "cover_chart_path"}
	_, err = orm.NewOrm().QueryTable(AncientBookTable).Filter("id__in", bookIds).All(&data, fields...)
	return
}
