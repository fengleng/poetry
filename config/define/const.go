/*
@Time : 2019/9/23 18:05
@Author : zxr
@File : const
@Software: GoLand
*/
package define

//诗词列表
type ContentAll struct {
	ContentArr []*Content
}

//诗词正文和作者总数据
type Content struct {
	PoetryText
	Author
}

//诗词正文数据
type PoetryText struct {
	Id          int
	Title       string
	Content     string
	AuthorId    int64
	SourceUrl   string
	GenreId     int64
	CreatBackId int64
	Sort        int
}

//作者信息
type Author struct {
	Id            int64
	Author        string
	SourceUrl     string
	WorksUrl      string
	DynastyId     int
	DynastyName   string
	AuthorsId     int
	PhotoUrl      string
	PhotoFileName string
	AuthorIntro   string
	PoetryCount   int
	IsRecommend   int
	Pinyin        string
	Acronym       string
	AuthorTitle   string
}
