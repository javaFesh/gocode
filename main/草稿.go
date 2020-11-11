package main

import (
	"fmt"
	_ "strings"
	"time"
)

type a struct {
	u int32
	b int32
}
type b struct {
	a
}

func main() {
	loc, _ := time.LoadLocation("Local")
	fmt.Println(time.Date(2099, 12, 12, 12, 12, 12, 0, loc).Unix())
	Str2Stamp := Str2Stamp("2099-12-12 12:12:12")
	fmt.Println(Str2Stamp)

}
func Str2Stamp(formatTimeStr string) int64 {
	timeStruct := Str2Time(formatTimeStr)
	millisecond := timeStruct.UnixNano() / 1e6
	return millisecond
}

/**字符串->时间对象*/
func Str2Time(formatTimeStr string) time.Time {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, formatTimeStr, loc) //使用模板在对应时区转化为time.time类型

	return theTime

}
