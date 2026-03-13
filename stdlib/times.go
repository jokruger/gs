package stdlib

import (
	"time"

	"github.com/jokruger/gs"
	gse "github.com/jokruger/gs/error"
)

var timesModule = map[string]gs.Object{
	"format_ansic":        &gs.String{Value: time.ANSIC},
	"format_unix_date":    &gs.String{Value: time.UnixDate},
	"format_ruby_date":    &gs.String{Value: time.RubyDate},
	"format_rfc822":       &gs.String{Value: time.RFC822},
	"format_rfc822z":      &gs.String{Value: time.RFC822Z},
	"format_rfc850":       &gs.String{Value: time.RFC850},
	"format_rfc1123":      &gs.String{Value: time.RFC1123},
	"format_rfc1123z":     &gs.String{Value: time.RFC1123Z},
	"format_rfc3339":      &gs.String{Value: time.RFC3339},
	"format_rfc3339_nano": &gs.String{Value: time.RFC3339Nano},
	"format_kitchen":      &gs.String{Value: time.Kitchen},
	"format_stamp":        &gs.String{Value: time.Stamp},
	"format_stamp_milli":  &gs.String{Value: time.StampMilli},
	"format_stamp_micro":  &gs.String{Value: time.StampMicro},
	"format_stamp_nano":   &gs.String{Value: time.StampNano},
	"nanosecond":          &gs.Int{Value: int64(time.Nanosecond)},
	"microsecond":         &gs.Int{Value: int64(time.Microsecond)},
	"millisecond":         &gs.Int{Value: int64(time.Millisecond)},
	"second":              &gs.Int{Value: int64(time.Second)},
	"minute":              &gs.Int{Value: int64(time.Minute)},
	"hour":                &gs.Int{Value: int64(time.Hour)},
	"january":             &gs.Int{Value: int64(time.January)},
	"february":            &gs.Int{Value: int64(time.February)},
	"march":               &gs.Int{Value: int64(time.March)},
	"april":               &gs.Int{Value: int64(time.April)},
	"may":                 &gs.Int{Value: int64(time.May)},
	"june":                &gs.Int{Value: int64(time.June)},
	"july":                &gs.Int{Value: int64(time.July)},
	"august":              &gs.Int{Value: int64(time.August)},
	"september":           &gs.Int{Value: int64(time.September)},
	"october":             &gs.Int{Value: int64(time.October)},
	"november":            &gs.Int{Value: int64(time.November)},
	"december":            &gs.Int{Value: int64(time.December)},
	"sleep": &gs.UserFunction{
		Name:  "sleep",
		Value: timesSleep,
	}, // sleep(int)
	"parse_duration": &gs.UserFunction{
		Name:  "parse_duration",
		Value: timesParseDuration,
	}, // parse_duration(str) => int
	"since": &gs.UserFunction{
		Name:  "since",
		Value: timesSince,
	}, // since(time) => int
	"until": &gs.UserFunction{
		Name:  "until",
		Value: timesUntil,
	}, // until(time) => int
	"duration_hours": &gs.UserFunction{
		Name:  "duration_hours",
		Value: timesDurationHours,
	}, // duration_hours(int) => float
	"duration_minutes": &gs.UserFunction{
		Name:  "duration_minutes",
		Value: timesDurationMinutes,
	}, // duration_minutes(int) => float
	"duration_nanoseconds": &gs.UserFunction{
		Name:  "duration_nanoseconds",
		Value: timesDurationNanoseconds,
	}, // duration_nanoseconds(int) => int
	"duration_seconds": &gs.UserFunction{
		Name:  "duration_seconds",
		Value: timesDurationSeconds,
	}, // duration_seconds(int) => float
	"duration_string": &gs.UserFunction{
		Name:  "duration_string",
		Value: timesDurationString,
	}, // duration_string(int) => string
	"month_string": &gs.UserFunction{
		Name:  "month_string",
		Value: timesMonthString,
	}, // month_string(int) => string
	"date": &gs.UserFunction{
		Name:  "date",
		Value: timesDate,
	}, // date(year, month, day, hour, min, sec, nsec) => time
	"now": &gs.UserFunction{
		Name:  "now",
		Value: timesNow,
	}, // now() => time
	"parse": &gs.UserFunction{
		Name:  "parse",
		Value: timesParse,
	}, // parse(format, str) => time
	"unix": &gs.UserFunction{
		Name:  "unix",
		Value: timesUnix,
	}, // unix(sec, nsec) => time
	"add": &gs.UserFunction{
		Name:  "add",
		Value: timesAdd,
	}, // add(time, int) => time
	"add_date": &gs.UserFunction{
		Name:  "add_date",
		Value: timesAddDate,
	}, // add_date(time, years, months, days) => time
	"sub": &gs.UserFunction{
		Name:  "sub",
		Value: timesSub,
	}, // sub(t time, u time) => int
	"after": &gs.UserFunction{
		Name:  "after",
		Value: timesAfter,
	}, // after(t time, u time) => bool
	"before": &gs.UserFunction{
		Name:  "before",
		Value: timesBefore,
	}, // before(t time, u time) => bool
	"time_year": &gs.UserFunction{
		Name:  "time_year",
		Value: timesTimeYear,
	}, // time_year(time) => int
	"time_month": &gs.UserFunction{
		Name:  "time_month",
		Value: timesTimeMonth,
	}, // time_month(time) => int
	"time_day": &gs.UserFunction{
		Name:  "time_day",
		Value: timesTimeDay,
	}, // time_day(time) => int
	"time_weekday": &gs.UserFunction{
		Name:  "time_weekday",
		Value: timesTimeWeekday,
	}, // time_weekday(time) => int
	"time_hour": &gs.UserFunction{
		Name:  "time_hour",
		Value: timesTimeHour,
	}, // time_hour(time) => int
	"time_minute": &gs.UserFunction{
		Name:  "time_minute",
		Value: timesTimeMinute,
	}, // time_minute(time) => int
	"time_second": &gs.UserFunction{
		Name:  "time_second",
		Value: timesTimeSecond,
	}, // time_second(time) => int
	"time_nanosecond": &gs.UserFunction{
		Name:  "time_nanosecond",
		Value: timesTimeNanosecond,
	}, // time_nanosecond(time) => int
	"time_unix": &gs.UserFunction{
		Name:  "time_unix",
		Value: timesTimeUnix,
	}, // time_unix(time) => int
	"time_unix_nano": &gs.UserFunction{
		Name:  "time_unix_nano",
		Value: timesTimeUnixNano,
	}, // time_unix_nano(time) => int
	"time_format": &gs.UserFunction{
		Name:  "time_format",
		Value: timesTimeFormat,
	}, // time_format(time, format) => string
	"time_location": &gs.UserFunction{
		Name:  "time_location",
		Value: timesTimeLocation,
	}, // time_location(time) => string
	"time_string": &gs.UserFunction{
		Name:  "time_string",
		Value: timesTimeString,
	}, // time_string(time) => string
	"is_zero": &gs.UserFunction{
		Name:  "is_zero",
		Value: timesIsZero,
	}, // is_zero(time) => bool
	"to_local": &gs.UserFunction{
		Name:  "to_local",
		Value: timesToLocal,
	}, // to_local(time) => time
	"to_utc": &gs.UserFunction{
		Name:  "to_utc",
		Value: timesToUTC,
	}, // to_utc(time) => time
	"in_location": &gs.UserFunction{
		Name:  "in_location",
		Value: timesInLocation,
	}, // in_location(time, location) => time
}

