/*
@Time : 2019/10/10 18:49
@Author : zxr
@File : funs
@Software: GoLand
*/
package logic

import "poetry/app/models"

//根据诗词数据获取作者ID
func ExtractAuthorId(contentList []models.Content) (authorIds []int64) {
	authorIds = make([]int64, len(contentList))
	for k, content := range contentList {
		authorIds[k] = content.AuthorId
	}
	return
}

//根据诗词数据获取诗词ID
func ExtractPoetryId(contentList []models.Content) (poetryIds []int) {
	poetryIds = make([]int, len(contentList))
	for k, content := range contentList {
		poetryIds[k] = content.Id
	}
	return
}

//根据诗词数据获取诗词ID
func ExtractPoetryIdTo64(contentList []models.Content) (poetryIds []int64) {
	poetryIds = make([]int64, len(contentList))
	for k, content := range contentList {
		poetryIds[k] = int64(content.Id)
	}
	return
}

//获取notesId
func ExtractNotesId(transData []models.Trans, appRecData []models.AppRec) (notesIds []int) {
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
