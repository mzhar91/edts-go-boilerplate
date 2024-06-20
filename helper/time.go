package utils

import (
	"time"
	
	_config "sg-edts.com/edts-go-boilerplate/config"
)

// ToLocalTimezone convert time to env TIMEZONE time
func ToLocalTimezone(times ...*time.Time) {
	timezone := _config.Cfg.Timezone
	loc, _ := time.LoadLocation(timezone)
	
	for _, o := range times {
		if !o.IsZero() {
			*o = o.In(loc)
		}
	}
}

// ToServerTimezone convert time to server timezone
func ToServerTimezone(times ...*time.Time) {
	for _, o := range times {
		if !o.IsZero() {
			if o.Location().String() == "UTC" {
				AddLocalTimezone(o)
			}
			*o = o.UTC()
		}
	}
}

// AddLocalTimezone convert time with no timezone to env TIMEZONE time
func AddLocalTimezone(times ...*time.Time) {
	timezone := _config.Cfg.Timezone
	loc, _ := time.LoadLocation(timezone)
	
	const format = "2006-01-02T15:04:05"
	for _, o := range times {
		if !o.IsZero() {
			*o, _ = time.ParseInLocation(format, o.Format(format), loc)
		}
	}
}

// NowInLocal convert time with no timezone to env TIMEZONE time
func NowInLocal() time.Time {
	timezone := _config.Cfg.Timezone
	loc, _ := time.LoadLocation(timezone)
	return time.Now().In(loc)
}

func TimestampToLocalTimeZone(timestamp int) string {
	timezone := _config.Cfg.Timezone
	loc, _ := time.LoadLocation(timezone)
	tm := time.Unix(int64(timestamp), 0).In(loc).Format("2006-01-02")
	return tm
}
