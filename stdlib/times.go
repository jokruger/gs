package stdlib

import (
	"time"

	gse "github.com/jokruger/gs/error"
	gst "github.com/jokruger/gs/types"
)

var timesModule = map[string]gst.Object{
	"format_ansic":        &gst.String{Value: time.ANSIC},
	"format_unix_date":    &gst.String{Value: time.UnixDate},
	"format_ruby_date":    &gst.String{Value: time.RubyDate},
	"format_rfc822":       &gst.String{Value: time.RFC822},
	"format_rfc822z":      &gst.String{Value: time.RFC822Z},
	"format_rfc850":       &gst.String{Value: time.RFC850},
	"format_rfc1123":      &gst.String{Value: time.RFC1123},
	"format_rfc1123z":     &gst.String{Value: time.RFC1123Z},
	"format_rfc3339":      &gst.String{Value: time.RFC3339},
	"format_rfc3339_nano": &gst.String{Value: time.RFC3339Nano},
	"format_kitchen":      &gst.String{Value: time.Kitchen},
	"format_stamp":        &gst.String{Value: time.Stamp},
	"format_stamp_milli":  &gst.String{Value: time.StampMilli},
	"format_stamp_micro":  &gst.String{Value: time.StampMicro},
	"format_stamp_nano":   &gst.String{Value: time.StampNano},
	"nanosecond":          &gst.Int{Value: int64(time.Nanosecond)},
	"microsecond":         &gst.Int{Value: int64(time.Microsecond)},
	"millisecond":         &gst.Int{Value: int64(time.Millisecond)},
	"second":              &gst.Int{Value: int64(time.Second)},
	"minute":              &gst.Int{Value: int64(time.Minute)},
	"hour":                &gst.Int{Value: int64(time.Hour)},
	"january":             &gst.Int{Value: int64(time.January)},
	"february":            &gst.Int{Value: int64(time.February)},
	"march":               &gst.Int{Value: int64(time.March)},
	"april":               &gst.Int{Value: int64(time.April)},
	"may":                 &gst.Int{Value: int64(time.May)},
	"june":                &gst.Int{Value: int64(time.June)},
	"july":                &gst.Int{Value: int64(time.July)},
	"august":              &gst.Int{Value: int64(time.August)},
	"september":           &gst.Int{Value: int64(time.September)},
	"october":             &gst.Int{Value: int64(time.October)},
	"november":            &gst.Int{Value: int64(time.November)},
	"december":            &gst.Int{Value: int64(time.December)},
	"sleep": &gst.UserFunction{
		Name:  "sleep",
		Value: timesSleep,
	}, // sleep(int)
	"parse_duration": &gst.UserFunction{
		Name:  "parse_duration",
		Value: timesParseDuration,
	}, // parse_duration(str) => int
	"since": &gst.UserFunction{
		Name:  "since",
		Value: timesSince,
	}, // since(time) => int
	"until": &gst.UserFunction{
		Name:  "until",
		Value: timesUntil,
	}, // until(time) => int
	"duration_hours": &gst.UserFunction{
		Name:  "duration_hours",
		Value: timesDurationHours,
	}, // duration_hours(int) => float
	"duration_minutes": &gst.UserFunction{
		Name:  "duration_minutes",
		Value: timesDurationMinutes,
	}, // duration_minutes(int) => float
	"duration_nanoseconds": &gst.UserFunction{
		Name:  "duration_nanoseconds",
		Value: timesDurationNanoseconds,
	}, // duration_nanoseconds(int) => int
	"duration_seconds": &gst.UserFunction{
		Name:  "duration_seconds",
		Value: timesDurationSeconds,
	}, // duration_seconds(int) => float
	"duration_string": &gst.UserFunction{
		Name:  "duration_string",
		Value: timesDurationString,
	}, // duration_string(int) => string
	"month_string": &gst.UserFunction{
		Name:  "month_string",
		Value: timesMonthString,
	}, // month_string(int) => string
	"date": &gst.UserFunction{
		Name:  "date",
		Value: timesDate,
	}, // date(year, month, day, hour, min, sec, nsec) => time
	"now": &gst.UserFunction{
		Name:  "now",
		Value: timesNow,
	}, // now() => time
	"parse": &gst.UserFunction{
		Name:  "parse",
		Value: timesParse,
	}, // parse(format, str) => time
	"unix": &gst.UserFunction{
		Name:  "unix",
		Value: timesUnix,
	}, // unix(sec, nsec) => time
	"add": &gst.UserFunction{
		Name:  "add",
		Value: timesAdd,
	}, // add(time, int) => time
	"add_date": &gst.UserFunction{
		Name:  "add_date",
		Value: timesAddDate,
	}, // add_date(time, years, months, days) => time
	"sub": &gst.UserFunction{
		Name:  "sub",
		Value: timesSub,
	}, // sub(t time, u time) => int
	"after": &gst.UserFunction{
		Name:  "after",
		Value: timesAfter,
	}, // after(t time, u time) => bool
	"before": &gst.UserFunction{
		Name:  "before",
		Value: timesBefore,
	}, // before(t time, u time) => bool
	"time_year": &gst.UserFunction{
		Name:  "time_year",
		Value: timesTimeYear,
	}, // time_year(time) => int
	"time_month": &gst.UserFunction{
		Name:  "time_month",
		Value: timesTimeMonth,
	}, // time_month(time) => int
	"time_day": &gst.UserFunction{
		Name:  "time_day",
		Value: timesTimeDay,
	}, // time_day(time) => int
	"time_weekday": &gst.UserFunction{
		Name:  "time_weekday",
		Value: timesTimeWeekday,
	}, // time_weekday(time) => int
	"time_hour": &gst.UserFunction{
		Name:  "time_hour",
		Value: timesTimeHour,
	}, // time_hour(time) => int
	"time_minute": &gst.UserFunction{
		Name:  "time_minute",
		Value: timesTimeMinute,
	}, // time_minute(time) => int
	"time_second": &gst.UserFunction{
		Name:  "time_second",
		Value: timesTimeSecond,
	}, // time_second(time) => int
	"time_nanosecond": &gst.UserFunction{
		Name:  "time_nanosecond",
		Value: timesTimeNanosecond,
	}, // time_nanosecond(time) => int
	"time_unix": &gst.UserFunction{
		Name:  "time_unix",
		Value: timesTimeUnix,
	}, // time_unix(time) => int
	"time_unix_nano": &gst.UserFunction{
		Name:  "time_unix_nano",
		Value: timesTimeUnixNano,
	}, // time_unix_nano(time) => int
	"time_format": &gst.UserFunction{
		Name:  "time_format",
		Value: timesTimeFormat,
	}, // time_format(time, format) => string
	"time_location": &gst.UserFunction{
		Name:  "time_location",
		Value: timesTimeLocation,
	}, // time_location(time) => string
	"time_string": &gst.UserFunction{
		Name:  "time_string",
		Value: timesTimeString,
	}, // time_string(time) => string
	"is_zero": &gst.UserFunction{
		Name:  "is_zero",
		Value: timesIsZero,
	}, // is_zero(time) => bool
	"to_local": &gst.UserFunction{
		Name:  "to_local",
		Value: timesToLocal,
	}, // to_local(time) => time
	"to_utc": &gst.UserFunction{
		Name:  "to_utc",
		Value: timesToUTC,
	}, // to_utc(time) => time
	"in_location": &gst.UserFunction{
		Name:  "in_location",
		Value: timesInLocation,
	}, // in_location(time, location) => time
}

