package kvdata

import (
	"gitee.com/johng/gkvdb/gkvdb"
	"SanpotelSpider/src/utils"
	"fmt"
	"sync"
)

//待爬取URL
type waitDB struct {
	waitUrl *gkvdb.DB
	lock	*sync.Mutex
}

//已完成URL
type finishedDB struct {
	finishedUrl *gkvdb.DB
	lock		*sync.Mutex
}

var wDB *waitDB

var fDB *finishedDB

func init() {
	wDB = new(waitDB)
	wDB.waitUrl, _ = gkvdb.New("db/gkvdb", "waitUrl")
	wDB.lock = new(sync.Mutex)
	fDB = new(finishedDB)
	fDB.finishedUrl, _ = gkvdb.New("db/gkvdb", "finishedUrl")
	fDB.lock = new(sync.Mutex)
	//waitUrl, _ = gkvdb.New("db/gkvdb", "waitUrl")
	//finishedUrl, _ = gkvdb.New("db/gkvdb", "finishedUrl")
}

//检查URL是否已经爬取过
func CheckFinishedUrl(url string) bool {
	fDB.lock.Lock()
	defer fDB.lock.Unlock()
	key := utils.Md5(url)
	s := fDB.finishedUrl.Get([]byte(key))
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
	fDB.lock.Lock()
	defer fDB.lock.Unlock()
	key := utils.Md5(url)
	value := "true"
	err := fDB.finishedUrl.Set([]byte(key), []byte(value))
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
	//fmt.Println("要插入的URL：", url)
	wDB.lock.Lock()
	wDB.lock.Unlock()
	if len(url) == 0 || url == " " {
		return false
	}
	mk := utils.Md5(url)
	key := []byte(mk)
	s := wDB.waitUrl.Get(key)
	//fmt.Println("s===>", string(s))
	if s == nil {
		err := wDB.waitUrl.Set(key, []byte(url))
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
	wDB.lock.Lock()
	defer wDB.lock.Unlock()
	result := wDB.waitUrl.Items(max)
	return result
}

//从待爬取表中删除已经爬取完成的URL
func RemoveForWaitUrl(url string) {
	wDB.lock.Lock()
	defer wDB.lock.Unlock()
	key := utils.Md5(url)
	wDB.waitUrl.Remove([]byte(key))
}


func GetDBLength() {
	wDB.lock.Lock()
	defer wDB.lock.Unlock()
	fmt.Println(len(wDB.waitUrl.Items(-1)))
}

