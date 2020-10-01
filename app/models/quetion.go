package models

import (
	"fmt"

	"gorm.io/gorm"
)

type SimpleQuetion struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"context"`
}
type Quetion struct {
	Id             int    `json:"id"`
	Title          string `json:"title"`
	Content        string `json:"context"`
	Created_at     int    `json:"created_at"`
	Modifyed_at    int    `json:"modifyed_at"`
	Group_id       string `json:"group_id"`
	Privated       bool   `json:"privated"`
	Create_user_id int    `json:"craete_user_id"`
}

// フロント側/bot側に渡す型
type QuetionWithTags struct {
	Id             int    `json:"id"`
	Title          string `json:"title"`
	Content        string `json:"context"`
	Created_at     int    `json:"created_at"`
	Modifyed_at    int    `json:"modifyed_at"`
	Group_id       string `json:"group_id"`
	Privated       bool   `json:"privated"`
	Create_user_id int    `json:"craete_user_id"`
	Tags           []Tag
}

// フロント側/bot側から受け取る型
type GetQuetionWithStringTags struct {
	Title          string `json:"title"`
	Content        string `json:"context"`
	Created_at     int    `json:"created_at"`
	Modifyed_at    int    `json:"modifyed_at"`
	Group_id       string `json:"group_id"`
	Privated       bool   `json:"privated"`
	Create_user_id int    `json:"craete_user_id"`
	Tags           []string
}

func (quetion *GetQuetionWithStringTags) CreateQuetionWithTags(db *gorm.DB) {
	allTags := GetAllTags(db)
	mapTags := map[string]int{}
	for _, value := range allTags {
		mapTags[value.Tag_name] = value.Id
	}
	setQuetion := Quetion{
		Title:          quetion.Title,
		Content:        quetion.Content,
		Created_at:     quetion.Created_at,
		Modifyed_at:    quetion.Modifyed_at,
		Group_id:       quetion.Group_id,
		Privated:       quetion.Privated,
		Create_user_id: quetion.Create_user_id}
	db.Create(&setQuetion)
	// var maxQuetion Quetion
	// maxCount := maxQuetion.Id + 1
	// db.Select("MAX(id)").Find(&maxQuetion)
	for _, value := range quetion.Tags {
		if _, ok := mapTags[value]; !ok {
			// ユニークなtagが追加
			setTag := Tag{Tag_name: value}
			db.Create(setTag)
			mapTags[value] = setTag.Id
		}
		result := db.Create(&QuetionTags{TargetId: setQuetion.Id, Tag_Id: mapTags[value]})
		if err := result.Error; err != nil {
			fmt.Println("err", err)
			// println(err)
		}
	}

}

func GetAllQuetion(db *gorm.DB) []Quetion {
	var quetions []Quetion
	db.Find(&quetions)
	fmt.Println("確認！！", quetions)
	return quetions
}

func GetQuetion(db *gorm.DB) Quetion {
	var quetion Quetion
	db.Find(&quetion)
	fmt.Println("確認!!T", quetion)
	return quetion
}

func (quetion *Quetion) CreateQuetion(db *gorm.DB) {
	result := db.Create(&quetion)
	if err := result.Error; err != nil {
		fmt.Println("err", err)
		// println(err)
	}
}
