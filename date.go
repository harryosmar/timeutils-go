package timeutils_go

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func GetMonthRange(month int, year int) (gte int64, lte int64, err error) {
	if month < 1 || month > 12 {
		return lte, gte, fmt.Errorf("invalid month")
	}

	monthStr := strconv.Itoa(month)
	if month < 10 {
		monthStr = fmt.Sprintf("0%s", monthStr)
	}
	gteStr := fmt.Sprintf("%d-%s-01T00:00:00+07:00", year, monthStr)
	gteTime, err := time.Parse(time.RFC3339, gteStr)
	if err != nil {
		return lte, gte, err
	}
	gte = gteTime.Unix()

	nextMonth := month + 1
	if nextMonth > 12 {
		nextMonth = 1
		year += 1
	}
	nextMonthStr := strconv.Itoa(nextMonth)
	if nextMonth < 10 {
		nextMonthStr = fmt.Sprintf("0%s", nextMonthStr)
	}
	nextMonthTimeStr := fmt.Sprintf("%d-%s-01T00:00:00+07:00", year, nextMonthStr)
	nextMonthTime, err := time.Parse(time.RFC3339, nextMonthTimeStr)
	if err != nil {
		return lte, gte, err
	}

	lteTime := time.Unix(nextMonthTime.Unix()-1, 0)
	lte = lteTime.Unix()

	return gte, lte, nil
}

func DayInUnix(t time.Time) float64 {
	x := t.Unix()
	if x%86400 == 0 {
		return float64(x / 86400)
	}

	// expect t time.Now calculate n days in Asia/Jakarta UTC+7 hours
	return math.Ceil(float64(x)/86400) - 1
}

// DayInUnixJakartaTimezone only used this for time.Now from host
// DayInUnixJakartaTimezone got time.Now convert to mysql JakartaTimezone
func DayInUnixJakartaTimezone(t time.Time) float64 {
	// expect t time.Now calculate n days in Asia/Jakarta UTC+7 hours
	return DayInUnix(time.Unix(t.Unix()+7*3600, 0))
}

func HourInUnix(t time.Time) float64 {
	x := t.Unix()
	if x%3600 == 0 {
		return float64(x / 3600)
	}

	// expect t time.Now calculate n days in Asia/Jakarta UTC+7 hours
	return math.Ceil(float64(x)/3600) - 1
}

// HourInUnixJakartaTimezone only used this for time.Now from host
// HourInUnixJakartaTimezone got time.Now convert to mysql JakartaTimezone
func HourInUnixJakartaTimezone(t time.Time) float64 {
	// expect t time.Now calculate n days in Asia/Jakarta UTC+7 hours
	return HourInUnix(time.Unix(t.Unix()+7*3600, 0))
}

func GetExpirationTillEndOfTodayJakartaTimezone(t time.Time) int64 {
	end := (DayInUnixJakartaTimezone(t) + 1) * 86400
	return int64(end) - t.Unix() - 7*3600
}

func PlusHourToTime(t time.Time, n int64) time.Time {
	return time.Unix((int64(HourInUnix(t))*3600)+(n*3600), 0)
}

type FormatParam struct {
	T        time.Time
	Location string
	Format   string
}

func Format(p FormatParam) (string, error) {
	if p.Location == "" {
		p.Location = "Asia/Jakarta"
	}

	if p.Format == "" {
		p.Format = time.RFC3339
	}

	jakartaLoc, err := time.LoadLocation(p.Location)
	if err != nil {
		return "", err
	}
	return p.T.In(jakartaLoc).Format(p.Format), nil
}

func FormatDate(t time.Time) (string, error) {
	return Format(FormatParam{
		T:        t,
		Location: "Asia/Jakarta",
		Format:   "02 Jan 2006 15:04 MST",
	})
}

func FormatMySQLDateJakartaTimezone(t time.Time) (string, error) {
	return Format(FormatParam{
		T:        t,
		Location: "Asia/Jakarta",
		Format:   "2006-01-02 15:04:05",
	})
}

func FormatMySQLDateUTCTimezone(t time.Time) (string, error) {
	return Format(FormatParam{
		T:        t,
		Location: "UTC",
		Format:   "2006-01-02 15:04:05",
	})
}

func GetNthDay(t time.Time) (int, error) {
	u := t.Unix() + (7 * 3600)
	i := float64(u) / 86400
	return int(math.Floor(i)), nil
}

