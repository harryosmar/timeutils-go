package timeutils_go_test

import (
	"fmt"
	timeutilsgo "github.com/harryosmar/timeutils-go"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestGetMonthRange(t *testing.T) {
	tests := []struct {
		name     string
		month    int
		year     int
		expected string
	}{
		{
			name:     "TC1",
			month:    1,
			year:     2018,
			expected: "2018-01-01T00:00:00+07:00 - 2018-01-31T23:59:59+07:00",
		},
		{
			name:     "TC2",
			month:    10,
			year:     2018,
			expected: "2018-10-01T00:00:00+07:00 - 2018-10-31T23:59:59+07:00",
		},
		{
			name:     "TC3",
			month:    12,
			year:     2018,
			expected: "2018-12-01T00:00:00+07:00 - 2018-12-31T23:59:59+07:00",
		},
		{
			name:     "TC3",
			month:    11,
			year:     2018,
			expected: "2018-11-01T00:00:00+07:00 - 2018-11-30T23:59:59+07:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gte, lte, err := timeutilsgo.GetMonthRange(tt.month, tt.year)
			assert.NoError(t, err)
			actual := fmt.Sprintf("%s - %s", time.Unix(gte, 0).Format(time.RFC3339), time.Unix(lte, 0).Format(time.RFC3339))
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestDayInUnixJakartaTimezone(t *testing.T) {
	type args struct {
		t time.Time
	}
	testData := []struct {
		name           string
		args           args
		expectedResult float64
	}{
		{
			name: "1.",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-01T00:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: 0,
		},
		{
			name: "2.",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-01T00:00:01+07:00")
					return parse
				}(),
			},
			expectedResult: 0,
		},
		{
			name: "3.",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-01T23:59:59+07:00")
					return parse
				}(),
			},
			expectedResult: 0,
		},
		{
			name: "4.",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-02T00:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: 1,
		},
		{
			name: "4.1",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-02T00:00:01+07:00")
					return parse
				}(),
			},
			expectedResult: 1,
		},
		{
			name: "4.2",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-02T23:59:59+07:00")
					return parse
				}(),
			},
			expectedResult: 1,
		},
		{

			name: "5.1",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-03T00:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: 2,
		},
		{

			name: "5.2",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-03T00:00:01+07:00")
					return parse
				}(),
			},
			expectedResult: 2,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			actual := timeutilsgo.DayInUnixJakartaTimezone(tt.args.t)
			if actual != tt.expectedResult {
				t.Errorf("expect %v got %v", tt.expectedResult, actual)
			}
		})
	}
}

func TestHourInUnixJakartaTimezone(t *testing.T) {
	type args struct {
		t time.Time
	}
	testData := []struct {
		name           string
		args           args
		expectedResult float64
	}{
		{
			name: "1.",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-01T00:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: 0,
		},
		{
			name: "2.",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-01T00:00:01+07:00")
					return parse
				}(),
			},
			expectedResult: 0,
		},
		{
			name: "2.1",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-01T00:59:59+07:00")
					return parse
				}(),
			},
			expectedResult: 0,
		},
		{
			name: "3.",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-01T01:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: 1,
		},
		{
			name: "3.1",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-01T01:59:59+07:00")
					return parse
				}(),
			},
			expectedResult: 1,
		},
		{

			name: "4.",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-02T00:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: 24,
		},
		{

			name: "4.1",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-02T00:59:59+07:00")
					return parse
				}(),
			},
			expectedResult: 24,
		},
		{

			name: "4.2",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-02T01:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: 25,
		},
		{

			name: "4.3",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-02T01:59:59+07:00")
					return parse
				}(),
			},
			expectedResult: 25,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			actual := timeutilsgo.HourInUnixJakartaTimezone(tt.args.t)
			if actual != tt.expectedResult {
				t.Errorf("expect %v got %v", tt.expectedResult, actual)
			}
		})
	}
}

