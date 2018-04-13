package handler

type XiamiSearchSongResponse struct {
	State     int       `form:"state" json:"state" binding:"required"`
	Message   string    `form:"message" json:"message" `
	RequestId string    `form:"request_id" json:"request_id" binding:"required"`
	Data      XiamiData `form:"data" json:"data" binding:"required"`
}

type XiamiData struct {
	Songs    []XiamiSong `form:"songs" json:"songs"`
	Total    int         `form:"total" json:"total" binding:"required"`
	Previous int         `form:"previous" json:"previous"`
	Next     int         `form:"next" json:"next"`
}

type XiamiSong struct {
	SongId      int    `form:"song_id" json:"song_id"`
	SongName    string `form:"song_name" json:"song_name"`
	AlbumLogo   string `form:"album_logo" json:"album_logo"`
	ArtistId    int    `form:"artist_id" json:"artist_id"`
	ArtistName  string `form:"artist_name" json:"artist_name"`
	ArtistLogo  string `form:"artist_logo" json:"artist_logo"`
	ListenFile  string `form:"listen_file" json:"listen_file"`
	NeedPayFlag int    `form:"need_pay_flag" json:"need_pay_flag"`
}
