package aihuishou

import (
	"container/list"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

var Db *gorm.DB

type base struct {
	Id int
}
type Phone struct {
	Base  base
	Name  string
	Price float64
}
type HtmlUrl struct {
	//域名
	Url  string
	Type string
	Page int
}

func getHtml(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {

	}
	resp, err := client.Do(req)
	if err != nil {

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}
	return string(body)
}
func ToGo(url *HtmlUrl) *list.List {
	length := url.Page
	if length <= 0 {
		return nil
	}
	phoneList := list.New()
	for i := 1; i <= length; i++ {
		httpUrl := url.Url + url.Type
		if i != 1 {
			httpUrl += "-p" + strconv.Itoa(i)
		}
		httpUrl += "?all=true"
		html := getHtml(httpUrl)
		//回收类型
		phoneType := `title="(.*?) 回收"`
		phoneType_regexp := regexp.MustCompile(phoneType)
		phoneType_txt := phoneType_regexp.FindAllStringSubmatch(html, -1)
		if phoneType_txt != nil {

		}
		//价格
		price := `<div class="price">回收最高价 <em>￥(.*?)</em></div>`
		price_regexp := regexp.MustCompile(price)
		price_txt := price_regexp.FindAllStringSubmatch(html, -1)
		for j := 0; j < len(phoneType_txt); j++ {
			phone := new(Phone)
			phone.Name = phoneType_txt[j][1]
			v, _ := strconv.ParseFloat(price_txt[j][1], 64)
			phone.Price = v
			phoneList.PushBack(phone)
		}
	}
	return phoneList
}
func SetDB(db *gorm.DB) {
	Db = db
}
func getDb() *gorm.DB {
	return Db
}