func TestPlusHourToTime(t *testing.T) {
	type args struct {
		t time.Time
		h int64
	}
	testData := []struct {
		name           string
		args           args
		expectedResult time.Time
	}{
		{
			name: "plus 1 00:00",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-01T00:00:00+07:00")
					return parse
				}(),
				h: 1,
			},
			expectedResult: func() time.Time {
				parse, _ := time.Parse(time.RFC3339, "1970-01-01T01:00:00+07:00")
				return parse
			}(),
		},
		{
			name: "plus 1 59:59",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2020-03-19T00:59:59+07:00")
					return parse
				}(),
				h: 1,
			},
			expectedResult: func() time.Time {
				parse, _ := time.Parse(time.RFC3339, "2020-03-19T01:00:00+07:00")
				return parse
			}(),
		},
		{
			name: "plus 2 59:59",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2020-03-19T00:59:59+07:00")
					return parse
				}(),
				h: 2,
			},
			expectedResult: func() time.Time {
				parse, _ := time.Parse(time.RFC3339, "2020-03-19T02:00:00+07:00")
				return parse
			}(),
		},
		{
			name: "minus 1 00:00",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2020-03-19T00:59:59+07:00")
					return parse
				}(),
				h: -1,
			},
			expectedResult: func() time.Time {
				parse, _ := time.Parse(time.RFC3339, "2020-03-18T23:00:00+07:00")
				return parse
			}(),
		},
		{
			name: "minus 1 59:59",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2020-03-19T00:00:00+07:00")
					return parse
				}(),
				h: -1,
			},
			expectedResult: func() time.Time {
				parse, _ := time.Parse(time.RFC3339, "2020-03-18T23:00:00+07:00")
				return parse
			}(),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			//format := timeutils_go.PlusHourToTime(tt.args.t, tt.args.h).Format("15:04")
			//log.Println(format)

			actual := timeutilsgo.PlusHourToTime(tt.args.t, tt.args.h)
			log.Println(actual)
			if actual.Unix() != tt.expectedResult.Unix() {
				t.Errorf("expect %v got %v", tt.expectedResult, actual)
			}
		})
	}
}

func TestGetExpirationTillEndOfTodayJakartaTimezone(t *testing.T) {
	type args struct {
		t time.Time
	}
	testData := []struct {
		name           string
		args           args
		expectedResult int64
	}{
		{
			name: "1.",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-01T00:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: 86400,
		},
		{
			name: "2.",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-01T20:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: 14400,
		},
		{
			name: "3.",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-01T23:59:59+07:00")
					return parse
				}(),
			},
			expectedResult: 1,
		},
		{
			name: "4.",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-02T23:59:50+07:00")
					return parse
				}(),
			},
			expectedResult: 10,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			actual := timeutilsgo.GetExpirationTillEndOfTodayJakartaTimezone(tt.args.t)
			if actual != tt.expectedResult {
				t.Errorf("expect %v got %v", tt.expectedResult, actual)
			}
		})
	}
}

func TestFormatDate(t *testing.T) {
	type args struct {
		t time.Time
	}
	testData := []struct {
		name           string
		args           args
		expectedResult string
	}{
		{
			name: "tc1",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: "28 Mar 2023 00:00 WIB",
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := timeutilsgo.FormatDate(tt.args.t)
			if err != nil {
				t.Error(err)
				return
			}
			log.Println(actual)
			if actual != tt.expectedResult {
				t.Errorf("expect %v got %v", tt.expectedResult, actual)
			}
		})
	}
}

func TestFormatMySQLDateJakartaTimezone(t *testing.T) {
	type args struct {
		t time.Time
	}
	testData := []struct {
		name           string
		args           args
		expectedResult string
	}{
		{
			name: "param using jakarta tz",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T08:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: "2023-03-28 08:00:00",
		},
		{
			name: "param using UTC tz",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T08:00:00Z")
					return parse
				}(),
			},
			expectedResult: "2023-03-28 15:00:00",
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := timeutilsgo.FormatMySQLDateJakartaTimezone(tt.args.t)
			if err != nil {
				t.Error(err)
				return
			}
			log.Println(actual)
			if actual != tt.expectedResult {
				t.Errorf("expect %v got %v", tt.expectedResult, actual)
			}
		})
	}
}

