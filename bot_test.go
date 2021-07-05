package ding_bot

import (
	"log"
	"testing"
)

func BenchmarkNew(b *testing.B) {
	accessToken := "xxxx"
	secret := "xxxx"
	dt := New(accessToken, WithSecret(secret))

	textContent := "我就是我, 是不一样的烟火"
	//atMobiles := SendWithAtMobiles([]string{"15210123291"})
	for i := 0; i < b.N; i++ {
		if err := dt.RobotSendText(textContent); err != nil {
			log.Fatal(err)
		}
	}

}

func TestNew(t *testing.T) {
	accessToken := "xxxx"
	secret := "xxxx"
	dt := New(accessToken, WithSecret(secret))

	// text类型
	textContent := "我就是我, 是不一样的烟火"
	atMobiles := SendWithAtMobiles([]string{"15210123291"})
	if err := dt.RobotSendText(textContent, atMobiles); err != nil {
		log.Fatal(err)
	}
	log.Println(dt)

	// link类型
	linkTitle := "时代的火车向前开"
	linkText := `这个即将发布的新版本，创始人xx称它为“红树林”。` +
		`而在此之前，每当面临重大升级，产品经理们都会取一个应景的代号，` +
		`这一次，为什么是“红树林”？`
	linkMessageURL := "https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI"
	linkPicURL := "https://cdn.pixabay.com/photo/2020/05/05/08/05/butterfly-5131967_960_720.jpg"
	if err := dt.RobotSendLink(linkTitle, linkText, linkMessageURL, linkPicURL); err != nil {
		log.Fatal(err)
	}
	log.Println(dt)

	// markdown类型
	markdownTitle := "markdown"
	markdownText := "#### 杭州天气 @176XXXXXXXX\n" +
		"> 9度，西北风1级，空气良89，相对温度73%\n" +
		"> ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n" +
		"> ###### 10点20分发布 [天气](https://www.dingtalk.com)\n"
	if err := dt.RobotSendMarkdown(markdownTitle, markdownText); err != nil {
		log.Fatal(err)
	}
	log.Println(dt)
}
