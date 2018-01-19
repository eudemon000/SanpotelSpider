package main

import (
	"fmt"
	"SanpotelSpider/src/scheduler"
	/*"SanpotelSpider/src/queue"
	"SanpotelSpider/src/kvdata"
	"strconv"
	"gotest/src/dis"*/
	"net/http"
	_"net/http/pprof"
)

//爬虫入口
func main() {
	fmt.Println("这是程序入口")
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	/*for i := 0; i < 100; i++ {
		queue.Push(i)
		fmt.Println("向队列插入", i)
	}
	 l := queue.Pull(12)
	 fmt.Println("l的长度为", len(l))
	 for i := 0; i < len(l); i++ {
	 	fmt.Println(l[i])
	 }
	 kvdata.CheckFinishedUrl("aaa")
	 for j := 0; j < 100; j++ {
	 	kvdata.AddUrlToWaitUrl("张三" + "李四" + strconv.Itoa(j))
	 	//kvdata.AddUrlToWaitUrl("https://studygolang.com/articles/2228")
	 }
	 kvdata.GetUrlForWaitUrl(100)*/
	 s := new(scheduler.SpiderDispther)
	 s.Start_urls = make([]string, 1)
	 s.Start_urls[0] = "http://www.99.com.cn/"
	//s.Start_urls[0] = "http://www.qq.com/"
	 s.Allowed_domains = make([]string, 1)
	 s.Allowed_domains[0] = "99.com.cn"
	 scheduler.InitData(s)
	 ch := make(chan int)
	 ch <- 1
}


