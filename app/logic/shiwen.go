/*
@Time : 2019/9/25 15:55
@Author : zxr
@File : nocde
@Software: GoLand
*/
package logic

import (
	"errors"
	"fmt"
	"poetry/app/models"
	"poetry/config/define"
	"poetry/tools"
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
	NotesAll         = "all"
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
		if transData, err = models.NewTrans().FindNotesIdByPoetryId(poetryId); err != nil || len(transData) == 0 {
			return
		}
	}
	if typeStr == NotesShangXiType {
		if appRecData, err = models.NewAppRec().FindNotesIdByPoetryId(poetryId); err != nil || len(appRecData) == 0 {
			return
		}
	}
	if typeStr == NotesAll {
		transData, _ = models.NewTrans().FindNotesIdByPoetryId(poetryId)
		appRecData, _ = models.NewAppRec().FindNotesIdByPoetryId(poetryId)
	}
	notesIds := n.extractNotesId(transData, appRecData)
	if len(notesIds) > 0 {
		if notesList, err = NewNotesLogic().GetNotesBytId(notesIds); err != nil {
			return
		}
	}
	var notesData []*models.Notes
	//详情内容HTML格式替换处理
	for _, note := range notesList {
		contentReplace := true
		if note.Title == "创作背景" {
			continue
		}
		if len(note.Introd) > 0 {
			note.Introd = tools.PreContentHtml(note.Introd)
			contentReplace = false
		}
		if contentReplace && len(note.Content) > 0 {
			note.Content = tools.PreContentHtml(note.Content)
		}
		notesData = append(notesData, note)
	}
	return notesData, err
}

//用于首页-诗文-》译或赏的操作时，将翻译数据和诗文详情整合成HTML格式字符串，用于页面点击AJAX获取具体内容时用到
func (n *ShiWenLogic) GetNotesContentHtml(notesData *models.Notes, typeStr string) string {
	content := notesData.Content
	idStr := strconv.Itoa(notesData.Id)
	content = tools.DealWithNotes(content, idStr)
	if typeStr == NotesShangXiType {
		return "<div class='hr'></div><strong>赏析<br></strong>" + content
	}
	return "<div class='hr'></div>" + content
}

//将翻译数据和诗文详情整合成HTML格式字符串，用于诗文详情页AJAX获取赏析或翻译信息时
func (n *ShiWenLogic) GetNotesDetailHtml(notes *models.Notes) (html string) {
	content := notes.Content
	idStr := strconv.Itoa(notes.Id)
	if len(content) == 0 {
		return
	}
	content = tools.DealWithNotes(content, idStr)
	html = fmt.Sprintf(`
            <div class="contyishang">
            <div style="height:30px; font-weight:bold; font-size:16px; margin-bottom:10px; clear:both;">
            <h2><span style="float:left;">%s</span></h2>
            <a style="float:left; margin-top:7px; margin-left:5px;" href="javascript:PlayFanyi(%d)">
<img id="speakerimgFanyi%d" src="%s/static/images/speaker.png"   alt="" width="16" height="16"/>
</a>
            </div>
              %s
            </div>
           `, notes.Title, notes.Id, notes.Id, define.CdnStaticDomain, content)
	return html
}

//获取notesId
func (n *ShiWenLogic) extractNotesId(transData []models.Trans, appRecData []models.AppRec) (notesIds []int) {
	if len(transData) == 0 && len(appRecData) == 0 {
		return
	}
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
