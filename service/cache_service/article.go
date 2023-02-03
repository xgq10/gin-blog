package cache_service

import (
	"gin-blog/utils/e"
	"strconv"
	"strings"
)

type Article struct {
	ID    int
	TagID int
	State int

	PageNum  int
	PageSize int
}

func (a *Article) GetArticleKey() string {
	return e.CacheArticle + "_" + strconv.Itoa(a.ID)
}

func (a *Article) GetArticlesKey() string {
	keys := []string{
		e.CacheArticle,
		"LIST",
	}

	if a.PageNum > 0 {
		keys = append(keys, strconv.Itoa(a.PageNum))
	}
	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.PageSize))
	}

	return strings.Join(keys, "_")
}
