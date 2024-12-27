package consts

type BizCode struct {
	// 错误码
	Code int32
	// 错误描述
	Msg string
}

var (
	ResSuccess   = BizCode{0, "success"}
	SystemErr    = BizCode{1001, "服务繁忙，请稍后重试"}
	ParamError   = BizCode{1002, "参数错误"}
	WriteDbError = BizCode{1003, "数据写入异常"}
	ReadDbError = BizCode{1004, "数据写入异常"}
	ShortURLExpired = BizCode{1005, "短链接过期"}
)
