package jptime

import "time"

var jst = time.FixedZone("Asia/Tokyo", 9*60*60)

const (
	LayoutDateTime = "2006-01-02 15:04:05"
	LayoutDate     = "2006-01-02"
)

func Now() time.Time {
	return time.Now().In(jst)
}

func ParseDateTime(str string) (time.Time, error) {
	return time.ParseInLocation(LayoutDateTime, str, jst)
}

func ParseDate(str string) (time.Time, error) {
	return time.ParseInLocation(LayoutDate, str, jst)
}

func FormatDateTime(t time.Time) string {
	return t.In(jst).Format(LayoutDateTime)
}

func FormatDate(t time.Time) string {
	return t.In(jst).Format(LayoutDate)
}