func TestFormatMySQLDateUTCTimezone(t *testing.T) {
	type args struct {
		t time.Time
	}
	testData := []struct {
		name           string
		args           args
		expectedResult string
	}{
		{
			name: "param using jakarta tz",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T08:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: "2023-03-28 01:00:00",
		},
		{
			name: "param using UTC tz",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T08:00:00Z")
					return parse
				}(),
			},
			expectedResult: "2023-03-28 08:00:00",
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := timeutilsgo.FormatMySQLDateUTCTimezone(tt.args.t)
			if err != nil {
				t.Error(err)
				return
			}
			log.Println(actual)
			if actual != tt.expectedResult {
				t.Errorf("expect %v got %v", tt.expectedResult, actual)
			}
		})
	}
}

func TestGetNthDay(t *testing.T) {
	type args struct {
		t time.Time
	}
	testData := []struct {
		name           string
		args           args
		expectedResult int
	}{
		{
			name: "tc1",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: 19444,
		},
		{
			name: "tc2",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T23:59:59+07:00")
					return parse
				}(),
			},
			expectedResult: 19444,
		},
		{
			name: "0d",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-01T00:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: 0,
		},
		{
			name: "0d",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-01T00:00:01+07:00")
					return parse
				}(),
			},
			expectedResult: 0,
		},
		{
			name: "0d",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-01T23:59:59+07:00")
					return parse
				}(),
			},
			expectedResult: 0,
		},
		{
			name: "1d",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-02T00:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: 1,
		},
		{
			name: "1d",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-02T23:59:59+07:00")
					return parse
				}(),
			},
			expectedResult: 1,
		},
		{
			name: "2d",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-03T00:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: 2,
		},
		{
			name: "2d",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "1970-01-03T23:59:59+07:00")
					return parse
				}(),
			},
			expectedResult: 2,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := timeutilsgo.GetNthDay(tt.args.t)
			if err != nil {
				t.Error(err)
				return
			}
			log.Println(actual)
			if actual != tt.expectedResult {
				t.Errorf("expect %v got %v", tt.expectedResult, actual)
			}
		})
	}
}

func TestFloorDay(t *testing.T) {
	type args struct {
		t      time.Time
		dRange int
	}
	testData := []struct {
		name           string
		args           args
		expectedResult time.Time
	}{
		{
			name: "dRange 0  00:00:00",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				dRange: 0,
			},
			expectedResult: func() time.Time {
				parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
				return parse
			}(),
		},
		{
			name: "dRange 0 23:59:59",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T23:59:59+07:00")
					return parse
				}(),
				dRange: 0,
			},
			expectedResult: func() time.Time {
				parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
				return parse
			}(),
		},
		{
			name: "dRange 4 23:59:59",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T23:59:59+07:00")
					return parse
				}(),
				dRange: 4,
			},
			expectedResult: func() time.Time {
				parse, _ := time.Parse(time.RFC3339, "2023-04-01T00:00:00+07:00")
				return parse
			}(),
		},
		{
			name: "dRange -2 23:59:59",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T23:59:59+07:00")
					return parse
				}(),
				dRange: -2,
			},
			expectedResult: func() time.Time {
				parse, _ := time.Parse(time.RFC3339, "2023-03-26T00:00:00+07:00")
				return parse
			}(),
		},
		{
			name: "dRange -5 23:59:59",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-02T23:59:59+07:00")
					return parse
				}(),
				dRange: -5,
			},
			expectedResult: func() time.Time {
				parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
				return parse
			}(),
		},
		{
			name: "dRange -5 23:59:59",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-15T23:59:59+07:00")
					return parse
				}(),
				dRange: -14,
			},
			expectedResult: func() time.Time {
				parse, _ := time.Parse(time.RFC3339, "2023-04-01T00:00:00+07:00")
				return parse
			}(),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := timeutilsgo.FloorDay(timeutilsgo.FloorDayParam{
				T:      tt.args.t,
				DRange: tt.args.dRange,
			})
			if err != nil {
				t.Error(err)
				return
			}
			log.Println(actual)
			if actual.Unix() != tt.expectedResult.Unix() {
				t.Errorf("expect %v got %v", tt.expectedResult, actual)
			}
		})
	}
}

