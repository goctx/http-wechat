# http-wechat
微信公众平台被动回复中间件，基于标准库http开发。

## Get Started

```go
package main

import (
	hw "github.com/goctx/http-wechat"
)

func main() {
	wechat := hw.New("token", "appid", "encodingAESKey", true)
	wechat.Run(func(req *hw.Request) interface{} {
		if req.MsgType == hw.MsgTypeEvent && req.Event == hw.EventTypeSubscribe {
			return &hw.TextResponse{
				Content: "谢谢关注",
			}
		}
		return "success"
	})
}
```

## TODO

+ [ ] 加密模式支持