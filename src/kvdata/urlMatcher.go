package kvdata

import (
	"gitee.com/johng/gkvdb/gkvdb"
	"SanpotelSpider/src/utils"
	"fmt"
)

//待爬取URL
var waitUrl *gkvdb.DB

//已完成URLe
var finishedUrl *gkvdb.DB

func init() {
	waitUrl, _ = gkvdb.New("db/gkvdb", "waitUrl")
	finishedUrl, _ = gkvdb.New("db/gkvdb", "finishedUrl")
}

//检查URL是否已经爬取过
func CheckFinishedUrl(url string) bool {
	key := utils.Md5(url)
	s := finishedUrl.Get([]byte(key))
	if s == nil {
		fmt.Println("s空")
		return false
	} else {
		fmt.Println("s不空")
		return true
	}
}

//向已爬取表插入URL
func AddUrlToFinishedUrl(url string) bool {
	key := utils.Md5(url)
	value := "true"
	err := finishedUrl.Set([]byte(key), []byte(value))
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}

//向待爬取表插入数据
func AddUrlToWaitUrl(url string) bool {
	//mk := utils.Md5(url)
	mk := utils.Md5(url)
	key := []byte(mk)
	s := waitUrl.Get(key)
	fmt.Println("s===>", string(s))
	if s == nil {
		err := waitUrl.Set(key, []byte(url))
		if err != nil {
			fmt.Println(err)
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}

//从待爬取表获取URL
func GetUrlForWaitUrl(max int) map[string][]byte {
	result := waitUrl.Items(max)
	return result
}

