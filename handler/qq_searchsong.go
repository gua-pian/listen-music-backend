package handler

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

var (
	qq_base_url = "http://c.y.qq.com/soso/fcgi-bin/search_cp?"
	aggr        = "1"
	lossless    = "1"
	cr          = "1"
	key_url     = "https://c.y.qq.com/base/fcgi-bin/fcg_musicexpress.fcg"
	song_url    = "http://dl.stream.qqmusic.qq.com/M500"
	poster_url  = "https://y.gtimg.cn/music/photo_new/T002R300x300M000"
)

func QQSearchSong(request *SearchSongRequest, ch_music chan MetaMusicInfo, ch_err chan string) {

	var metaMusicInfo = MetaMusicInfo{}

	// Init a http request to qq-music.
	u, _ := url.Parse(qq_base_url)
	q := u.Query()
	q.Set("p", strconv.Itoa(request.Page))
	q.Set("n", strconv.Itoa(request.Limit))
	q.Set("w", request.Key)
	q.Set("aggr", aggr)
	q.Set("lossless", lossless)
	q.Set("cr", cr)

	u.RawQuery = q.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		ch_err <- "QQ " + err.Error()
		return
	}

	result, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	// Unmarshal result into response.
	var response QQSearchSongResponse
	newresult := result[9 : len(result)-1]
	err = json.Unmarshal(newresult, &response)
	if err != nil {
		ch_err <- "QQ " + err.Error()
		return
	}

	MusicInfo := make([]SongInfo, 0)
	for _, song := range response.Data.Song.List {
		if v, ok := generateUrl(song); ok == nil {
			MusicInfo = append(MusicInfo, v)
		}
	}

	metaMusicInfo.Data = MusicInfo
	metaMusicInfo.Meta = "qq"
	ch_music <- metaMusicInfo
	return

}

func generateUrl(list QQList) (SongInfo, error) {
	var songInfo SongInfo
	// Generate a random.
	source := rand.NewSource(10000)
	r := rand.New(source)
	guid := int(r.Float32() * 1000000000)

	// Send request to qq-music, get key.
	u, _ := url.Parse(key_url)
	q := u.Query()
	q.Set("json", "3")
	q.Set("guid", strconv.Itoa(guid))

	u.RawQuery = q.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		return songInfo, err
	}
	result, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	newresult := result[13 : len(result)-2]
	var response QQGetKeyResponse
	err = json.Unmarshal(newresult, &response)
	if err != nil {
		return songInfo, err
	}
	key := response.Key

	generate_url := song_url + list.StrMediaMid + ".mp3?" + "vkey=" + key + "&guid=" + strconv.Itoa(guid) + "&fromtag=30"

	author := []string{}
	for _, v := range list.Singer {
		author = append(author, v.Name)
	}
	return SongInfo{Name: list.Songname, Author: author, Src: generate_url, Poster: poster_url + list.AlbumMid + ".jpg"}, nil

}