func timesSleep(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	i1, ok := gs.ToInt64(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	time.Sleep(time.Duration(i1))
	ret = gs.UndefinedValue

	return
}

func timesParseDuration(args ...gs.Object) (
	ret gs.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	s1, ok := gs.ToString(args[0])
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

	ret = &gs.Int{Value: int64(dur)}

	return
}

func timesSince(args ...gs.Object) (
	ret gs.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gs.Int{Value: int64(time.Since(t1))}

	return
}

func timesUntil(args ...gs.Object) (
	ret gs.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gs.Int{Value: int64(time.Until(t1))}

	return
}

func timesDurationHours(args ...gs.Object) (
	ret gs.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	i1, ok := gs.ToInt64(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gs.Float{Value: time.Duration(i1).Hours()}

	return
}

func timesDurationMinutes(args ...gs.Object) (
	ret gs.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	i1, ok := gs.ToInt64(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gs.Float{Value: time.Duration(i1).Minutes()}

	return
}

func timesDurationNanoseconds(args ...gs.Object) (
	ret gs.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	i1, ok := gs.ToInt64(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gs.Int{Value: time.Duration(i1).Nanoseconds()}

	return
}

func timesDurationSeconds(args ...gs.Object) (
	ret gs.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	i1, ok := gs.ToInt64(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gs.Float{Value: time.Duration(i1).Seconds()}

	return
}

func timesDurationString(args ...gs.Object) (
	ret gs.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	i1, ok := gs.ToInt64(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gs.String{Value: time.Duration(i1).String()}

	return
}

func timesMonthString(args ...gs.Object) (
	ret gs.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	i1, ok := gs.ToInt64(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gs.String{Value: time.Month(i1).String()}

	return
}

func timesDate(args ...gs.Object) (
	ret gs.Object,
	err error,
) {
	if len(args) < 7 || len(args) > 8 {
		err = gse.ErrWrongNumArguments
		return
	}

	i1, ok := gs.ToInt(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}
	i2, ok := gs.ToInt(args[1])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}
	i3, ok := gs.ToInt(args[2])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}
	i4, ok := gs.ToInt(args[3])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}
	i5, ok := gs.ToInt(args[4])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "fifth",
			Expected: "int(compatible)",
			Found:    args[4].TypeName(),
		}
		return
	}
	i6, ok := gs.ToInt(args[5])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "sixth",
			Expected: "int(compatible)",
			Found:    args[5].TypeName(),
		}
		return
	}
	i7, ok := gs.ToInt(args[6])
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
		i8, ok := gs.ToString(args[7])
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

	ret = &gs.Time{
		Value: time.Date(i1,
			time.Month(i2), i3, i4, i5, i6, i7, loc),
	}

	return
}

