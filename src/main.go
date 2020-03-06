package main

import (
	"aihuishou"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type BaseModel struct {
	Id    int `gorm:"primary_key;AUTO_INCREMENT"`
	Name  string
	Price float64
	Type  string
}

func main() {
	url := new(aihuishou.HtmlUrl)
	url.Url = "https://www.aihuishou.com/"
	url.Type = "shouji"
	url.Page = 135
	list := aihuishou.ToGo(url)
	for i := list.Front(); i != nil; i = i.Next() {
		phone := i.Value.(*aihuishou.Phone)
		fmt.Println(phone)
	}
}
