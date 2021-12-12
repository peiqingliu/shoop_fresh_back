package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"time"
)

/*
 * 分页方式
 */
func Page(Limit,Page int) (limit,offset int)  {
	if Limit > 0 {
		limit = Limit
	}else {
		limit = 10
	}
	if Page > 0 {
		offset = (Page - 1) * limit
	}else {
		offset = -1
	}
	return limit,offset
}

/*
 * 排序方式
 */
func Sort(Sort string) (sort string)  {
	if Sort != "" {
		sort = Sort
	}else {
		sort = "create_at desc"
	}
	return sort
}

const TimeLayout  = "2006-01-02 15:04:05"

var local  = time.FixedZone("CST",8*3600)

/*
 *	获取当前时间
 */
func GetNow() string  {
		now := time.Now().In(local).Format(TimeLayout)
		return now
}

func TimeFormat(s string) string  {
	result,error := time.ParseInLocation(TimeLayout,s,time.Local)
	if error != nil {
		logrus.Error("日期转换错误")
	}
	return result.In(local).Format(TimeLayout)
}

func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}