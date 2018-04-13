package handler

type SearchSongRequest struct {
	Key   string `form:"key" json:"key" binding:"required"`
	Page  int    `form:"page" json:"page"`
	Limit int    `form:"limit" json:"limit"`
}

type QQSearchSongResponse struct {
	Code    int        `form:"code" json:"code" binding:"required"`
	Data    QQSongData `form:"data" json:"data" binding:"required"`
	Message string     `form:"message" json:"message" binding:"required"`
	Notice  string     `form:"notice" json:"notice" binding:"required"`
	Subcode int        `form:"subcode" json:"subcode" binding:"required"`
	Time    int        `form:"time" json:"time" binding:"required"`
	Tips    string     `form:"tips" json:"tips" binding:"required"`
}

type QQSongData struct {
	KeyWord   string `form:"keyword" json:"keyword" binding:"required"`
	Song      QQSong `form:"song" json:"song" binding:"required"`
	TotalTime int    `form:"totaltime" json:"totaltime" binding:"required"`
}

type QQSong struct {
	Curnum   int      `form:"curnum" json:"curnum" binding:"required"`
	Curpage  int      `form:"curpage" json:"curpage" binding:"required"`
	List     []QQList `form:"list" json:"list" binding:"required"`
	Totalnum int      `form:"totalnum" json:"totalnum" binding:"required"`
}

type QQList struct {
	Albumname   string     `form:"albumname" json:"albumname" binding:"required"`
	AlbumMid    string     `form:"albummid" json:"albummid" binding:"required"`
	Songname    string     `form:"songname" json:"songname" binding:"required"`
	SongId      int        `form:"songid" json:"songid" binding:"required"`
	SongMid     string     `form:"songmid" json:"songmid" binding:"required"`
	StrMediaMid string     `form:"strMediaMid" json:"strMediaMid" binding:"required"`
	Stream      int        `form:"stream" json:"stream" binding:"required"`
	Switch      int        `form:"switch" json:"switch" binding:"required"`
	Singer      []QQSinger `form:"singer" json:"singer" binding:"required"`
	Pay         QQPay      `form:"pay" json:"pay" binding:"required"`
}

type QQSinger struct {
	Id   int    `form:"id" json:"id" binding:"required"`
	Mid  string `form:"mid" json:"mid" binding:"required"`
	Name string `form:"name" json:"name" binding:"required"`
}

type QQPay struct {
	PayPlay int `form:"payplay" json:"payplay" binding:"required"`
}

type QQGetKeyResponse struct {
	Code int    `form:"code" json:"code"`
	Key  string `form:"key" json:"key"`
}

type SongInfo struct {
	Poster string   `json:"poster"`
	Name   string   `json:"name"`
	Author []string `json:"author"`
	Src    string   `json:"src"`
}

type MetaMusicInfo struct {
	Data []SongInfo
	Meta string
}

type SearchSongResponse struct {
	QQSongDatas    []SongInfo `form:"QQSongdatas" json:"QQSongdatas"`
	XiamiSongDatas []SongInfo `form:"XiamiSongdatas" json:"XiamiSongdatas"`
}
