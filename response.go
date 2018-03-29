package http_wechat

import "encoding/xml"

type response struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
}

type TextResponse struct {
	response
	Content string
}

type ImageResponse struct {
	response
	Image Media
}

type Media struct {
	MediaId string
}

type VoiceResponse struct {
	response
	Voice Media
}

type VideoResponse struct {
	response
	Video Video
}

type Video struct {
	MediaId     string
	Title       string `xml:"omitempty"`
	Description string `xml:"omitempty"`
}

type MusicResponse struct {
	response
	Music Music
}

type Music struct {
	Title        string `xml:"omitempty"`
	Description  string `xml:"omitempty"`
	MusicURL     string `xml:"omitempty"`
	HQMusicUrl   string `xml:"omitempty"`
	ThumbMediaId string
}

type NewsResponse struct {
	response
	ArticleCount int
	Articles     Articles
}

type Articles struct {
	Articles []Article
}

type Article struct {
	XMLName     xml.Name `xml:"item"`
	Title       string
	Description string
	PicUrl      string
	Url         string
}

type CustomerServiceResponse struct {
	response
	TransInfo TransInfo
}

type TransInfo struct {
	KfAccount string
}
