package GPTService

import (
	"time"
)

const ApiKey = "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"
const Proxy = "198.44.136.53:17873"

// 结构体
type RespData struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Text string `json:"text"`
}

type UserInfo struct {
	ParentID       string
	ConversationId interface{}
	Ttl            time.Time
}

type ChatReq struct {
	Messages    string `json:"prompt"`
	Model       string `json:"model"`
	MaxToken    int    `json:"max_tokens"`
	Temperature int    `json:"temperature"`
}
