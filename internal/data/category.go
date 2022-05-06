package data

import (
	"gorm.io/gorm"
	"log"
	"time"
)

func CategoryAdd(c *Category) error {
	t := time.Now()
	if c.Id == 0 {
		c.Id = t.UnixMilli()
	}
	tx := GetDbCli().Session(&gorm.Session{}).Table("category")
	return tx.Create(&c).Error
}

func CategoryDelete(id int64, lang string) error {
	tx := GetDbCli().Session(&gorm.Session{}).Table("category")
	return tx.Delete(Category{}, "id = ? and lang = ?", id, lang).Error
}

func CategoryUpdate(c *Category) error {
	tx := GetDbCli().Session(&gorm.Session{}).Table("category")
	return tx.Where("id = ? and lang = ?", c.Id, c.Lang).Updates(&c).Error
}

func (c *CategoryQuery) CategorySearch() interface{} {
	var list = make([]Category, 0, 20)
	tx := GetDbCli().Session(&gorm.Session{}).Table("category").Order("id desc, lang")
	if c.Id != 0 {
		tx = tx.Where("id = ?", c.Id)
	}
	if c.Lang != "" {
		tx = tx.Where("lang = ?", c.Lang)
	}
	if c.ParentId != 0 {
		tx = tx.Where("parent_id = ?", c.ParentId)
	}
	//if c.Page > 0 && c.PageSize > 0 {
	//	tx = tx.Limit(c.PageSize).Offset((c.Page - 1) * c.PageSize)
	//}
	err := tx.Find(&list).Error
	if err != nil {
		log.Println(err.Error())
	}
	return list
}

func CategoryCount(pid int64, lang string) int64 {
	var count int64
	tx := GetDbCli().Session(&gorm.Session{}).Table("category")
	if lang != "" {
		tx = tx.Where("lang = ?", lang)
	}
	if pid != 0 {
		tx = tx.Where("parent_id = ?", pid)
	}
	err := tx.Count(&count).Error
	if err != nil {
		log.Println(err.Error())
	}
	return count
}
