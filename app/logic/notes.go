/*
@Time : 2019/9/25 16:51
@Author : zxr
@File : notes
@Software: GoLand
*/
package logic

import (
	"poetry/app/models"
	"strings"
)

//诗词译文，注释 内容服务
type NotesLogic struct {
	notesModel *models.Notes
}

func NewNotesLogic() *NotesLogic {
	return &NotesLogic{
		notesModel: models.NewNotes(),
	}
}

//获取内容文本
func (n *NotesLogic) GetNotesBytId(notesId []int) (data []*models.Notes, err error) {
	if len(notesId) == 0 {
		return
	}
	var notesList []models.Notes
	notesList, err = n.notesModel.GetNotesByIds(notesId)
	data = make([]*models.Notes, len(notesList))
	for k, notes := range notesList {
		tmpNotes := notes
		tmpNotes.HtmlSrcUrl = strings.Replace(tmpNotes.HtmlSrcUrl, "shiwen2017", "nocdn", -1)
		data[k] = &tmpNotes
	}
	defer func() {
		notesList = nil
		notesId = nil
	}()
	return
}
