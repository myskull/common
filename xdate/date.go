package xdate

import (
	"regexp"
	"time"
)

/**
日期包
*/

func ToTime(date string) int64 {
	if len(date) <= 10 {
		date += " 00:00:00"
	}
	t, _ := time.Parse(parseSplit(date), date)
	return t.Unix()
}

/**
 *  @param timestamp 时间戳
 *  @param _format 日期格式:  Y-m-d H:i:s
 */
func ToDate(timestamp int64, _format ...string) string {
	t := time.Unix(timestamp, 0)
	format := "2006-01-02 15:04:05"
	if len(_format) > 0 && _format[0] != "" {
		format = _format[0]
		format = parseYMDHIS(format)
	}
	return t.Format(format)
}

func parseSplit(date string) string {
	reg, _ := regexp.Compile(`\/`)
	if reg.MatchString(date) {
		return "2006/01/02 15:04:05"
	}
	return "2006-01-02 15:04:05"
}

func parseYMDHIS(format string) string {
	checkReg, _ := regexp.Compile(`[YyMmDdHhIiSs]`)
	if !checkReg.MatchString(format) {
		return format
	}
	year, _ := regexp.Compile(`[Yy]+`)
	format = year.ReplaceAllString(format, "2006")
	month, _ := regexp.Compile(`[Mm]+`)
	format = month.ReplaceAllString(format, "01")
	day, _ := regexp.Compile(`[Dd]+`)
	format = day.ReplaceAllString(format, "02")
	hour, _ := regexp.Compile(`[Hh]+`)
	format = hour.ReplaceAllString(format, "15")
	i, _ := regexp.Compile(`[Ii]+`)
	format = i.ReplaceAllString(format, "04")
	s, _ := regexp.Compile(`[Ss]+`)
	format = s.ReplaceAllString(format, "05")
	return format
}
