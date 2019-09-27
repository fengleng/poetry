/*
@Time : 2019/9/25 15:55
@Author : zxr
@File : nocde
@Software: GoLand
*/
package logic

import (
	"errors"
	"poetry/app/models"
	"strconv"
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

const (
	NotesShangXiType = "shangxi"
	NotesYiWenType   = "yiwen"
)

//根据诗词URL CRC32 查询具体的翻译或详情内容,只返回一条数据
func (n *ShiWenLogic) GetOneNotesDetailByCrcId(crcId uint32, typeStr string) (notesData *models.Notes, err error) {
	if crcId == 0 {
		return
	}
	var (
		poetryData models.Content
		notesList  []*models.Notes
	)
	if poetryData, err = models.NewContent().GetContentByCrc32Id(crcId); err != nil || poetryData.Id == 0 {
		return
	}
	if notesList, err = n.GetAllNotesByPoetryId(poetryData.Id, typeStr); err != nil || len(notesList) == 0 {
		return
	}
	for _, notes := range notesList {
		if typeStr == NotesShangXiType && strings.Contains(notes.Title, "赏析") == true {
			notesData = notes
			break
		}
		if typeStr == NotesYiWenType && strings.Contains(notes.Title, "注释") == true {
			notesData = notes
			break
		}
	}
	if notesData == nil {
		notesData = notesList[0]
	}
	return notesData, nil
}

//查询诗词所有的赏析信息，注释信息
//typeStr:shangxi,zhushi
func (n *ShiWenLogic) GetAllNotesByPoetryId(poetryId int, typeStr string) (notesList []*models.Notes, err error) {
	var (
		transData  []models.Trans
		appRecData []models.AppRec
	)
	defer func() {
		transData = nil
		appRecData = nil
	}()
	if typeStr == NotesYiWenType {
		if transData, err = models.NewTrans().FindNotesIdByPoetryId(poetryId); err != nil {
			return
		}
	}
	if typeStr == NotesShangXiType {
		if appRecData, err = models.NewAppRec().FindNotesIdByPoetryId(poetryId); err != nil {
			return
		}
	}
	notesIds := n.extractNotesId(transData, appRecData)
	notesList, err = NewNotesLogic().GetNotesBytId(notesIds)
	return
}

//将翻译数据和诗文详情整合成HTML格式字符串，用于页面点击AJAX获取具体内容时用到
func (n *ShiWenLogic) GetNotesContentHtml(notesData *models.Notes, typeStr string) string {
	if typeStr == NotesShangXiType {
		return "<div class='hr'></div><strong>赏析<br></strong>" + notesData.Content
	}
	return "<div class='hr'></div>" + notesData.Content
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

//根据url path 获取CrcId
func (n *ShiWenLogic) GetCrcIdByUrlPath(path string) (crcId uint64, err error) {
	var lastI int
	if lastI = strings.LastIndex(path, "/"); lastI == 0 {
		return 0, errors.New("url params error")
	}
	if crcId, _ = strconv.ParseUint(path[lastI+1:], 10, 64); crcId == 0 {
		return 0, errors.New("url query params error")
	}
	return
}
