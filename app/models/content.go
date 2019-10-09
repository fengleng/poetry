/*
@Time : 2019/9/20 15:42
@Author : zxr
@File : content
@Software: GoLand
*/
package models

import "github.com/astaxie/beego/orm"

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
