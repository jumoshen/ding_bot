# ding_bot

[![Build Status](https://travis-ci.com/jumoshen/ding_bot.svg?branch=master)](https://travis-ci.com/jumoshen/ding_bot)

Send message to dingding`s group chat by golang

### install
- go get github.com/jumoshen/ding.bot@v0.0.5
- or use `go get -u` to update this package

### example

- we can get the usage for ding_bot`s message sender at bot_test.go

```golang
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
```

#### Documents
- [pkg.go](https://pkg.go.dev/github.com/jumoshen/ding_bot)
- [钉钉开发文档](https://developers.dingtalk.com/document/app/custom-robot-access/title-r82-8g5-0sk)
