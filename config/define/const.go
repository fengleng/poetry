/*
@Time : 2019/9/23 18:05
@Author : zxr
@File : const
@Software: GoLand
*/
package define

import "poetry/app/models"

const (
	ProductEnvStr   = "product"
	DevEnvStr       = "test"
	CdnStaticDomain = "http://127.0.0.1:81"
)

//诗词列表
type ContentAll struct {
	ContentArr []*Content
}

//诗词正文和作者总数据
type Content struct {
	PoetryText                      //诗词信息
	PoetryAuthor                    //作者信息
	Tags         []*models.Category //标签分类信息
}

//诗词正文数据
type PoetryText struct {
	PoetryInfo models.Content //诗词信息
	OriContent string         //诗词正文内容
	LinkUrl    string         //诗词详情页链接地址
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
	AuthorInfo     models.Author //作者信息
	DynastyName    string        //朝代名称
	DynastyLinkUrl string        //朝代详情页链接地址
	AuthorLinkUrl  string        //作者详情页链接地址
}
