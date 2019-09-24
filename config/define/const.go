/*
@Time : 2019/9/23 18:05
@Author : zxr
@File : const
@Software: GoLand
*/
package define

import "poetry/app/models"

//诗词列表
type ContentAll struct {
	ContentArr []*Content
}

//诗词正文和作者总数据
type Content struct {
	PoetryText
	PoetryAuthor
	Tags []*models.Category
}

//诗词正文数据
type PoetryText struct {
	PoetryInfo models.Content
}

//诗词标签
type PoetryTextTag struct {
	Id             int
	Pid            int
	TagName        string
	SourceUrlCrc32 uint32
	ShowPosition   int
}

//作者信息
type PoetryAuthor struct {
	AuthorInfo  models.Author
	DynastyName string
}
