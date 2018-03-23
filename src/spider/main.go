package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
	"math/rand"
	"strconv"
	"time"
	"sync"
)

var wg sync.WaitGroup

// 主协程首先获取到所有的视频链接页面(非具体链接)
func getVideoUrls(videoPage chan string) {
	baseUri := "http://www.333avtb.com"
	for i := 1; i <= 515; i++ {
		tmpUrl := "http://www.333avtb.com/recent/"
		if i != 1 {
			tmpUrl += strconv.Itoa(i)
		}
		doc, err := goquery.NewDocument(tmpUrl)
		if err != nil {
			glog.Errorf("request %s url raise error:%s\n", strconv.Itoa(i), err.Error())
		} else {
			doc.Find(".video").Each(func(i int, s *goquery.Selection) {
				tmpUri := s.Find("a").AttrOr("href", "")
				videoPage <- fmt.Sprintf("%s%s", baseUri, tmpUri)
				glog.Infof("%s%s\n", baseUri, tmpUri)
			})
		}
		// 遍历每页需要sleep 0~2秒
		rand.Seed(time.Now().UnixNano())
		x := rand.Intn(2)
		time.Sleep(time.Duration(x) * time.Second)
	}
	close(videoPage)
}

// 分析视频链接页面,获取具体视频链接
// 慢些抓取分析
func analyzePage(videoPage chan string) {
	for {
		videoPageUrl, ok := <-videoPage
		if !ok {
			glog.Infoln("one goroutine of analyzePage end")
			break
		}
		if videoPageUrl != "" {
			doc, err := goquery.NewDocument(videoPageUrl)
			if err != nil {
				glog.Errorf("analyze url:%s raise error:%s\n", videoPageUrl, err.Error())
			} else {
				doc.Find("source").Each(func(i int, s *goquery.Selection) {
					videoUrl := s.AttrOr("src", "")
					glog.Infof("video URL is %s\n", videoUrl)
				})
			}
			rand.Seed(time.Now().UnixNano())
			x := rand.Intn(2)
			time.Sleep(time.Duration(x) * time.Second)
		}
	}
	wg.Done()
}

// 设计方案1
// 主协程触发一个协程不断获取到视频页面的链接，传入一个通道里面 done
// 主协程启动若干子协程获取到具体视频链接，传入一个下载通道 doing
// 主协程启动若干子协程下载具体视频，是否压缩视频内容看清再说
// todo 所有的视频页面链接的后缀部分存储起来
// todo 所有的视频链接的后缀存储起来
func main() {
	flag.Parse()
	defer glog.Flush()
	var videoPageChan = make(chan string, 30)

	go getVideoUrls(videoPageChan)

	for i := 0; i < 60; i++ {
		wg.Add(1)
		go analyzePage(videoPageChan)
	}

	wg.Wait()
	glog.Infoln("all goroutine execute over")
}
