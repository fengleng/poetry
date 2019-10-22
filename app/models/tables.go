/*
@Time : 2019/9/19 17:06
@Author : zxr
@File : tables
@Software: GoLand
*/
package models

const (
	RecommendTable          = "poetry_recommend"                     //诗文推荐表
	ContentTable            = "poetry_content"                       // 诗词正文表
	AuthorTable             = "poetry_author"                        //作者表
	DynastyTable            = "poetry_dynasty"                       //朝代表
	ContentTagTable         = "poetry_detail_category"               //诗词标签表
	CategoryTable           = "poetry_category"                      //诗文分类表
	TransTable              = "poetry_content_trans"                 //诗词详情翻译信息关联表
	RecTable                = "poetry_content_apprec"                //诗词详情赏析信息关联表
	NotesTable              = "poetry_detail_notes"                  //诗词详情内容表
	AncientBookTable        = "poetry_ancient_book"                  //古籍-书名表
	AuthorDataTable         = "poetry_author_data"                   //作者资料信息表
	FamousSentenceTable     = "poetry_famous_sentence"               //名句表
	AncientCatalogueTable   = "poetry_ancient_book_catalog"          //古籍-书名目录表
	AncientClassifyTable    = "poetry_ancient_category"              //古籍-栏目分类表
	AnCataClassTable        = "poetry_ancient_book_catalog_category" //古籍-书名目录分类表
	AncientBookContentTable = "poetry_ancient_book_content"          //古籍-正文内容
	AncientAuthorTable      = "poetry_ancient_author"                //古籍作者表
	PerfectTable            = "poetry_perfect"                       //资料完善表
)