func TestDayDiff(t *testing.T) {
	type args struct {
		t1 time.Time
		t2 time.Time
	}
	testData := []struct {
		name           string
		args           args
		expectedResult int
	}{
		{
			name: "tc0",
			args: args{
				t1: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				t2: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-27T23:59:59+07:00")
					return parse
				}(),
			},
			expectedResult: -1,
		},
		{
			name: "tc1",
			args: args{
				t1: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				t2: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: 0,
		},
		{
			name: "tc2",
			args: args{
				t1: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				t2: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T23:59:59+07:00")
					return parse
				}(),
			},
			expectedResult: 0,
		},
		{
			name: "tc3",
			args: args{
				t1: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				t2: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-29T00:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: 1,
		},
		{
			name: "3 23:59:59",
			args: args{
				t1: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-04T23:59:59+07:00")
					return parse
				}(),
				t2: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-07T23:59:59+07:00")
					return parse
				}(),
			},
			expectedResult: 3,
		},
		{
			name: "3 00:00:00",
			args: args{
				t1: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-04T23:59:59+07:00")
					return parse
				}(),
				t2: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-07T00:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: 3,
		},
		{
			name: "2 23:59:59",
			args: args{
				t1: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-04T23:59:59+07:00")
					return parse
				}(),
				t2: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-06T23:59:59+07:00")
					return parse
				}(),
			},
			expectedResult: 2,
		},
		{
			name: "2 00:00:00",
			args: args{
				t1: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-04T23:59:59+07:00")
					return parse
				}(),
				t2: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-06T00:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: 2,
		},
		{
			name: "1 23:59:59",
			args: args{
				t1: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-04T23:59:59+07:00")
					return parse
				}(),
				t2: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-05T23:59:59+07:00")
					return parse
				}(),
			},
			expectedResult: 1,
		},
		{
			name: "0 00:00:00",
			args: args{
				t1: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-04T23:59:59+07:00")
					return parse
				}(),
				t2: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-05T00:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: 1,
		},
		{
			name: "0 00:00:00",
			args: args{
				t1: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-04T23:59:59+07:00")
					return parse
				}(),
				t2: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-04T00:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: 0,
		},
		{
			name: "-1 23:59:59",
			args: args{
				t1: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-04T23:59:59+07:00")
					return parse
				}(),
				t2: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-03T23:59:59+07:00")
					return parse
				}(),
			},
			expectedResult: -1,
		},
		{
			name: "-1 00:00:00",
			args: args{
				t1: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-04T23:59:59+07:00")
					return parse
				}(),
				t2: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-03T00:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: -1,
		},
		{
			name: "-2 23:59:59",
			args: args{
				t1: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-04T23:59:59+07:00")
					return parse
				}(),
				t2: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-02T23:59:59+07:00")
					return parse
				}(),
			},
			expectedResult: -2,
		},
		{
			name: "-2 00:00:00",
			args: args{
				t1: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-04T23:59:59+07:00")
					return parse
				}(),
				t2: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-02T00:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: -2,
		},
		{
			name: "-3 23:59:59",
			args: args{
				t1: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-04T23:59:59+07:00")
					return parse
				}(),
				t2: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-01T23:59:59+07:00")
					return parse
				}(),
			},
			expectedResult: -3,
		},
		{
			name: "-3 00:00:00",
			args: args{
				t1: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-04T23:59:59+07:00")
					return parse
				}(),
				t2: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-01T00:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: -3,
		},
		{
			name: "-4 23:59:59",
			args: args{
				t1: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-04T23:59:59+07:00")
					return parse
				}(),
				t2: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-31T23:59:59+07:00")
					return parse
				}(),
			},
			expectedResult: -4,
		},
		{
			name: "-4 00:00:00",
			args: args{
				t1: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-04-04T23:59:59+07:00")
					return parse
				}(),
				t2: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-31T00:00:00+07:00")
					return parse
				}(),
			},
			expectedResult: -4,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := timeutilsgo.DayDiff(tt.args.t1, tt.args.t2)
			if err != nil {
				t.Error(err)
				return
			}
			log.Println(actual)
			if actual != tt.expectedResult {
				t.Errorf("expect %v got %v", tt.expectedResult, actual)
			}
		})
	}
}

