# ding_bot

Send message to dingding`s group chat by golang

### usage
- get go get github.com/jumoshen/ding.bot@v0.0.2
- or use `go get -u` to update this package

#### example

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
