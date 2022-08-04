package xdate

type sTime struct {
	timestamp int64
	date      string
	format    string
}

func New(timestamp int64) *sTime {
	t := &sTime{
		timestamp: timestamp,
	}
	return t
}

func NewDate(date string) *sTime {
	t := &sTime{
		date: date,
	}
	return t
}

func (this *sTime) GetDayTime(day int) (int64, int64) {

	return 0, 0
}

func (this *sTime) parseTime() *sTime {
	if this.timestamp == 0 {
		this.timestamp = ToTime(this.date)
	}
	return this
}

func (this *sTime) parseDate() *sTime {
	if this.date == "" {
		this.date = ToDate(this.timestamp, this.format)
	}
	return this
}

// 读取时间戳
func (this *sTime) GetTime() int64 {
	this.parseTime()
	return this.timestamp
}

// 读取日期
func (this *sTime) GetDate(_format ...string) string {
	if len(_format) > 0 {
		this.format = _format[0]
	}
	this.parseDate()
	return this.date
}
