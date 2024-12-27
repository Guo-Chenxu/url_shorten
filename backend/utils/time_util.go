package utils

import (
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func TimeSub(t time.Time) string {
	return fmt.Sprintf("%.4f s", time.Since(t).Seconds())
}

func ParseTime(t, format string) time.Time {
	utc8, _ := time.LoadLocation("Asia/Shanghai")
	res, err := time.ParseInLocation(format, t, utc8)
	if err != nil {
		hlog.Errorf("时间转换失败, time = %s", t)
	}

	return res
}
