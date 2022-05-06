package data

import (
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

func ArticleAdd(a *Article) error {
	t := time.Now()
	if a.Id == 0 {
		a.Id = t.UnixMilli()
	}
	a.Created = t
	a.Updated = t
	tx := GetDbCli().Session(&gorm.Session{})
	return tx.Table("article").Create(&a).Error
}

func ArticleDelete(id int64, lang string) error {
	tx := GetDbCli().Session(&gorm.Session{})
	return tx.Table("article").Delete(Article{}, "id = ? and lang = ?", id, lang).Error
}

func ArticleUpdate(a *Article) error {
	tx := GetDbCli().Session(&gorm.Session{})
	return tx.Table("article").Where("id = ? and lang = ?", a.Id, a.Lang).Omit("created").Updates(&a).Error
}

func (a *ArticleQuery) ArticleSearch(adm bool) interface{} {
	var list = make([]Article, 0, a.PageSize)
	tx := GetDbCli().Session(&gorm.Session{}).Table("article").Order("created desc, id")
	if a.Id != 0 {
		tx = tx.Where("id = ?", a.Id)
	}
	if a.Status != 0 {
		tx = tx.Where("status = ?", a.Status)
	}
	if a.CateId != 0 {
		tx = tx.Where("cate_id = ?", a.CateId)
	}
	if a.Lang != "" {
		tx = tx.Where("lang = ?", a.Lang)
	}
	if a.Page > 0 && a.PageSize > 0 {
		tx = tx.Limit(a.PageSize).Offset((a.Page - 1) * a.PageSize)
	}
	if !adm {
		type Result struct {
			Id       int64  `json:"id"`
			Lang     string `json:"lang"`
			CateId   int64  `json:"cate_id"`
			Title    string `json:"title"`
			Summary  string `json:"summary"`
			RichText string `json:"rich_text"`
			Uri      string `json:"uri"`
		}
		var result = make([]Result, 0, a.PageSize)
		tx = tx.Select("id", "lang", "cate_id", "title", "summary", "rich_text", "uri")
		err := tx.Scan(&result).Error
		if err != nil {
			log.Println(err.Error())
		}
		return result
	} else {
		err := tx.Find(&list).Error
		if err != nil {
			log.Println(err.Error())
		}
		return list
	}
}
func (a *ArticleQuery) ArticleCount() int {
	var count int64
	tx := GetDbCli().Session(&gorm.Session{}).Table("article")
	if a.Id != 0 {
		tx = tx.Where("id = ?", a.Id)
	}
	if a.Status != 0 {
		tx = tx.Where("status = ?", a.Status)
	}
	if a.CateId != 0 {
		tx = tx.Where("cate_id = ?", a.CateId)
	}
	if a.Lang != "" {
		tx = tx.Where("lang = ?", a.Lang)
	}
	//if a.Page > 0 && a.PageSize > 0 {
	//	tx = tx.Limit(a.PageSize).Offset((a.Page - 1) * a.PageSize)
	//}
	err := tx.Count(&count).Error
	if err != nil {
		log.Println(err.Error())
	}
	strCount := strconv.FormatInt(count, 10)
	intCount, err := strconv.Atoi(strCount)
	if err != nil {
		log.Println(err.Error())
	}
	return intCount
}
func ArticleMatch(subStr, lang string) (interface{}, int) {
	tx := GetDbCli().Session(&gorm.Session{}).Table("article").Order("created desc, id")
	tx = tx.Where("lang = ?", lang)
	tx = tx.Where("title like ? ", "%"+subStr+"%")
	tx = tx.Limit(30)
	type Result struct {
		Id       int64  `json:"id"`
		Lang     string `json:"lang"`
		CateId   int64  `json:"cate_id"`
		Title    string `json:"title"`
		Summary  string `json:"summary"`
		RichText string `json:"rich_text"`
		Uri      string `json:"uri"`
	}
	var result []Result
	tx = tx.Select("id", "lang", "cate_id", "title", "summary", "rich_text", "uri")
	err := tx.Scan(&result).Error
	if err != nil {
		log.Println(err.Error())
	}
	return result, len(result)

}
