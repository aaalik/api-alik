package asql

import (
	"strconv"
	"time"
)

const (
	FormatDateTimeUs = "2006-01-02 15:04:05"

	KeyStartDate = "start_date"
	KeyEndDate   = "end_date"

	KeyOffset = "offset"
	KeyLimit  = "limit"

	DefaultOffset = 1
	DefaultLimit  = 20
)

func BetweenDate(params map[string]string) (startDate, endDate string) {
	defaultStartDate := time.Now().Format(FormatDateTimeUs)
	defaultEndDate := time.Now().Format(FormatDateTimeUs)

	startDate = params[KeyStartDate]
	if startDate == "" {
		startDate = defaultStartDate
	}

	endDate = params[KeyEndDate]
	if endDate == "" {
		endDate = defaultEndDate
	}

	return
}

func Pagination(params map[string]string) (limit, offset int) {
	if result, err := strconv.Atoi(params[KeyOffset]); err == nil {
		offset = result
	} else {
		offset = DefaultOffset
	}

	if result, err := strconv.Atoi(params[KeyLimit]); err == nil {
		limit = result
	} else {
		limit = DefaultLimit
	}

	return
}