func TestIsInDayRange(t *testing.T) {
	type args struct {
		t   time.Time
		now time.Time
		min timeutilsgo.Range
		max timeutilsgo.Range
	}
	testData := []struct {
		name           string
		args           args
		expectedResult bool
	}{
		{
			name: "Given rule t <= now <= t + 1 day and now = t Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   0,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   1,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule t <= now <= t + 1 day and t+1 = now Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-29T00:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   0,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   1,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule t <= now < t + 1 day and t+1 = now Then expect return false",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-29T00:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   0,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   1,
					IsEqual: false,
				},
			},
			expectedResult: false,
		},
		{
			name: "Given rule t <= now < t + 1 day and t = now and time 23:59:59 Then expect return false",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T23:59:59+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   0,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   1,
					IsEqual: false,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule t <= now <= t + 1 day and t + 1 = now and time 23:59:59 Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-29T23:59:59+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   0,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   1,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule t <= now <= t + 1 day and t + 2 = now Then expect return false",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-30T00:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   0,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   1,
					IsEqual: true,
				},
			},
			expectedResult: false,
		},
		{
			name: "Given rule t <= now <= t + 1 day and t - 1= now Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-27T00:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -1,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   1,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule t <= now <= t + 1 day and t - 1= now and time 23:59:59 Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-27T23:59:59+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -1,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   1,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule t < now <= t + 1 day and t - 1= now Then expect return false",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-27T00:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -1,
					IsEqual: false,
				},
				max: timeutilsgo.Range{
					Value:   1,
					IsEqual: true,
				},
			},
			expectedResult: false,
		},
		{
			name: "Given rule t < now <= t + 1 day and t - 1= now and time 23:59:59 Then expect return false",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-27T23:59:59+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -1,
					IsEqual: false,
				},
				max: timeutilsgo.Range{
					Value:   1,
					IsEqual: true,
				},
			},
			expectedResult: false,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := timeutilsgo.IsInDayRange(tt.args.t, tt.args.now, tt.args.min, tt.args.max)
			if err != nil {
				t.Error(err)
				return
			}
			log.Println(actual)
			if actual != tt.expectedResult {
				t.Errorf("expect %v got %v", tt.expectedResult, actual)
			}
		})
	}
}

func TestIsInHourRange(t *testing.T) {
	type args struct {
		t   time.Time
		now time.Time
		min timeutilsgo.Range
		max timeutilsgo.Range
	}
	testData := []struct {
		name           string
		args           args
		expectedResult bool
	}{
		{
			name: "Given rule 1 - t <= now <= t + 2 hour and now = t Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -1,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   2,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule 1 - t <= now <= t + 2 hour and now = t + 1 Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T01:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -1,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   2,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule 1 - t <= now <= t + 2 hour and now = t + 2 Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T02:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -1,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   2,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule 1 - t <= now < t + 2 hour and now = t + 1 and min:seconds 59:59 Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T01:59:59+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -1,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   2,
					IsEqual: false,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule 1 - t <= now < t + 2 hour and now = t + 2 Then expect return false",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T02:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -1,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   2,
					IsEqual: false,
				},
			},
			expectedResult: false,
		},
		{
			name: "Given rule 1 - t <= now <= t + 2 hour and now = t + 3 Then expect return false",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T03:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -1,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   2,
					IsEqual: true,
				},
			},
			expectedResult: false,
		},
		{
			name: "Given rule 1 - t <= now <= t + 2 hour and now = t - 1 Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-27T23:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -1,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   2,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule 1 - t <= now <= t + 2 hour and now = t - 2 Then expect return false",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-27T22:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -1,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   2,
					IsEqual: true,
				},
			},
			expectedResult: false,
		},
		{
			name: "Given rule 1 - t <= now <= t + 2 hour and now = t - 2 and minute:seconds 59:59 Then expect return false",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-27T22:59:59+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -1,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   2,
					IsEqual: true,
				},
			},
			expectedResult: false,
		},
		{
			name: "Given rule 1 - t <= now <= t + 2 hour and now = t and minute:second 59:59 Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-27T23:59:59+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -1,
					IsEqual: false,
				},
				max: timeutilsgo.Range{
					Value:   2,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule 1 - t <= now <= t + 2 hour and now = t - 1 Then expect return false",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-27T23:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -1,
					IsEqual: false,
				},
				max: timeutilsgo.Range{
					Value:   2,
					IsEqual: true,
				},
			},
			expectedResult: false,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := timeutilsgo.IsInHourRange(
				tt.args.t,
				tt.args.now,
				tt.args.min,
				tt.args.max,
			)
			if err != nil {
				t.Error(err)
				return
			}
			log.Println(actual)
			if actual != tt.expectedResult {
				t.Errorf("expect %v got %v", tt.expectedResult, actual)
			}
		})
	}
}

