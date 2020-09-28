package service

import (
	"goblog/model"
	"log"

	"github.com/google/uuid"
)

func ListBlogs(pageOptions model.PageOptions) []model.Blog {
	limit, offset := pageOptions.Limit, pageOptions.Offset
	var blogs []model.Blog
	db.Limit(limit).Offset(offset).Order("create_at desc").Find(&blogs)
	return blogs
}

func GetBlogsCount() int {
	var count int
	db.Model(&model.Blog{}).Count(&count)
	return count
}

func ListBlogsByTypeId(pageOptions model.PageOptions, typeId string) []model.Blog {
	limit, offset := pageOptions.Limit, pageOptions.Offset
	var blogs []model.Blog
	db.Limit(limit).Offset(offset).Where("type_id = ?", typeId).Find(&blogs)
	return blogs
}

func ListBlogsByTagId(pageOptions model.PageOptions, tagId string) []model.Blog {
	limit, offset := pageOptions.Limit, pageOptions.Offset
	var blogs []model.Blog
	rows, err := db.Table("blog_tag").Limit(limit).Offset(offset).Where("tag_id = ?", tagId).Select("blog_id").Rows()
	if err != nil {
		log.Printf("error occurred while listing blogs by tag id, err is %s\n", err)
	}
	for rows.Next() {
		var blogId string
		rows.Scan(&blogId)
		blog := GetBlogById(blogId)
		blogs = append(blogs, blog)
	}
	return blogs
}

func GetBlogById(blogId string) model.Blog {
	blog := model.Blog{}
	db.Where("blog_id = ?", blogId).First(&blog)
	return blog
}

func CreateBlog(blog *model.Blog) {
	db.Create(blog)
}

func UpdateBlog(blog *model.Blog) error {
	if err := db.Save(blog).Error; err != nil {
		log.Printf("error occurred while updating blog: %s\n", err)
		return err
	}
	return nil
}

func DeleteBlog(blogId string) error {
	if err := db.Where("blog_id = ?", blogId).Delete(model.Blog{}).Error; err != nil {
		log.Printf("error occurred while deleting blog err is: %s\n", err)
		return err
	}
	// TODO: delete association between this blog and the tags
	db.Exec("DELETE FROM blog_tag WHERE blog_id = ?", blogId)
	return nil
}

func SaveAssociationBetweenBlogAndTags(blogId string, tags []model.Tag) {
	oldTags := ListTagsByBlogId(blogId)
	newTagIds := make(map[string]bool)
	for _, tag := range tags {
		newTagIds[tag.TagId] = true
	}

	for _, oldTag := range oldTags {
		oldTagId := oldTag.TagId
		if !newTagIds[oldTagId] { // the tag is removed
			db.Exec("DELETE FROM blog_tag where tag_id = ?", oldTagId)
		} else { // the tag is still attached to the blog, remove it from newTagIds
			delete(newTagIds, oldTagId)
		}
	}

	for newTagId := range newTagIds {
		id := uuid.New().String()
		db.Exec("INSERT INTO blog_tag(id, blog_id, tag_id) VALUES(?, ?, ?)", id, blogId, newTagId)
	}
}

func ListBlogsGroupedByYear() map[string][]model.Blog {
	blogsGroupedByYear := make(map[string][]model.Blog)
	// list all the years
	rows, err := db.Raw("SELECT DISTINCT YEAR(create_at) FROM blog").Rows()
	if err != nil {
		log.Printf("error occurred while listing years of blogs, err is %s\n", err)
		return nil
	}
	var years []string
	for rows.Next() {
		var year string
		rows.Scan(&year)
		years = append(years, year)
	}
	// select blogs by year
	for _, year := range years {
		var blogsOfThisYear []model.Blog
		db.Where("YEAR(create_at) = ?", year).Find(&blogsOfThisYear)
		blogsGroupedByYear[year] = blogsOfThisYear
	}
	return blogsGroupedByYear
}
