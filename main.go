package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/zs5460/my"
)

const (
	// GetTokenURL ...
	GetTokenURL = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?"
	// SendMsgURL ...
	SendMsgURL = "https://qyapi.weixin.qq.com/cgi-bin/message/send?"
)

type wechat struct {
	AppID       string
	AppKey      string
	AgentID     string
	AccessToken string
}

type result struct {
	Code    int    `json:"errcode"`
	Message string `json:"errmsg"`
	Token   string `json:"access_token"`
}

type message struct {
	ToUser  string `json:"touser"`
	MsgType string `json:"msgtype"`
	AgentID string `json:"agentid"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

func (wc *wechat) GetToken() {
	var ret result
	url := GetTokenURL + "corpid=" + wc.AppID + "&corpsecret=" + wc.AppKey
	err := my.GetJSON(url, &ret)
	if err != nil {
		log.Fatal(err)
	}
	if ret.Code != 0 {
		log.Fatal(ret.Message)
	}
	wc.AccessToken = ret.Token
}

func (wc *wechat) Send(msg string) {
	var ret result
	url := SendMsgURL + "&access_token=" + wc.AccessToken
	m := message{
		ToUser:  "@all",
		MsgType: "text",
		AgentID: wc.AgentID,
		Text: struct {
			Content string `json:"content"`
		}{
			Content: msg,
		},
	}

	data, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	reply, err := my.PostURL(url, string(data))
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(reply, &ret)
	if err != nil {
		log.Fatal(err)
	}
	if ret.Code != 0 {
		log.Fatal(ret.Message)
	}
	fmt.Println("The message was sent successfully!")
}

func main() {
	msg := strings.Join(os.Args[1:], " ")
	if msg == "" {
		log.Fatalln("No message to send.")
	}

	var wc wechat
	my.MustLoadConfig("config.json", &wc)
	wc.GetToken()
	wc.Send(msg)
}
