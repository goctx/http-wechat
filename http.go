package http_wechat

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

type handleFunc func(req *Request) interface{}

type wechat struct {
	token          string
	appid          string
	encodingAESKey string
	debug          bool
}

func New(token, appid, encodingAESKey string, debug bool) *wechat {
	return &wechat{
		token:          token,
		appid:          appid,
		encodingAESKey: encodingAESKey,
		debug:          debug,
	}
}

func (p *wechat) makeSign(nonce, timestamp string) string {
	sl := []string{p.token, nonce, timestamp}
	sort.Strings(sl)
	s := sha1.New()
	io.Copy(s, strings.NewReader(strings.Join(sl, "")))
	return fmt.Sprintf("%x", s.Sum(nil))
}

func (p *wechat) Run(handleFunc handleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if p.debug {
			log.Printf("%s %s\n", r.Method, r.URL.String())
		}
		q := r.URL.Query()
		nonce := q.Get("nonce")
		timestamp := q.Get("timestamp")
		signature := q.Get("signature")
		if p.makeSign(nonce, timestamp) != signature {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		switch r.Method {
		case "GET":
			fmt.Fprint(w, q.Get("echostr"))
		case "POST":
			req := Request{}
			defer r.Body.Close()
			if err := xml.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			w.Header().Set("Content-Type", "text/xml;charset=UTF-8")
			response := handleFunc(&req)
			now := time.Now().Unix()
			switch v := response.(type) {
			case *TextResponse:
				v.MsgType = "text"
				v.FromUserName = req.ToUserName
				v.ToUserName = req.FromUserName
				v.CreateTime = now
				xml.NewEncoder(w).Encode(&v)
			case *ImageResponse:
				v.MsgType = "image"
				v.FromUserName = req.ToUserName
				v.ToUserName = req.FromUserName
				v.CreateTime = now
				xml.NewEncoder(w).Encode(&v)
			case *VoiceResponse:
				v.MsgType = "voice"
				v.FromUserName = req.ToUserName
				v.ToUserName = req.FromUserName
				v.CreateTime = now
				xml.NewEncoder(w).Encode(&v)
			case *VideoResponse:
				v.MsgType = "video"
				v.FromUserName = req.ToUserName
				v.ToUserName = req.FromUserName
				v.CreateTime = now
				xml.NewEncoder(w).Encode(&v)
			case *MusicResponse:
				v.MsgType = "music"
				v.FromUserName = req.ToUserName
				v.ToUserName = req.FromUserName
				v.CreateTime = now
				xml.NewEncoder(w).Encode(&v)
			case *NewsResponse:
				v.MsgType = "news"
				v.ArticleCount = len(v.Articles.Articles)
				v.FromUserName = req.ToUserName
				v.ToUserName = req.FromUserName
				v.CreateTime = now
				xml.NewEncoder(w).Encode(&v)
			case *CustomerServiceResponse:
				v.MsgType = "transfer_customer_service"
				v.FromUserName = req.ToUserName
				v.ToUserName = req.FromUserName
				v.CreateTime = now
				xml.NewEncoder(w).Encode(&v)
			case string:
				w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
				fmt.Fprint(w, v)
			default:
				w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
				fmt.Fprint(w, "success")
			}
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	}
}
