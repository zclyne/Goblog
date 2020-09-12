package service

import (
	"goblog/model"
	"log"
)

func ListTags() []model.Tag {
	var tags []model.Tag
	db.Find(&tags)
	return tags
}

func GetTagById(tagId string) model.Tag {
	var tag model.Tag
	db.Where("tag_id = ?", tagId).First(&tag)
	return tag
}

func ListTagsByBlogId(blogId string) []model.Tag {
	var tags []model.Tag
	rows, err := db.Table("blog_tag").Where("blog_id = ?", blogId).Select("tag_id").Rows()
	if err != nil {
		log.Printf("error occurred while listing tags by blog id, err is: %s\n", err)
		return nil
	}
	for rows.Next() {
		var tagId string
		rows.Scan(&tagId)
		tag := GetTagById(tagId)
		tags = append(tags, tag)
	}
	return tags
}

func CreateTag(t *model.Tag) {
	db.Create(t)
}

func DeleteTag(tagId string) error {
	if err := db.Where("tag_id = ?", tagId).Delete(model.Tag{}).Error; err != nil {
		log.Printf("error occurred while deleting tag, err is: %s\n", err)
		return err
	}
	return nil
}

func UpdateTag(t *model.Tag) error {
	if err := db.Save(t).Error; err != nil {
		log.Printf("error occurred while updating tag: %s\n", err)
		return err
	}
	return nil
}