func timesNow(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 0 {
		err = gse.ErrWrongNumArguments
		return
	}

	ret = &gs.Time{Value: time.Now()}

	return
}

func timesParse(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 2 {
		err = gse.ErrWrongNumArguments
		return
	}

	s1, ok := gs.ToString(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := gs.ToString(args[1])
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

	ret = &gs.Time{Value: parsed}

	return
}

func timesUnix(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 2 {
		err = gse.ErrWrongNumArguments
		return
	}

	i1, ok := gs.ToInt64(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := gs.ToInt64(args[1])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &gs.Time{Value: time.Unix(i1, i2)}

	return
}

func timesAdd(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 2 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := gs.ToInt64(args[1])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &gs.Time{Value: t1.Add(time.Duration(i2))}

	return
}

func timesSub(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 2 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := gs.ToTime(args[1])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &gs.Int{Value: int64(t1.Sub(t2))}

	return
}

func timesAddDate(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 4 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := gs.ToInt(args[1])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	i3, ok := gs.ToInt(args[2])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	i4, ok := gs.ToInt(args[3])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}

	ret = &gs.Time{Value: t1.AddDate(i2, i3, i4)}

	return
}

func timesAfter(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 2 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := gs.ToTime(args[1])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	if t1.After(t2) {
		ret = gs.TrueValue
	} else {
		ret = gs.FalseValue
	}

	return
}

func timesBefore(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 2 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := gs.ToTime(args[1])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	if t1.Before(t2) {
		ret = gs.TrueValue
	} else {
		ret = gs.FalseValue
	}

	return
}

func timesTimeYear(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gs.Int{Value: int64(t1.Year())}

	return
}

func timesTimeMonth(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gs.Int{Value: int64(t1.Month())}

	return
}

func timesTimeDay(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gs.Int{Value: int64(t1.Day())}

	return
}

func timesTimeWeekday(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gs.Int{Value: int64(t1.Weekday())}

	return
}

func timesTimeHour(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gs.Int{Value: int64(t1.Hour())}

	return
}

func timesTimeMinute(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gs.Int{Value: int64(t1.Minute())}

	return
}

func timesTimeSecond(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gs.Int{Value: int64(t1.Second())}

	return
}

func timesTimeNanosecond(args ...gs.Object) (
	ret gs.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gs.Int{Value: int64(t1.Nanosecond())}

	return
}

func timesTimeUnix(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gs.Int{Value: t1.Unix()}

	return
}

func timesTimeUnixNano(args ...gs.Object) (
	ret gs.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gs.Int{Value: t1.UnixNano()}

	return
}

func timesTimeFormat(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 2 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := gs.ToString(args[1])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	s := t1.Format(s2)
	if len(s) > gs.MaxStringLen {

		return nil, gse.ErrStringLimit
	}

	ret = &gs.String{Value: s}

	return
}

func timesIsZero(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	if t1.IsZero() {
		ret = gs.TrueValue
	} else {
		ret = gs.FalseValue
	}

	return
}

func timesToLocal(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gs.Time{Value: t1.Local()}

	return
}

func timesToUTC(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gs.Time{Value: t1.UTC()}

	return
}

func timesTimeLocation(args ...gs.Object) (
	ret gs.Object,
	err error,
) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gs.String{Value: t1.Location().String()}

	return
}

func timesInLocation(args ...gs.Object) (
	ret gs.Object,
	err error,
) {
	if len(args) != 2 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := gs.ToString(args[1])
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

	ret = &gs.Time{Value: t1.In(location)}

	return
}

func timesTimeString(args ...gs.Object) (ret gs.Object, err error) {
	if len(args) != 1 {
		err = gse.ErrWrongNumArguments
		return
	}

	t1, ok := gs.ToTime(args[0])
	if !ok {
		err = gse.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &gs.String{Value: t1.String()}

	return
}
