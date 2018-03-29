package http_wechat

import "encoding/xml"

const (
	MsgTypeText       = "text"
	MsgTypeImage      = "image"
	MsgTypeVoice      = "voice"
	MsgTypeShortvideo = "shortvideo"
	MsgTypeLink       = "link"
	MsgTypeEvent      = "event"
)

const (
	EventTypeSubscribe   = "subscribe"
	EventTypeUnSubscribe = "unsubscribe"
	EventTypeLocation    = "LOCATION"
	EventTypeScan        = "SCAN"
	EventTypeClick       = "CLICK"
	EventTypeView        = "VIEW"
)

type Request struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	MsgId        int64
	// pic
	PicUrl  string
	MediaId string
	// voice
	Format      string
	Recognition string
	// video
	ThumbMediaId string
	// location
	LocationX string `xml:"Location_X"`
	LocationY string `xml:"Location_Y"`
	Scale     float64
	Label     string
	// link
	Title       string
	Description string
	Url         string
	// event:subscribe/unsubscribe
	Event string
	// event:scan
	EventKey string
	Ticket   string
	// event:LOCATION
	Latitude  float64
	Longitude float64
	Precision float64
}
