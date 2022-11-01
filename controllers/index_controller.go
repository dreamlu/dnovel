package controllers

import (
	"dnovel/services"
	"dnovel/util/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
	Service services.BookService
}

func (c IndexController) GetClassify(u *gin.Context) {
	results := c.Service.GetClassify()

	u.JSON(http.StatusOK, result.GetSuccess(results))
}

func (c IndexController) GetClassifyInfo(u *gin.Context) {
	classify := u.Request.FormValue("name")
	results := c.Service.GetClassifyInfo(classify)
	u.JSON(http.StatusOK, result.GetSuccess(results))
}

func (c IndexController) GetSearch(u *gin.Context) {
	k := u.Request.FormValue("k")
	results := c.Service.GetSearch(k)
	u.JSON(http.StatusOK, result.GetSuccess(results))
}

func (c IndexController) GetInfo(u *gin.Context) {
	detailURL := u.Request.FormValue("detail_url")
	source := u.Request.FormValue("source")
	info := c.Service.GetInfo(detailURL, source)

	u.JSON(http.StatusOK, result.GetSuccess(info))
}

func (c IndexController) GetChapters(u *gin.Context) {
	detailURL := u.Request.FormValue("detail_url")
	source := u.Request.FormValue("source")
	chapterList := c.Service.GetChapterList(detailURL, source)

	u.JSON(http.StatusOK, result.GetSuccess(chapterList))
}

func (c IndexController) GetRead(u *gin.Context) {
	detailURL := u.Request.FormValue("detail_url") // 可不传
	chapterURL := u.Request.FormValue("chapter_url")
	source := u.Request.FormValue("source")
	content := c.Service.GetRead(detailURL, chapterURL, source)

	u.JSON(http.StatusOK, result.GetSuccess(content))
}