func TestIsInMinuteRange(t *testing.T) {
	type args struct {
		t   time.Time
		now time.Time
		min timeutilsgo.Range
		max timeutilsgo.Range
	}
	testData := []struct {
		name           string
		args           args
		expectedResult bool
	}{
		// now < t - 5m
		{
			name: "Given rule now < t - 5m and now = t Then expect return false",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					IsSkipCheck: true,
				},
				max: timeutilsgo.Range{
					Value:   -5,
					IsEqual: false,
				},
			},
			expectedResult: false,
		},
		{
			name: "Given rule now < t - 5m and now = t - 5m Then expect return false",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-27T23:55:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					IsSkipCheck: true,
				},
				max: timeutilsgo.Range{
					Value:   -5,
					IsEqual: false,
				},
			},
			expectedResult: false,
		},
		{
			name: "Given rule now < t - 5m and now = t - 5m - 1s Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-27T23:54:59+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					IsSkipCheck: true,
				},
				max: timeutilsgo.Range{
					Value:   -5,
					IsEqual: false,
				},
			},
			expectedResult: true,
		},
		// now <= t - 5m
		{
			name: "Given rule now <= t - 5m and now = t - 5m Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-27T23:55:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					IsSkipCheck: true,
				},
				max: timeutilsgo.Range{
					Value:   -5,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule now <= t - 5m and now = t - 5m + 1s Then expect return false",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-27T23:55:01+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					IsSkipCheck: true,
				},
				max: timeutilsgo.Range{
					Value:   -5,
					IsEqual: true,
				},
			},
			expectedResult: false,
		},
		// now > t + 5m
		{
			name: "Given rule now > t + 5m and now = t Then expect return false",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   5,
					IsEqual: false,
				},
				max: timeutilsgo.Range{
					IsSkipCheck: true,
				},
			},
			expectedResult: false,
		},
		{
			name: "Given rule now > t + 5m and now = t + 5m Then expect return false",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:05:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   5,
					IsEqual: false,
				},
				max: timeutilsgo.Range{
					IsSkipCheck: true,
				},
			},
			expectedResult: false,
		},
		{
			name: "Given rule now > t + 5m and now = t + 5m + 1s Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T05:00:01+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   5,
					IsEqual: false,
				},
				max: timeutilsgo.Range{
					IsSkipCheck: true,
				},
			},
			expectedResult: true,
		},
		// now >= t - 5m
		{
			name: "Given rule now >= t + 5m and now = t + 5m Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:05:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   5,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					IsSkipCheck: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule now >= t + 5m and now = t + 5m - 1s Then expect return false",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:04:59+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   5,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					IsSkipCheck: true,
				},
			},
			expectedResult: false,
		},
		// t - 3m <= now <= t + 5m
		{
			name: "Given rule t - 3m <= now <= t + 5m and now = t - 3m - 1s Then expect return false",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-27T23:56:59+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -3,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   5,
					IsEqual: true,
				},
			},
			expectedResult: false,
		},
		{
			name: "Given rule t - 3m <= now <= t + 5m and now = t - 3m Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-27T23:57:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -3,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   5,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule t - 3m <= now <= t + 5m and now = t - 2m Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-27T23:58:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -3,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   5,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule t - 3m <= now <= t + 5m and now = t - 1m Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-27T23:59:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -3,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   5,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule t - 3m <= now <= t + 5m and now = t Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -3,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   5,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule t - 3m <= now <= t + 5m and now = t + 1m Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:01:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -3,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   5,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule t - 3m <= now <= t + 5m and now = t + 2m Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:02:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -3,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   5,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule t - 3m <= now <= t + 5m and now = t + 3m Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:03:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -3,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   5,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule t - 3m <= now <= t + 5m and now = t + 4m Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:04:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -3,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   5,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule t - 3m <= now <= t + 5m and now = t + 5m Then expect return true",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:05:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -3,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   5,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule t - 3m <= now <= t + 5m and now = t + 5m + 1s Then expect return false",
			args: args{
				t: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:05:01+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -3,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   5,
					IsEqual: true,
				},
			},
			expectedResult: false,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := timeutilsgo.IsInMinuteRange(
				tt.args.t,
				tt.args.now,
				tt.args.min,
				tt.args.max,
			)
			if err != nil {
				t.Error(err)
				return
			}
			log.Println(actual)
			if actual != tt.expectedResult {
				t.Errorf("expect %v got %v", tt.expectedResult, actual)
			}
		})
	}
}

