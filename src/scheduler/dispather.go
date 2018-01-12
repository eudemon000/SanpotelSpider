package scheduler

import (
	"sync"
	"runtime"
	"SanpotelSpider/src/queue"
	"SanpotelSpider/src/kvdata"
	"SanpotelSpider/src/downloader"
	"fmt"
	"time"
)

type SpiderDispther struct {
	Start_urls		[]string
	Allowed_domains	[]string
	lock			*sync.Mutex
}

var sDispther *SpiderDispther

func InitData(s *SpiderDispther) {
	cpuNum := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuNum)
	sDispther = s
	//1、将入口URL放入待爬取表
	for _, data := range sDispther.Start_urls {
		kvdata.AddUrlToWaitUrl(data)
	}
	time.Sleep(time.Second * 5)
	//2、开启一个协程从待爬取表中获取URL，并放入队列
	go getTaskForDb(50)
	time.Sleep(time.Second * 2)
	//3、从队列中取出指定数量的任务，并执行
	go getTaskForQueue(50)



}

//从待爬取表中获取URL，并放入队列
func getTaskForDb(num int) {
	for {
		urls := kvdata.GetUrlForWaitUrl(num)
		for _, url := range urls {
			s := string(url)
			queue.Push(s)
		}
	}
}

//从任务队列取出指定数量的任务并执行
func getTaskForQueue(num int) {
	wait := queue.Pull(num)
	//var chUrls chan [num]interface{}
	chUrls := make([] chan interface{}, num)
	for index, url := range wait {
		chUrls[index] = make(chan interface{})
		go Downloader.Parser(url.(string), chUrls[index])
	}
	time.Sleep(time.Second * 20)
	for i, c := range chUrls {
		a := <- c
		fmt.Println("这是下一次爬取的URL：", i, "	", a)
	}
}

