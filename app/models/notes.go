package models

import "github.com/astaxie/beego/orm"

type Notes struct {
	Id         int    `orm:"column(id);auto"`
	Title      string `orm:"column(title)"`
	Content    string `orm:"column(content)"`
	PlayUrl    string `orm:"column(play_url)"`
	PlaySrcUrl string `orm:"column(play_src_url)"`
	HtmlSrcUrl string `orm:"column(html_src_url)"`
	Type       int    `orm:"column(type)"`
	Introd     string `orm:"column(introd)"`
	FileName   string `orm:"column(file_name)"`
	AddDate    int64  `orm:"column(add_date)"`
	UpdateDate int64  `orm:"column(update_date)"`
}

func init() {
	orm.RegisterModel(new(Notes))
}

func NewNotes() *Notes {
	return new(Notes)
}

func (n *Notes) TableName() string {
	return NotesTable
}

var NotesFields = []string{"id", "title", "introd", "content", "file_name", "html_src_url", "play_src_url", "play_url"}

//根据id查询notes内容
func (n *Notes) GetNotesById(id int) (data Notes, err error) {
	_, err = orm.NewOrm().QueryTable(NotesTable).Filter("id", id).All(&data, NotesFields...)
	return
}

//根据id数组批量查询notes内容
func (n *Notes) GetNotesByIds(ids []int) (data []Notes, err error) {
	if len(ids) == 1 {
		_, err = orm.NewOrm().QueryTable(NotesTable).Filter("id", ids).All(&data, NotesFields...)
	} else {
		_, err = orm.NewOrm().QueryTable(NotesTable).Filter("id__in", ids).OrderBy("id").All(&data, NotesFields...)
	}
	return
}