func timesSleep(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	i1, ok := args[0].ToInt64()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	time.Sleep(time.Duration(i1))
	ret = gst.UndefinedValue

	return
}

func timesParseDuration(args ...gst.Object) (
	ret gst.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	s1, ok := args[0].ToString()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	dur, err := time.ParseDuration(s1)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &gst.Int{Value: int64(dur)}

	return
}

func timesSince(args ...gst.Object) (
	ret gst.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gst.Int{Value: int64(time.Since(t1))}

	return
}

func timesUntil(args ...gst.Object) (
	ret gst.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gst.Int{Value: int64(time.Until(t1))}

	return
}

func timesDurationHours(args ...gst.Object) (
	ret gst.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	i1, ok := args[0].ToInt64()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gst.Float{Value: time.Duration(i1).Hours()}

	return
}

func timesDurationMinutes(args ...gst.Object) (
	ret gst.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	i1, ok := args[0].ToInt64()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gst.Float{Value: time.Duration(i1).Minutes()}

	return
}

func timesDurationNanoseconds(args ...gst.Object) (
	ret gst.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	i1, ok := args[0].ToInt64()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gst.Int{Value: time.Duration(i1).Nanoseconds()}

	return
}

func timesDurationSeconds(args ...gst.Object) (
	ret gst.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	i1, ok := args[0].ToInt64()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gst.Float{Value: time.Duration(i1).Seconds()}

	return
}

func timesDurationString(args ...gst.Object) (
	ret gst.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	i1, ok := args[0].ToInt64()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gst.String{Value: time.Duration(i1).String()}

	return
}

func timesMonthString(args ...gst.Object) (
	ret gst.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	i1, ok := args[0].ToInt64()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gst.String{Value: time.Month(i1).String()}

	return
}

