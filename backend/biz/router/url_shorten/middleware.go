// Code generated by hertz generator.

package url_shorten

import (
	"context"
	"time"
	"url_shorten/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		ReqRespLogMiddleware(),
	}
}

type RequestInfo struct {
	Route  string `json:"route"`
	Params string `json:"params"`
}

func ReqRespLogMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 记录请求时间
		beginAt := time.Now()
		method := string(c.Method())
		req := &RequestInfo{
			Route: c.FullPath(),
		}
		if method == "GET" {
			params := c.Request.URI().QueryString()
			req.Params = string(params)
		}
		if method == "POST" {
			if req.Route != "/api/tools/file_post" {
				body := c.Request.Body()
				req.Params = string(body)
			} else {
				req.Params = "this is a file"
			}
		}

		hlog.CtxInfof(ctx, "Request route:%s, Method:%s, RequestBody:%s", req.Route, method, req.Params)
		c.Next(ctx)
		hlog.CtxInfof(ctx, "Response route:%v, code:%v, cost:%v", req.Route, c.Response.StatusCode(), utils.TimeSub(beginAt))
	}
}

func _queryMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _apiMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createshorturlMw() []app.HandlerFunc {
	// your code...
	return nil
}