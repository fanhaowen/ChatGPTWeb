package GPTService

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

// 服务
func (msg *ChatReq) ToJson() []byte {
	body, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return body
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

func BuildReq(msg string, method string) *http.Request {
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", CreateChatReqBody(msg))
	if err != nil {
		log.Errorln(err)
		return nil
	}
	req.Header.Set("Authorization", "Bearer "+ApiKey)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Content-Type", "application/json")
	return req
}
func BuildClient(proxy string) *http.Client {
	if len(proxy) == 0 {
		return http.DefaultClient
	}
	proxyAddress, _ := url.Parse(Proxy)
	//fmt.Println(proxyAddress)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyAddress),
		},
	}
	return client
}

func Request2GPT(msg string, method string) *http.Response {
	req := BuildReq(msg, method)

	client := BuildClient(Proxy)

	resp, err := client.Do(req)
	if err != nil {
		log.Errorln(err)
		return nil
	}
	return resp
}
