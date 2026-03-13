package stdlib

import (
	"regexp"

	gse "github.com/jokruger/gs/error"
	gst "github.com/jokruger/gs/types"
)

func makeTextRegexp(re *regexp.Regexp) *gst.ImmutableMap {
	return &gst.ImmutableMap{
		Value: map[string]gst.Object{
			// match(text) => bool
			"match": &gst.UserFunction{
				Value: func(args ...gst.Object) (
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

					if re.MatchString(s1) {
						ret = gst.TrueValue
					} else {
						ret = gst.FalseValue
					}

					return
				},
			},

			// find(text) 			=> array(array({text:,begin:,end:}))/undefined
			// find(text, maxCount) => array(array({text:,begin:,end:}))/undefined
			"find": &gst.UserFunction{
				Value: func(args ...gst.Object) (
					ret gst.Object,
					err error,
				) {
					numArgs := len(args)
					if numArgs != 1 && numArgs != 2 {
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

					if numArgs == 1 {
						m := re.FindStringSubmatchIndex(s1)
						if m == nil {
							ret = gst.UndefinedValue
							return
						}

						arr := &gst.Array{}
						for i := 0; i < len(m); i += 2 {
							arr.Value = append(arr.Value,
								&gst.ImmutableMap{
									Value: map[string]gst.Object{
										"text": &gst.String{
											Value: s1[m[i]:m[i+1]],
										},
										"begin": &gst.Int{
											Value: int64(m[i]),
										},
										"end": &gst.Int{
											Value: int64(m[i+1]),
										},
									}})
						}

						ret = &gst.Array{Value: []gst.Object{arr}}

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
					m := re.FindAllStringSubmatchIndex(s1, i2)
					if m == nil {
						ret = gst.UndefinedValue
						return
					}

					arr := &gst.Array{}
					for _, m := range m {
						subMatch := &gst.Array{}
						for i := 0; i < len(m); i += 2 {
							subMatch.Value = append(subMatch.Value,
								&gst.ImmutableMap{
									Value: map[string]gst.Object{
										"text": &gst.String{
											Value: s1[m[i]:m[i+1]],
										},
										"begin": &gst.Int{
											Value: int64(m[i]),
										},
										"end": &gst.Int{
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
			"replace": &gst.UserFunction{
				Value: func(args ...gst.Object) (
					ret gst.Object,
					err error,
				) {
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

					s, ok := doTextRegexpReplace(re, s1, s2)
					if !ok {
						return nil, gse.ErrStringLimit
					}

					ret = &gst.String{Value: s}

					return
				},
			},

			// split(text) 			 => array(string)
			// split(text, maxCount) => array(string)
			"split": &gst.UserFunction{
				Value: func(args ...gst.Object) (
					ret gst.Object,
					err error,
				) {
					numArgs := len(args)
					if numArgs != 1 && numArgs != 2 {
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

					var i2 = -1
					if numArgs > 1 {
						i2, ok = args[1].ToInt()
						if !ok {
							err = gse.ErrInvalidArgumentType{
								Name:     "second",
								Expected: "int(compatible)",
								Found:    args[1].TypeName(),
							}
							return
						}
					}

					arr := &gst.Array{}
					for _, s := range re.Split(s1, i2) {
						arr.Value = append(arr.Value,
							&gst.String{Value: s})
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
		if len(out)+m[0]-idx+len(exp) > gst.MaxStringLen {
			return "", false
		}
		out += src[idx:m[0]] + string(exp)
		idx = m[1]
	}
	if idx < len(src) {
		if len(out)+len(src)-idx > gst.MaxStringLen {
			return "", false
		}
		out += src[idx:]
	}
	return out, true
}
