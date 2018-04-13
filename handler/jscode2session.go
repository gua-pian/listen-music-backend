package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

var (
	jscode2session_url = "https://api.weixin.qq.com/sns/jscode2session"
)

type (
	Jscode2SessionRequest struct {
		JsCode string `form:"jscode" json:"jscode" binding:"required"`
	}
)

func Jscode2Session(c *gin.Context) {
	// Get all the parameters passed in.
	var request Jscode2SessionRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Init a http request to qq-music.
	u, _ := url.Parse(jscode2session_url)
	q := u.Query()
	q.Set("appid", AppId)
	q.Set("secret", AppSecret)
	q.Set("grant_type", "authorization_code")
	q.Set("js_code", request.JsCode)

	u.RawQuery = q.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	c.JSON(http.StatusOK, fmt.Sprintf("%s", result))
}
