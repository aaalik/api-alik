package autils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const ISO8601_LAYOUT = "2006-01-02T15:04:05.999+07:00"
const DATE_LAYOUT = "2006-01-02"
const DATE2_LAYOUT = "2 Jan 2006"
const DATETIME_LAYOUT = "2006-01-02 15:04:05"
const DATETIME2_LAYOUT = "2006-01-02T15:04:05Z" //FORMAT GOLANG GET DATETIME DARI MYSQL
const BCA_LAYOUT = "15/03/2014 22:07:40"
const TIME_STRING_LAYOUT = "2006-01-02 15:04:05.999999999 -0700 MST"

func DefaultTimeZone() *time.Location {
	tz := time.FixedZone("UTC+7", 7*60*60)
	if tz == nil {
		loc, err := time.LoadLocation("Asia/Jakarta")
		if err == nil {
			tz = loc
		}
	}

	return tz
}

func DateRange(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}

func GetMonth(m time.Month) string {
	if m < 10 {
		return "0" + strconv.Itoa(int(m))
	} else {
		return strconv.Itoa(int(m))
	}
}

func GetDateTime(timeZone *time.Location) time.Time {
	if timeZone == nil {
		timeZone = DefaultTimeZone()
	}

	return time.Now().In(timeZone)
}

func GetDateTimeWithMinutesInterval(minutes int) time.Time {
	return GetDateTime(nil).Add(time.Duration(minutes) * time.Minute)
}

func GetMinutesIntervalFromNow(t time.Time) float64 {
	t = t.Add(time.Hour * -7)
	duration := time.Since(t)

	return duration.Minutes()
}

func GetDayIntervalFromNow(startWorking time.Time) int {
	duration := time.Since(startWorking)
	interval := duration.Hours()

	return int((interval / 24) + 0.5)
}

func FixDateTime(datetime string) (string, error) {
	var err error

	formattedHour := ""
	if strings.Contains(datetime, "PM") {
		addColon := false
		hour24format := 00
		formattedHour = datetime[11:13]

		if strings.Contains(formattedHour, ":") {
			addColon = true
			formattedHour = string(formattedHour[0])

			if tmp, err := strconv.Atoi(formattedHour); err == nil {
				hour24format = tmp + 12
			} else {
				err = errors.New(fmt.Sprintf("error while parsing fixing datetime: %s", err))
			}
		} else {
			if tmp, err := strconv.Atoi(formattedHour); err == nil {
				hour24format = tmp + 12
			} else {
				err = errors.New(fmt.Sprintf("error while parsing fixing datetime: %s", err))
			}
		}

		if hour24format == 24 {
			formattedHour = "12"
		} else if hour24format > 9 {
			formattedHour = strconv.Itoa(hour24format)
		} else {
			formattedHour = "0" + strconv.Itoa(hour24format)
		}

		if addColon {
			formattedHour += ":"
		}
	} else if strings.Contains(datetime, "AM") {
		formattedHour = datetime[11:13]

		if formattedHour == "12" {
			formattedHour = "00"
		}
	}

	if formattedHour != "" {
		datetime = datetime[0:11] + formattedHour + datetime[13:19]
	}

	datetime = strings.TrimSpace(datetime)
	if datetime == "" {
		err = errors.New("empty datetime")
	}

	return datetime, err
}
