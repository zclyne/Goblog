package service

import (
	"goblog/model"
	"log"
)

func ListTypes() []model.Type {
	var types []model.Type
	db.Find(&types)
	return types
}

func GetTypeById(typeId string) model.Type {
	var blogType model.Type
	db.Where("type_id = ?", typeId).First(&blogType)
	return blogType
}

func CreateType(t *model.Type) {
	db.Create(t)
}

func DeleteType(typeId string) error {
	if err := db.Where("type_id = ?", typeId).Delete(model.Type{}).Error; err != nil {
		log.Printf("error occurred while deleting type: %s\n", err)
		return err
	}
	return nil
}

func UpdateType(t *model.Type) error {
	if err := db.Save(t).Error; err != nil {
		log.Printf("error occurred while updating type: %s\n", err)
		return err
	}
	return nil
}
