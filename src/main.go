package main

import (
	"aihuishou"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"strconv"
)

type BaseModel struct {
	Id    int `gorm:"primary_key;AUTO_INCREMENT"`
	Name  string
	Price float64
	Type  string
}

var file *os.File
var htmlUrl []*aihuishou.HtmlUrl

func main() {
	//创建excel文件
	f, err := os.Create("D:/haha3.txt")
	if err != nil {
		panic(err)
	}
	file = f
	defer f.Close()
	htmlUrl = make([]*aihuishou.HtmlUrl, 3)
	htmlUrl[0] = &aihuishou.HtmlUrl{Url: "https://www.aihuishou.com/", Type: "shouji", Page: 135}
	htmlUrl[1] = &aihuishou.HtmlUrl{Url: "https://www.aihuishou.com/", Type: "pingban", Page: 35}
	htmlUrl[2] = &aihuishou.HtmlUrl{Url: "https://www.aihuishou.com/", Type: "laptop", Page: 313}
	write()

}

func write() {
	file.WriteString("类型 \t 价格 \t\r\n")
	for _, url := range htmlUrl {
		list := aihuishou.ToGo(url)
		file.WriteString(url.Type + "\t\r\n")
		for i := list.Front(); i != nil; i = i.Next() {
			phone := i.Value.(*aihuishou.Phone)
			//fmt.Println(phone)
			file.WriteString(phone.Name + "\t" + strconv.FormatFloat(phone.Price, 'g', 6, 64) + "\t" + "\r\n")
		}
	}

}
