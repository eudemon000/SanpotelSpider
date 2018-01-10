package scheduler

import (
	"sync"
	"runtime"
	"SanpotelSpider/src/queue"
	"SanpotelSpider/src/kvdata"
	"SanpotelSpider/src/downloader"
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
	//2、从待爬取表中获取URL，并放入队列
	urls := kvdata.GetUrlForWaitUrl(100)
	for _, url := range urls {
		s := string(url)
		queue.Push(s)
	}

	wait := queue.Pull(50)

	for _, w := range wait {
		go Downloader.Parser(w.(string))
	}

}

