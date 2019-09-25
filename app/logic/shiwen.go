/*
@Time : 2019/9/25 15:55
@Author : zxr
@File : nocde
@Software: GoLand
*/
package logic

import (
	"poetry/app/models"
	"poetry/tools"
	"strings"
)

//诗词译文，注释服务
type ShiWenLogic struct {
	contentLogic *contentLogic
}

func NewShiWenLogic() *ShiWenLogic {
	return &ShiWenLogic{
		contentLogic: NewContentLogic(),
	}
}

//查询诗词赏析信息，注释信息
//typeStr:shangxi,zhushi
func (n *ShiWenLogic) GetNotesByPoetryCrcId(poetryId int, typeStr string) (notesData *models.Notes, err error) {
	var (
		transData  []models.Trans
		appRecData []models.AppRec
		notesList  []*models.Notes
	)
	defer func() {
		transData = nil
		appRecData = nil
		notesList = nil
	}()
	if typeStr == "zhushi" {
		if transData, err = models.NewTrans().FindNotesIdByPoetryId(poetryId); err != nil {
			return
		}
	}
	if typeStr == "shangxi" {
		if appRecData, err = models.NewAppRec().FindNotesIdByPoetryId(poetryId); err != nil {
			return
		}
	}
	notesIds := n.extractNotesId(transData, appRecData)
	if notesList, err = NewNotesLogic().GetNotesBytId(notesIds); err != nil || len(notesList) == 0 {
		return
	}
	notesData = notesList[0]
	return
}

//将翻译数据和诗文详情整合成HTML格式字符串，用于页面点击AJAX获取具体内容时用到
func (n *ShiWenLogic) GetNotesContentHtml(notesData *models.Notes, poetryData models.Content) string {
	var builder strings.Builder
	content := tools.AddHtmlLabel(poetryData.Content)
	builder.WriteString(content)
	builder.WriteString("<div class='hr'></div><span>" + notesData.Content + "</span>")
	return builder.String()
}

//获取notesId
func (n *ShiWenLogic) extractNotesId(transData []models.Trans, appRecData []models.AppRec) (notesIds []int) {
	for _, trans := range transData {
		notesIds = append(notesIds, int(trans.NotesId))
	}
	for _, appRec := range appRecData {
		notesIds = append(notesIds, int(appRec.NotesId))
	}
	return
}
