package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	authorization = "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"
)

func main() {
	var OpenID string
	var msg string
	msg = "我有点饿了，你有什么推荐给我的吗"
	userInfoMap := make(map[string]*userInfo)
	info, ok := userInfoMap[OpenID]
	if !ok || info.ttl.Before(time.Now()) {
		//log.Infof("用户 %s 启动新的对话", OpenID)
		info = &userInfo{
			parentID:       uuid.New().String(),
			conversationId: nil,
		}
		userInfoMap[OpenID] = info
	} else {
		//log.Infof("用户 %s 继续对话", OpenID)
	}
	info.ttl = time.Now().Add(5 * time.Minute)
	// 发送请求
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", CreateChatReqBody(msg))
	if err != nil {
		log.Errorln(err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+authorization)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Content-Type", "application/json")
	proxy := "198.44.136.53:17873"
	proxyAddress, _ := url.Parse(proxy)
	//fmt.Println(proxyAddress)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyAddress),
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Errorln(err)
		return
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	s := string(bodyBytes)
	var respData RespData
	err = json.Unmarshal(bodyBytes, &respData)
	if err != nil {
		return
	}
	fmt.Println(s)
	//fmt.Println("***")
	//fmt.Println(respData.Choices[0])
	//fmt.Println("***")
	fmt.Println(respData.Choices[0].Text)

}

type RespData struct {
	Choices []Choice `json:"choices"`
}
type Choice struct {
	Text string `json:"text"`
}
type userInfo struct {
	parentID       string
	conversationId interface{}
	ttl            time.Time
}

func CreateChatReqBody(msg string) *bytes.Buffer {
	req := &ChatReq{
		Messages:    msg,
		Model:       "text-davinci-003",
		MaxToken:    256,
		Temperature: 0,
	}
	return bytes.NewBuffer(req.ToJson())
}
func (msg *ChatReq) ToJson() []byte {
	body, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return body
}

type ChatReq struct {
	Messages    string `json:"prompt"`
	Model       string `json:"model"`
	MaxToken    int    `json:"max_tokens"`
	Temperature int    `json:"temperature"`
}
