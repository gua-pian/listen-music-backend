package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	xiami_base_url = "http://api.xiami.com/web?"
	referer        = "http://h.xiami.com/"
	cookie         = "user_from=2;XMPLAYER_addSongsToggler=0;XMPLAYER_isOpen=0;_xiamitoken=cb8bfadfe130abdbf5e2282c30f0b39a;"
	user_agent     = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.75 Safari/537.36"
	content_type   = "application/x-www-form-urlencoded"
)

func XiamiSearchSong(request *SearchSongRequest, ch_music chan MetaMusicInfo, ch_err chan string) {

	var metaMusicInfo = MetaMusicInfo{}

	header := http.Header{}
	header.Set("Referer", referer)
	header.Set("Cookie", cookie)
	header.Set("User-Agent", user_agent)
	header.Set("Content-Type", content_type)

	u, _ := url.Parse(xiami_base_url)
	q := u.Query()
	q.Set("v", "2.0")
	q.Set("app_key", "1")
	q.Set("key", request.Key)
	q.Set("page", "1")
	q.Set("limit", "10")
	q.Set("r", "search/songs")
	u.RawQuery = q.Encode()

	b := strings.NewReader(q.Encode())

	req, err := http.NewRequest("POST", u.String(), b)
	req.Header = header

	if err != nil {
		ch_err <- "Xiami " + err.Error()
		return
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ch_err <- "Xiami " + err.Error()
		return
	}

	result, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var response XiamiSearchSongResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		ch_err <- "Xiami " + err.Error()
		return
	}

	MusicInfo := make([]SongInfo, 0)
	for _, v := range response.Data.Songs {
		MusicInfo = append(MusicInfo, SongInfo{Name: v.SongName, Author: []string{v.ArtistName}, Src: v.ListenFile, Poster: v.AlbumLogo})
	}
	metaMusicInfo.Data = MusicInfo
	metaMusicInfo.Meta = "xiami"
	ch_music <- metaMusicInfo
}
