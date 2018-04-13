package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SearchSong(c *gin.Context) {
	MusicInfo := SearchSongResponse{}

	// Get all the parameters passed in.
	var request SearchSongRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"line": "27", "error": err.Error()})
		return
	}

	ch_err := make(chan string, 2)
	ch_music := make(chan MetaMusicInfo, 2)

	// Launch QQ Music Search.
	go QQSearchSong(&request, ch_music, ch_err)

	// Launch Xiami Music Search.
	go XiamiSearchSong(&request, ch_music, ch_err)

	for i := 1; i <= 2; i++ {
		select {
		case err_info := <-ch_err:
			log.Println(err_info + "Music Search Error!")
		case info := <-ch_music:
			if info.Meta == "qq" {
				MusicInfo.QQSongDatas = info.Data
			} else {
				MusicInfo.XiamiSongDatas = info.Data
			}
		case <-time.After(time.Second * 10):
			log.Println("Music Search Timeout!")
		}
	}

	c.JSON(http.StatusOK, MusicInfo)
}