func TestIsInDayRangeStartEnd(t *testing.T) {
	type args struct {
		t   timeutilsgo.TimeRange
		now time.Time
		min timeutilsgo.Range
		max timeutilsgo.Range
	}
	testData := []struct {
		name           string
		args           args
		expectedResult bool
	}{
		{
			name: "Given rule start + 0 <= now <= end + 0 day and now = start Then expect return true",
			args: args{
				t: timeutilsgo.TimeRange{
					Start: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-25T00:00:00+07:00")
						return parse
					}(),
					End: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
						return parse
					}(),
				},
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-25T00:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   0,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   0,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule start + 0 <= now <= end + 0 day and now = end Then expect return true",
			args: args{
				t: timeutilsgo.TimeRange{
					Start: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-25T00:00:00+07:00")
						return parse
					}(),
					End: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
						return parse
					}(),
				},
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   0,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   0,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule start + 0 <= now <= end + 0 day and now = start - 1 Then expect return false",
			args: args{
				t: timeutilsgo.TimeRange{
					Start: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-25T00:00:00+07:00")
						return parse
					}(),
					End: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
						return parse
					}(),
				},
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-24T00:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   0,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   0,
					IsEqual: true,
				},
			},
			expectedResult: false,
		},
		{
			name: "Given rule start + 0 <= now <= end + 0 day and now = end + 1 Then expect return false",
			args: args{
				t: timeutilsgo.TimeRange{
					Start: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-25T00:00:00+07:00")
						return parse
					}(),
					End: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
						return parse
					}(),
				},
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-29T00:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   0,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   0,
					IsEqual: true,
				},
			},
			expectedResult: false,
		},
		{
			name: "Given rule start + 0 < now < end + 0 day and now = start Then expect return false",
			args: args{
				t: timeutilsgo.TimeRange{
					Start: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-25T00:00:00+07:00")
						return parse
					}(),
					End: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
						return parse
					}(),
				},
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-25T00:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   0,
					IsEqual: false,
				},
				max: timeutilsgo.Range{
					Value:   0,
					IsEqual: false,
				},
			},
			expectedResult: false,
		},
		{
			name: "Given rule start + 0 < now < end + 0 day and now = end Then expect return false",
			args: args{
				t: timeutilsgo.TimeRange{
					Start: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-25T00:00:00+07:00")
						return parse
					}(),
					End: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
						return parse
					}(),
				},
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   0,
					IsEqual: false,
				},
				max: timeutilsgo.Range{
					Value:   0,
					IsEqual: false,
				},
			},
			expectedResult: false,
		},
		{
			name: "Given rule start + 0 < now < end + 0 day and now = end - 1second Then expect return true",
			args: args{
				t: timeutilsgo.TimeRange{
					Start: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-25T00:00:00+07:00")
						return parse
					}(),
					End: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
						return parse
					}(),
				},
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-27T23:59:59+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   0,
					IsEqual: false,
				},
				max: timeutilsgo.Range{
					Value:   0,
					IsEqual: false,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule start - 1 <= now <= end + 2 day and now = start - 1 Then expect return true",
			args: args{
				t: timeutilsgo.TimeRange{
					Start: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-25T00:00:00+07:00")
						return parse
					}(),
					End: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
						return parse
					}(),
				},
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-24T00:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -1,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   2,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule start - 1 <= now <= end + 2 day and now = start - 2 Then expect return false",
			args: args{
				t: timeutilsgo.TimeRange{
					Start: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-25T00:00:00+07:00")
						return parse
					}(),
					End: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
						return parse
					}(),
				},
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-23T00:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -1,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   2,
					IsEqual: true,
				},
			},
			expectedResult: false,
		},
		{
			name: "Given rule start - 1 <= now <= end + 2 day and now = end + 1 Then expect return true",
			args: args{
				t: timeutilsgo.TimeRange{
					Start: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-25T00:00:00+07:00")
						return parse
					}(),
					End: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
						return parse
					}(),
				},
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-29T00:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -1,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   2,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule start - 1 <= now <= end + 2 day and now = end + 2 Then expect return true",
			args: args{
				t: timeutilsgo.TimeRange{
					Start: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-25T00:00:00+07:00")
						return parse
					}(),
					End: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
						return parse
					}(),
				},
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-30T00:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -1,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   2,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule start - 1 <= now <= end + 2 day and now = end + 2 with time 23:59:59 Then expect return true",
			args: args{
				t: timeutilsgo.TimeRange{
					Start: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-25T00:00:00+07:00")
						return parse
					}(),
					End: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
						return parse
					}(),
				},
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-30T23:59:59+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -1,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   2,
					IsEqual: true,
				},
			},
			expectedResult: true,
		},
		{
			name: "Given rule start - 1 <= now <= end + 2 day and now = end + 3 Then expect return false",
			args: args{
				t: timeutilsgo.TimeRange{
					Start: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-25T00:00:00+07:00")
						return parse
					}(),
					End: func() time.Time {
						parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
						return parse
					}(),
				},
				now: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-31T00:00:00+07:00")
					return parse
				}(),
				min: timeutilsgo.Range{
					Value:   -1,
					IsEqual: true,
				},
				max: timeutilsgo.Range{
					Value:   2,
					IsEqual: true,
				},
			},
			expectedResult: false,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := timeutilsgo.IsInDayRangeStartEnd(tt.args.t, tt.args.now, tt.args.min, tt.args.max)
			if err != nil {
				t.Error(err)
				return
			}
			log.Println(actual)
			if actual != tt.expectedResult {
				t.Errorf("expect %v got %v", tt.expectedResult, actual)
			}
		})
	}
}

func TestCombineDateAndHour(t *testing.T) {
	type args struct {
		date    time.Time
		hourStr string
	}
	testData := []struct {
		name           string
		args           args
		expectedResult time.Time
	}{
		{
			name: "tc1",
			args: args{
				date: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				hourStr: "17:00:00",
			},
			expectedResult: func() time.Time {
				parse, _ := time.Parse(time.RFC3339, "2023-03-28T17:00:00+07:00")
				return parse
			}(),
		},
		{
			name: "tc2",
			args: args{
				date: func() time.Time {
					parse, _ := time.Parse(time.RFC3339, "2023-03-28T00:00:00+07:00")
					return parse
				}(),
				hourStr: "23:05:45",
			},
			expectedResult: func() time.Time {
				parse, _ := time.Parse(time.RFC3339, "2023-03-28T23:05:45+07:00")
				return parse
			}(),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := timeutilsgo.CombineDateAndHour(tt.args.date, tt.args.hourStr)
			if err != nil {
				t.Error(err)
				return
			}
			if actual.Unix() != tt.expectedResult.Unix() {
				t.Errorf("expect %+v got %+v", tt.expectedResult, actual)
				return
			}
		})
	}
}
