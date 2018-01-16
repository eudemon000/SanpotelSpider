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

var num int = 50

var dbCh chan int = make(chan int)

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

	kvdata.GetDBLength()

	//1、将入口URL放入待爬取表
	for _, data := range sDispther.Start_urls {
		kvdata.AddUrlToWaitUrl(data)
	}
	//time.Sleep(time.Second * 5)
	//2、开启一个协程从待爬取表中获取URL，并放入队列
	//var dbCh chan int = make(chan int)
	go getTaskForDb(num)
	time.Sleep(time.Second * 5)
	//3、从队列中取出指定数量的任务，并执行
	go getTaskForQueue(num)
}

//从待爬取表中获取URL，并放入队列
func getTaskForDb(num int) {
	for {
		urls := kvdata.GetUrlForWaitUrl(num)
		for _, url := range urls {
			s := string(url)
			fmt.Println("取出的", s)
			queue.Push(s)
		}
		//每次取出数据后，需要等到所有任务爬取完毕后，再进行下次任务
		dbCh <- 1
	}
}

//从任务队列取出指定数量的任务并执行
func getTaskForQueue(num int) {
	for{
	//等待从DB中获取数据后，再从队列获取任务

	wait := queue.Pull(num)
	waitLength := len(wait)
	n := 0
	if num > waitLength {
		n = waitLength
	} else {
		n = num
	}
	//var chUrls chan [num]interface{}
	chUrls := make([] chan interface{}, n)
	for index, url := range wait {
		chUrls[index] = make(chan interface{})
		go Downloader.Parser(url.(string), chUrls[index])
	}

	for _, c := range chUrls {
		a := <- c
		//fmt.Println("这是下一次爬取的URL：", a)
		fUrl := a.(Downloader.NextUrl).FinishUrl
		kvdata.AddUrlToFinishedUrl(fUrl)
		kvdata.RemoveForWaitUrl(fUrl)
		for _, u := range a.(Downloader.NextUrl).ResultUrl {
			kvdata.AddUrlToWaitUrl(u)
			//fmt.Println("已经插入表的URL：", u)
		}
	}
	fmt.Println("循环执行完毕")
	d := <- dbCh
	fmt.Println("getTaskForQueue", d)
	}
}