func timesDate(args ...gst.Object) (
	ret gst.Object,
	err error,
) {
	if len(args) < 7 || len(args) > 8 {
		err = gse.ErrWrongNumArguments
		return
	}

	i1, ok := args[0].ToInt()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}
	i2, ok := args[1].ToInt()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}
	i3, ok := args[2].ToInt()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}
	i4, ok := args[3].ToInt()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}
	i5, ok := args[4].ToInt()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "fifth",
			Expected: "int(compatible)",
			Found:    args[4].TypeName(),
		}
		return
	}
	i6, ok := args[5].ToInt()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "sixth",
			Expected: "int(compatible)",
			Found:    args[5].TypeName(),
		}
		return
	}
	i7, ok := args[6].ToInt()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "seventh",
			Expected: "int(compatible)",
			Found:    args[6].TypeName(),
		}
		return
	}

	var loc *time.Location
	if len(args) == 8 {
		i8, ok := args[7].ToString()
		if !ok {
			err = gse.ErrInvalidArgumentType{
				Name:     "eighth",
				Expected: "string(compatible)",
				Found:    args[7].TypeName(),
			}
			return
		}
		loc, err = time.LoadLocation(i8)
		if err != nil {
			ret = wrapError(err)
			return
		}
	} else {
		loc = time.Now().Location()
	}

	ret = &gst.Time{
		Value: time.Date(i1,
			time.Month(i2), i3, i4, i5, i6, i7, loc),
	}

	return
}

func timesNow(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 0 {
		err = gse.ErrWrongNumArguments
		return
	}

	ret = &gst.Time{Value: time.Now()}

	return
}

func timesParse(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 2 {
		err = gse.ErrWrongNumArguments
		return
	}

	s1, ok := args[0].ToString()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := args[1].ToString()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	parsed, err := time.Parse(s1, s2)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &gst.Time{Value: parsed}

	return
}

func timesUnix(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 2 {
		err = gse.ErrWrongNumArguments
		return
	}

	i1, ok := args[0].ToInt64()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := args[1].ToInt64()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &gst.Time{Value: time.Unix(i1, i2)}

	return
}

func timesAdd(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 2 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := args[1].ToInt64()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &gst.Time{Value: t1.Add(time.Duration(i2))}

	return
}

func timesSub(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 2 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := args[1].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &gst.Int{Value: int64(t1.Sub(t2))}

	return
}

func timesAddDate(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 4 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := args[1].ToInt()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	i3, ok := args[2].ToInt()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	i4, ok := args[3].ToInt()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}

	ret = &gst.Time{Value: t1.AddDate(i2, i3, i4)}

	return
}

func timesAfter(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 2 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := args[1].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	if t1.After(t2) {
		ret = gst.TrueValue
	} else {
		ret = gst.FalseValue
	}

	return
}

func timesBefore(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 2 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := args[1].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	if t1.Before(t2) {
		ret = gst.TrueValue
	} else {
		ret = gst.FalseValue
	}

	return
}

func timesTimeYear(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gst.Int{Value: int64(t1.Year())}

	return
}

func timesTimeMonth(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gst.Int{Value: int64(t1.Month())}

	return
}

func timesTimeDay(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gst.Int{Value: int64(t1.Day())}

	return
}

func timesTimeWeekday(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gst.Int{Value: int64(t1.Weekday())}

	return
}

func timesTimeHour(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gst.Int{Value: int64(t1.Hour())}

	return
}

func timesTimeMinute(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gst.Int{Value: int64(t1.Minute())}

	return
}

func timesTimeSecond(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gst.Int{Value: int64(t1.Second())}

	return
}

func timesTimeNanosecond(args ...gst.Object) (
	ret gst.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gst.Int{Value: int64(t1.Nanosecond())}

	return
}

func timesTimeUnix(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gst.Int{Value: t1.Unix()}

	return
}

func timesTimeUnixNano(args ...gst.Object) (
	ret gst.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gst.Int{Value: t1.UnixNano()}

	return
}

func timesTimeFormat(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 2 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := args[1].ToString()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	s := t1.Format(s2)
	if len(s) > gst.MaxStringLen {

		return nil, gse.ErrStringLimit
	}

	ret = &gst.String{Value: s}

	return
}

func timesIsZero(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	if t1.IsZero() {
		ret = gst.TrueValue
	} else {
		ret = gst.FalseValue
	}

	return
}

func timesToLocal(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gst.Time{Value: t1.Local()}

	return
}

func timesToUTC(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gst.Time{Value: t1.UTC()}

	return
}

func timesTimeLocation(args ...gst.Object) (
	ret gst.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gst.String{Value: t1.Location().String()}

	return
}

func timesInLocation(args ...gst.Object) (
	ret gst.Object,
	err error,
) {
	if len(args) != 2 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := args[1].ToString()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	location, err := time.LoadLocation(s2)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &gst.Time{Value: t1.In(location)}

	return
}

func timesTimeString(args ...gst.Object) (ret gst.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := args[0].ToTime()
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gst.String{Value: t1.String()}

	return
}
