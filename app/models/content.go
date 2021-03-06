/*
@Time : 2019/9/20 15:42
@Author : zxr
@File : content
@Software: GoLand
*/
package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strconv"
)

//poetry_content 诗词表
type Content struct {
	Id             int    `orm:"column(id);auto"`
	Title          string `orm:"column(title)"`
	Content        string `orm:"column(content)"`
	AuthorId       int64  `orm:"column(author_id)"`
	SourceUrl      string `orm:"column(source_url)"`
	SourceUrlCrc32 uint32 `orm:"column(sourceurl_crc32)"`
	GenreId        int64  `orm:"column(genre_id)"`
	CreatBackId    int64  `orm:"column(creat_back_id)"`
	Sort           int    `orm:"column(sort)"`
	AddDate        int64  `orm:"column(add_date)"`
	UpdateDate     int64  `orm:"column(update_date)"`
}

func init() {
	orm.RegisterModel(new(Content))
}

func NewContent() *Content {
	return new(Content)
}

func (c *Content) TableName() string {
	return ContentTable
}

var Fields = []string{"id", "title", "content", "author_id", "source_url", "sourceurl_crc32", "genre_id", "creat_back_id"}

//根据诗词ID查询正文数据
func (c *Content) GetContentByIdList(poetryId []int64) (data []Content, err error) {
	_, err = orm.NewOrm().QueryTable(ContentTable).Filter("id__in", poetryId).All(&data, Fields...)
	return
}

//根据sourceurl_crc32 查询正文数据
func (c *Content) GetContentByCrc32Id(crc32Id uint32) (data Content, err error) {
	fields := []string{"id"}
	_, err = orm.NewOrm().QueryTable(ContentTable).Filter("sourceurl_crc32", crc32Id).All(&data, fields...)
	return
}

//根据sourceurl_crc32数组 批量查询正文数据
func (c *Content) GetContentByCrc32IdArr(crc32Ids []uint32) (data []Content, err error) {
	fields := []string{"id", "title", "sourceurl_crc32", "author_id"}
	_, err = orm.NewOrm().QueryTable(ContentTable).Filter("sourceurl_crc32__in", crc32Ids).All(&data, fields...)
	return
}

//根据作者ID查询作者诗词列表
func (c *Content) GetContentListByAuthorId(authorId int64, offset, limit int, orderFiled string) (data []Content, err error) {
	_, err = orm.NewOrm().QueryTable(ContentTable).Filter("author_id", authorId).Limit(limit, offset).OrderBy(orderFiled).All(&data, Fields...)
	return
}

//根据作者ID查询作者诗词总数
func (c *Content) GetContentCountByAuthorId(authorId int64) (count int64, err error) {
	count, err = orm.NewOrm().QueryTable(ContentTable).Filter("author_id", authorId).Count()
	return
}

//根据朝代ID查询诗词列表
func (c *Content) GetContentListByDynastyId(dynastyId int, offset, limit int) (data []Content, err error) {
	sql := fmt.Sprintf(`SELECT a.id,a.title,a.content,a.source_url,a.sourceurl_crc32,a.author_id,b.author,b.dynasty_id FROM %s AS a left join %s AS b on a.author_id=b.id WHERE b.dynasty_id=%d Limit %d,%d`, ContentTable, AuthorTable, dynastyId, offset, limit)
	_, err = orm.NewOrm().Raw(sql).QueryRows(&data)
	return
}

//根据朝代ID查询诗词总数
func (c *Content) GetContentCountByDynastyId(dynastyId int) (count int, err error) {
	var (
		maps  []orm.Params
		mpVal interface{}
		ok    bool
	)
	sql := fmt.Sprintf(`SELECT  COUNT(*) AS C FROM %s AS a left join %s AS b on a.author_id=b.id WHERE b.dynasty_id=%d`, ContentTable, AuthorTable, dynastyId)
	_, err = orm.NewOrm().Raw(sql).Values(&maps)
	if len(maps) > 0 {
		if mpVal, ok = maps[0]["C"]; !ok {
			return
		}
		if reflect.TypeOf(mpVal).String() == reflect.String.String() {
			c := reflect.ValueOf(mpVal).String()
			count, err = strconv.Atoi(c)
		}
	}
	return
}

//根据分类ID查询诗词列表
func (c *Content) GetContentListByCategoryId(categoryId int, offset, limit int) (data []Content, err error) {
	sql := fmt.Sprintf("SELECT a.id,a.title,a.content,a.source_url,a.sourceurl_crc32,a.author_id FROM %s AS a  WHERE a.id IN (SELECT DISTINCT(poetry_id)  FROM %s WHERE category_id=%d) Limit %d,%d", ContentTable, ContentTagTable, categoryId, offset, limit)
	_, err = orm.NewOrm().Raw(sql).QueryRows(&data)
	return
}
