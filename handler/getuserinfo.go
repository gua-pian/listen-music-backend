package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	GetUserInfoRequest struct {
		EncryptedData string   `form:"encryptedData" json:"encryptedData" binding:"required"`
		Iv            string   `form:"iv" json:"iv" binding:"required"`
		RawData       string   `form:"rawData" json:"rawData" binding:"required"`
		Signature     string   `form:"signature" json:"signature" binding:"required"`
		UserInfo      UserInfo `form:"userInfo" json:"userInfo" binding:"required"`
		SessionData   string   `form:"sessionData" json:"sessionData" binding:"required"`
	}
	SessionData struct {
		SessionKey string `form:"session_key" json:"session_key" binding:"required"`
		OpenId     string `form:"openid" json:"openid" binding:"required"`
	}
	UserInfo struct {
		AvatarUrl string `form:"avatarUrl" json:"avatarUrl" binding:"required"`
		NickName  string `form:"nickName" json:"nickName" binding:"required"`
		Gender    int    `form:"gender" json:"gender" binding:"required"`
		City      string `form:"city" json:"city"`
		Province  string `form:"province" json:"province"`
	}
)

func GetUserInfo(c *gin.Context) {
	var request GetUserInfoRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}

	var sessionData SessionData
	err := json.Unmarshal([]byte(request.SessionData), &sessionData)
	if err != nil {
		fmt.Println(err)
	}

	key := sessionData.OpenId

	// Set openid:userinfo -> UserInfo
	fmt.Println(request.UserInfo)
	field := "userinfo"
	marshalUserInfo, _ := json.Marshal(request.UserInfo)
	_, err = redisClient.HSet(key, field, marshalUserInfo).Result()
	if err != nil {
		fmt.Println(err)
	}

	// Set openid:session_key -> session_key
	field = "session_key"
	_, err = redisClient.HSet(key, field, sessionData.SessionKey).Result()
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, "ok")
}