type FloorDayParam struct {
	T      time.Time
	DRange int
}

func FloorDay(p FloorDayParam) (time.Time, error) {
	i, err := GetNthDay(p.T)
	sec := ((i + p.DRange) * 86400) - (7 * 3600)
	return time.Unix(int64(sec), 0), err
}

func DayDiff(t1 time.Time, t2 time.Time) (int, error) {
	nthDay1, err := GetNthDay(t1)
	if err != nil {
		return 0, err
	}
	nthDay2, err := GetNthDay(t2)
	if err != nil {
		return 0, err
	}
	return nthDay2 - nthDay1, nil
}

type Range struct {
	Value       int
	IsEqual     bool
	IsSkipCheck bool
}

func IsInDayRange(t time.Time, now time.Time, minD Range, maxD Range) (bool, error) {
	diff, err := DayDiff(t, now)
	if err != nil {
		return false, err
	}

	isMinValid := false
	if !minD.IsSkipCheck {
		if minD.IsEqual {
			isMinValid = diff >= minD.Value
		} else {
			isMinValid = diff > minD.Value
		}
	} else {
		isMinValid = true
	}

	isMaxValid := false
	if !maxD.IsSkipCheck {
		if maxD.IsEqual {
			isMaxValid = diff <= maxD.Value
		} else {
			isMaxValid = diff < maxD.Value
		}
	} else {
		isMaxValid = true
	}

	return isMinValid && isMaxValid, nil
}

func IsInHourRange(t time.Time, now time.Time, minD Range, maxD Range) (bool, error) {
	var diff float32
	diff = float32(now.Unix()-t.Unix()) / 3600

	isMinValid := false
	if !minD.IsSkipCheck {
		if minD.IsEqual {
			isMinValid = diff >= float32(minD.Value)
		} else {
			isMinValid = diff > float32(minD.Value)
		}
	} else {
		isMinValid = true
	}

	isMaxValid := false
	if !maxD.IsSkipCheck {
		if maxD.IsEqual {
			isMaxValid = diff <= float32(maxD.Value)
		} else {
			isMaxValid = diff < float32(maxD.Value)
		}
	} else {
		isMaxValid = true
	}

	return isMinValid && isMaxValid, nil
}

func IsInMinuteRange(t time.Time, now time.Time, minM Range, maxM Range) (bool, error) {
	var diff float32
	diff = float32(now.Unix()-t.Unix()) / 60

	isMinValid := false
	if !minM.IsSkipCheck {
		if minM.IsEqual {
			isMinValid = diff >= float32(minM.Value)
		} else {
			isMinValid = diff > float32(minM.Value)
		}
	} else {
		isMinValid = true
	}

	isMaxValid := false
	if !maxM.IsSkipCheck {
		if maxM.IsEqual {
			isMaxValid = diff <= float32(maxM.Value)
		} else {
			isMaxValid = diff < float32(maxM.Value)
		}
	} else {
		isMaxValid = true
	}

	return isMinValid && isMaxValid, nil
}

type TimeRange struct {
	Start time.Time
	End   time.Time
}

func IsInDayRangeStartEnd(r TimeRange, now time.Time, minD Range, maxD Range) (bool, error) {
	startD, err := GetNthDay(r.Start)
	if err != nil {
		return false, err
	}
	endD, err := GetNthDay(r.End)
	if err != nil {
		return false, err
	}
	nowD, err := GetNthDay(now)
	if err != nil {
		return false, err
	}

	isMinValid := false
	if !minD.IsSkipCheck {
		startD = startD + minD.Value
		if minD.IsEqual {
			isMinValid = nowD >= startD
		} else {
			isMinValid = nowD > startD
		}
	} else {
		isMinValid = true
	}

	isMaxValid := false
	if !maxD.IsSkipCheck {
		endD = endD + maxD.Value
		if maxD.IsEqual {
			isMaxValid = nowD <= endD
		} else {
			isMaxValid = nowD < endD
		}
	} else {
		isMaxValid = true
	}

	return isMinValid && isMaxValid, nil
}

// CombineDateAndHour param: hourStr: "23:00:00"
func CombineDateAndHour(d time.Time, hourStr string) (time.Time, error) {
	hourTime, err := time.Parse(time.RFC3339, fmt.Sprintf("1970-01-01T%sZ", hourStr))
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(d.Unix()+hourTime.Unix(), 0), nil
}
