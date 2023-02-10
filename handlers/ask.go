package handlers

import (
	gpt "ChatGPTWeb/GPTService"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"strings"
)

func ASK(c *gin.Context) {
	question := c.Query("q")
	if len(question) < 1 {
		c.JSON(400, gin.H{
			"error": "请求不能为空",
		})
		return
	}
	fmt.Println(question)
	resp := gpt.Request2GPT(question, "POST")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error(err.Error())
		}
	}(resp.Body)

	bodyBytes, err := io.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))

	var respData gpt.RespData
	err = json.Unmarshal(bodyBytes, &respData)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "GPT返回数据解析失败",
		})
		return
	}
	fmt.Println(respData)
	ans := respData.Choices[0].Text
	ans = strings.Replace(ans, "\n", "", -1)

	c.JSON(200, ans)
	return
}
