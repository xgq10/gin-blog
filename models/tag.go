package models

import "github.com/jinzhu/gorm"

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// GetTags gets a list of tags based on paging and constraints
func GetTags(pageNum int, pageSize int, maps interface{}) ([]*Tag, error) {
	var (
		tags []*Tag
		err  error
	)
	if pageNum > 0 && pageSize > 0 {
		err = db.Where(maps).Find(&tags).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&tags).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

// GetTagTotal counts the total number of tags based on the constraint
func GetTagTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Tag{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// ExistTagByName checks if there is a tag with the same name
func ExistTagByName(name string) (bool, error) {
	var tag Tag
	if err := db.Select("id").Where("name=? AND deleted_on=?", name, 0).First(&tag).Error; err != nil {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

// ExistTagById determines whether a Tag exists based on the ID
func ExistTagById(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id=? AND deleted_on=?", id, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

// DeleteTag delete a tag
func DeleteTag(id int) error {
	if err := db.Where("id=?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}
	return nil
}

// EditTag modify a single tag
func EditTag(id int, data interface{}) error {
	if err := db.Model(&Tag{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// AddTag Add a Tag
func AddTag(name string, state int, createdBy string) error {
	tag := Tag{
		Name:      name,
		CreatedBy: createdBy,
		State:     state,
	}
	if err := db.Create(&tag).Error; err != nil {
		return err
	}
	return nil
}

// CleanAllTag Cron定时模拟硬删除
func CleanAllTag() (bool, error) {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{}).Error; err != nil {
		return false, err
	}

	return true, nil
}
