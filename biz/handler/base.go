package handler

import (
	"context"
	"url_shorten/consts"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hz_consts "github.com/cloudwego/hertz/pkg/protocol/consts"
)

type BaseResponse struct {
	Code int32         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type BaseHandler struct{}

func (b BaseHandler) ErrorResponse(ctx context.Context, c *app.RequestContext, err *consts.BizCode, data interface{}) {
	hlog.CtxErrorf(ctx, "ErrorResponse code:%d, errmsg:%s", err.Code, err.Msg)
	c.JSON(hz_consts.StatusOK, map[string]interface{}{
		"code": err.Code,
		"msg":  err.Msg,
		"data": data,
	})
}

func (b BaseHandler) SuccessResponse(c *app.RequestContext, resp interface{}) {
	c.JSON(hz_consts.StatusOK, resp)
}

func (b BaseHandler) Response(c *app.RequestContext, bizCode *consts.BizCode, data interface{}) {
	c.JSON(hz_consts.StatusOK, map[string]interface{}{
		"code": bizCode.Code,
		"msg":  bizCode.Msg,
		"data": data,
	})
}
