package models

import (
	"gorm.io/gorm"
)

/*
CREATE TABLE tags {
    id int auto_increment not null primary key,
    tag_name VARCHAR (256)
}

CREATE TABLE quetion_tags {
    id int auto_increment not null primary key,
    target_id int,
    tag_id int
}

*/
type Tag struct {
	Id       int    `json:"id"`
	Tag_name string `json:"tag_name"`
}

type QuetionTags struct {
	Id       int `json:"id"`
	TargetId int `json:"TargetId"` // 対象のquetion
	Tag_Id   int `json:"Tag_Id"`   // 対象のtag
}

func GetAllTags(db *gorm.DB) []Tag {
	var tags []Tag
	db.Find(&tags)
	return tags
}
