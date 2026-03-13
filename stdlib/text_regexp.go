package stdlib

import (
	"regexp"

	"github.com/jokruger/gs"
	gse "github.com/jokruger/gs/error"
)

func makeTextRegexp(re *regexp.Regexp) *gs.ImmutableMap {
	return &gs.ImmutableMap{
		Value: map[string]gs.Object{
			// match(text) => bool
			"match": &gs.UserFunction{
				Value: func(args ...gs.Object) (
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

					if re.MatchString(s1) {
						ret = gs.TrueValue
					} else {
						ret = gs.FalseValue
					}

					return
				},
			},

			// find(text) 			=> array(array({text:,begin:,end:}))/undefined
			// find(text, maxCount) => array(array({text:,begin:,end:}))/undefined
			"find": &gs.UserFunction{
				Value: func(args ...gs.Object) (
					ret gs.Object,
					err error,
				) {
					numArgs := len(args)
					if numArgs != 1 && numArgs != 2 {
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

					if numArgs == 1 {
						m := re.FindStringSubmatchIndex(s1)
						if m == nil {
							ret = gs.UndefinedValue
							return
						}

						arr := &gs.Array{}
						for i := 0; i < len(m); i += 2 {
							arr.Value = append(arr.Value,
								&gs.ImmutableMap{
									Value: map[string]gs.Object{
										"text": &gs.String{
											Value: s1[m[i]:m[i+1]],
										},
										"begin": &gs.Int{
											Value: int64(m[i]),
										},
										"end": &gs.Int{
											Value: int64(m[i+1]),
										},
									}})
						}

						ret = &gs.Array{Value: []gs.Object{arr}}

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
					m := re.FindAllStringSubmatchIndex(s1, i2)
					if m == nil {
						ret = gs.UndefinedValue
						return
					}

					arr := &gs.Array{}
					for _, m := range m {
						subMatch := &gs.Array{}
						for i := 0; i < len(m); i += 2 {
							subMatch.Value = append(subMatch.Value,
								&gs.ImmutableMap{
									Value: map[string]gs.Object{
										"text": &gs.String{
											Value: s1[m[i]:m[i+1]],
										},
										"begin": &gs.Int{
											Value: int64(m[i]),
										},
										"end": &gs.Int{
											Value: int64(m[i+1]),
										},
									}})
						}

						arr.Value = append(arr.Value, subMatch)
					}

					ret = arr

					return
				},
			},

			// replace(src, repl) => string
			"replace": &gs.UserFunction{
				Value: func(args ...gs.Object) (
					ret gs.Object,
					err error,
				) {
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

					s, ok := doTextRegexpReplace(re, s1, s2)
					if !ok {
						return nil, gse.ErrStringLimit
					}

					ret = &gs.String{Value: s}

					return
				},
			},

			// split(text) 			 => array(string)
			// split(text, maxCount) => array(string)
			"split": &gs.UserFunction{
				Value: func(args ...gs.Object) (
					ret gs.Object,
					err error,
				) {
					numArgs := len(args)
					if numArgs != 1 && numArgs != 2 {
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

					var i2 = -1
					if numArgs > 1 {
						i2, ok = gs.ToInt(args[1])
						if !ok {
							err = gse.ErrInvalidArgumentType{
								Name:     "second",
								Expected: "int(compatible)",
								Found:    args[1].TypeName(),
							}
							return
						}
					}

					arr := &gs.Array{}
					for _, s := range re.Split(s1, i2) {
						arr.Value = append(arr.Value,
							&gs.String{Value: s})
					}

					ret = arr

					return
				},
			},
		},
	}
}

// Size-limit checking implementation of regexp.ReplaceAllString.
func doTextRegexpReplace(re *regexp.Regexp, src, repl string) (string, bool) {
	idx := 0
	out := ""
	for _, m := range re.FindAllStringSubmatchIndex(src, -1) {
		var exp []byte
		exp = re.ExpandString(exp, repl, src, m)
		if len(out)+m[0]-idx+len(exp) > gs.MaxStringLen {
			return "", false
		}
		out += src[idx:m[0]] + string(exp)
		idx = m[1]
	}
	if idx < len(src) {
		if len(out)+len(src)-idx > gs.MaxStringLen {
			return "", false
		}
		out += src[idx:]
	}
	return out, true
}
