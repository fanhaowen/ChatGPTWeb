package main

import (
	gpt "ChatGPTWeb/GPTService"
	"ChatGPTWeb/handlers"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func main() {
	var OpenID string
	var port string
	port = "8080"
	//UserInfoMap是用来记录多个用户状态的，要存储上下文，但是暂时不用,不走上下文。
	userInfoMap := make(map[string]*gpt.UserInfo)
	info, ok := userInfoMap[OpenID]
	if !ok || info.Ttl.Before(time.Now()) {
		info = &gpt.UserInfo{
			ParentID:       uuid.New().String(),
			ConversationId: nil,
		}
		info.Ttl = time.Now().Add(5 * time.Minute)
		userInfoMap[OpenID] = info
	}

	router := gin.Default()
	router.GET("/api/ask", handlers.ASK)
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	err := router.Run(":" + port)
	if err != nil {
		return
	}

}
