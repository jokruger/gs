package gs_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/jokruger/gs"
	gse "github.com/jokruger/gs/error"
	gst "github.com/jokruger/gs/types"
)

func Test_builtinDelete(t *testing.T) {
	var builtinDelete func(args ...gst.Object) (gst.Object, error)
	for _, f := range gs.GetAllBuiltinFunctions() {
		if f.Name == "delete" {
			builtinDelete = f.Value
			break
		}
	}
	if builtinDelete == nil {
		t.Fatal("builtin delete not found")
	}
	type args struct {
		args []gst.Object
	}
	tests := []struct {
		name      string
		args      args
		want      gst.Object
		wantErr   bool
		wantedErr error
		target    interface{}
	}{
		{name: "invalid-arg", args: args{[]gst.Object{&gst.String{},
			&gst.String{}}}, wantErr: true,
			wantedErr: gse.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "map",
				Found:    "string"},
		},
		{name: "no-args",
			wantErr: true, wantedErr: gse.ErrWrongNumArguments},
		{name: "empty-args", args: args{[]gst.Object{}}, wantErr: true,
			wantedErr: gse.ErrWrongNumArguments,
		},
		{name: "3-args", args: args{[]gst.Object{
			(*gst.Map)(nil), (*gst.String)(nil), (*gst.String)(nil)}},
			wantErr: true, wantedErr: gse.ErrWrongNumArguments,
		},
		{name: "nil-map-empty-key",
			args: args{[]gst.Object{&gst.Map{}, &gst.String{}}},
			want: gst.UndefinedValue,
		},
		{name: "nil-map-nonstr-key",
			args: args{[]gst.Object{
				&gst.Map{}, &gst.Int{}}}, wantErr: true,
			wantedErr: gse.ErrInvalidArgumentType{
				Name: "second", Expected: "string", Found: "int"},
		},
		{name: "nil-map-no-key",
			args: args{[]gst.Object{&gst.Map{}}}, wantErr: true,
			wantedErr: gse.ErrWrongNumArguments,
		},
		{name: "map-missing-key",
			args: args{
				[]gst.Object{
					&gst.Map{Value: map[string]gst.Object{
						"key": &gst.String{Value: "value"},
					}},
					&gst.String{Value: "key1"}}},
			want: gst.UndefinedValue,
			target: &gst.Map{
				Value: map[string]gst.Object{
					"key": &gst.String{
						Value: "value"}}},
		},
		{name: "map-emptied",
			args: args{
				[]gst.Object{
					&gst.Map{Value: map[string]gst.Object{
						"key": &gst.String{Value: "value"},
					}},
					&gst.String{Value: "key"}}},
			want:   gst.UndefinedValue,
			target: &gst.Map{Value: map[string]gst.Object{}},
		},
		{name: "map-multi-keys",
			args: args{
				[]gst.Object{
					&gst.Map{Value: map[string]gst.Object{
						"key1": &gst.String{Value: "value1"},
						"key2": &gst.Int{Value: 10},
					}},
					&gst.String{Value: "key1"}}},
			want: gst.UndefinedValue,
			target: &gst.Map{Value: map[string]gst.Object{
				"key2": &gst.Int{Value: 10}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := builtinDelete(tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("builtinDelete() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
			if tt.wantErr && !errors.Is(err, tt.wantedErr) {
				if err.Error() != tt.wantedErr.Error() {
					t.Errorf("builtinDelete() error = %v, wantedErr %v",
						err, tt.wantedErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("builtinDelete() = %v, want %v", got, tt.want)
				return
			}
			if !tt.wantErr && tt.target != nil {
				switch v := tt.args.args[0].(type) {
				case *gst.Map, *gst.Array:
					if !reflect.DeepEqual(tt.target, tt.args.args[0]) {
						t.Errorf("builtinDelete() objects are not equal "+
							"got: %+v, want: %+v", tt.args.args[0], tt.target)
					}
				default:
					t.Errorf("builtinDelete() unsuporrted arg[0] type %s",
						v.TypeName())
					return
				}
			}
		})
	}
}

func Test_builtinSplice(t *testing.T) {
	var builtinSplice func(args ...gst.Object) (gst.Object, error)
	for _, f := range gs.GetAllBuiltinFunctions() {
		if f.Name == "splice" {
			builtinSplice = f.Value
			break
		}
	}
	if builtinSplice == nil {
		t.Fatal("builtin splice not found")
	}
	tests := []struct {
		name      string
		args      []gst.Object
		deleted   gst.Object
		Array     *gst.Array
		wantErr   bool
		wantedErr error
	}{
		{name: "no args", args: []gst.Object{}, wantErr: true,
			wantedErr: gse.ErrWrongNumArguments,
		},
		{name: "invalid args", args: []gst.Object{&gst.Map{}},
			wantErr: true,
			wantedErr: gse.ErrInvalidArgumentType{
				Name: "first", Expected: "array", Found: "map"},
		},
		{name: "invalid args",
			args:    []gst.Object{&gst.Array{}, &gst.String{}},
			wantErr: true,
			wantedErr: gse.ErrInvalidArgumentType{
				Name: "second", Expected: "int", Found: "string"},
		},
		{name: "negative index",
			args:      []gst.Object{&gst.Array{}, &gst.Int{Value: -1}},
			wantErr:   true,
			wantedErr: gse.ErrIndexOutOfBounds},
		{name: "non int count",
			args: []gst.Object{
				&gst.Array{}, &gst.Int{Value: 0},
				&gst.String{Value: ""}},
			wantErr: true,
			wantedErr: gse.ErrInvalidArgumentType{
				Name: "third", Expected: "int", Found: "string"},
		},
		{name: "negative count",
			args: []gst.Object{
				&gst.Array{Value: []gst.Object{
					&gst.Int{Value: 0},
					&gst.Int{Value: 1},
					&gst.Int{Value: 2}}},
				&gst.Int{Value: 0},
				&gst.Int{Value: -1}},
			wantErr:   true,
			wantedErr: gse.ErrIndexOutOfBounds,
		},
		{name: "insert with zero count",
			args: []gst.Object{
				&gst.Array{Value: []gst.Object{
					&gst.Int{Value: 0},
					&gst.Int{Value: 1},
					&gst.Int{Value: 2}}},
				&gst.Int{Value: 0},
				&gst.Int{Value: 0},
				&gst.String{Value: "b"}},
			deleted: &gst.Array{Value: []gst.Object{}},
			Array: &gst.Array{Value: []gst.Object{
				&gst.String{Value: "b"},
				&gst.Int{Value: 0},
				&gst.Int{Value: 1},
				&gst.Int{Value: 2}}},
		},
		{name: "insert",
			args: []gst.Object{
				&gst.Array{Value: []gst.Object{
					&gst.Int{Value: 0},
					&gst.Int{Value: 1},
					&gst.Int{Value: 2}}},
				&gst.Int{Value: 1},
				&gst.Int{Value: 0},
				&gst.String{Value: "c"},
				&gst.String{Value: "d"}},
			deleted: &gst.Array{Value: []gst.Object{}},
			Array: &gst.Array{Value: []gst.Object{
				&gst.Int{Value: 0},
				&gst.String{Value: "c"},
				&gst.String{Value: "d"},
				&gst.Int{Value: 1},
				&gst.Int{Value: 2}}},
		},
		{name: "insert with zero count",
			args: []gst.Object{
				&gst.Array{Value: []gst.Object{
					&gst.Int{Value: 0},
					&gst.Int{Value: 1},
					&gst.Int{Value: 2}}},
				&gst.Int{Value: 1},
				&gst.Int{Value: 0},
				&gst.String{Value: "c"},
				&gst.String{Value: "d"}},
			deleted: &gst.Array{Value: []gst.Object{}},
			Array: &gst.Array{Value: []gst.Object{
				&gst.Int{Value: 0},
				&gst.String{Value: "c"},
				&gst.String{Value: "d"},
				&gst.Int{Value: 1},
				&gst.Int{Value: 2}}},
		},
		{name: "insert with delete",
			args: []gst.Object{
				&gst.Array{Value: []gst.Object{
					&gst.Int{Value: 0},
					&gst.Int{Value: 1},
					&gst.Int{Value: 2}}},
				&gst.Int{Value: 1},
				&gst.Int{Value: 1},
				&gst.String{Value: "c"},
				&gst.String{Value: "d"}},
			deleted: &gst.Array{
				Value: []gst.Object{&gst.Int{Value: 1}}},
			Array: &gst.Array{Value: []gst.Object{
				&gst.Int{Value: 0},
				&gst.String{Value: "c"},
				&gst.String{Value: "d"},
				&gst.Int{Value: 2}}},
		},
		{name: "insert with delete multi",
			args: []gst.Object{
				&gst.Array{Value: []gst.Object{
					&gst.Int{Value: 0},
					&gst.Int{Value: 1},
					&gst.Int{Value: 2}}},
				&gst.Int{Value: 1},
				&gst.Int{Value: 2},
				&gst.String{Value: "c"},
				&gst.String{Value: "d"}},
			deleted: &gst.Array{Value: []gst.Object{
				&gst.Int{Value: 1},
				&gst.Int{Value: 2}}},
			Array: &gst.Array{
				Value: []gst.Object{
					&gst.Int{Value: 0},
					&gst.String{Value: "c"},
					&gst.String{Value: "d"}}},
		},
		{name: "delete all with positive count",
			args: []gst.Object{
				&gst.Array{Value: []gst.Object{
					&gst.Int{Value: 0},
					&gst.Int{Value: 1},
					&gst.Int{Value: 2}}},
				&gst.Int{Value: 0},
				&gst.Int{Value: 3}},
			deleted: &gst.Array{Value: []gst.Object{
				&gst.Int{Value: 0},
				&gst.Int{Value: 1},
				&gst.Int{Value: 2}}},
			Array: &gst.Array{Value: []gst.Object{}},
		},
		{name: "delete all with big count",
			args: []gst.Object{
				&gst.Array{Value: []gst.Object{
					&gst.Int{Value: 0},
					&gst.Int{Value: 1},
					&gst.Int{Value: 2}}},
				&gst.Int{Value: 0},
				&gst.Int{Value: 5}},
			deleted: &gst.Array{Value: []gst.Object{
				&gst.Int{Value: 0},
				&gst.Int{Value: 1},
				&gst.Int{Value: 2}}},
			Array: &gst.Array{Value: []gst.Object{}},
		},
		{name: "nothing2",
			args: []gst.Object{
				&gst.Array{Value: []gst.Object{
					&gst.Int{Value: 0},
					&gst.Int{Value: 1},
					&gst.Int{Value: 2}}}},
			Array: &gst.Array{Value: []gst.Object{}},
			deleted: &gst.Array{Value: []gst.Object{
				&gst.Int{Value: 0},
				&gst.Int{Value: 1},
				&gst.Int{Value: 2}}},
		},
		{name: "pop without count",
			args: []gst.Object{
				&gst.Array{Value: []gst.Object{
					&gst.Int{Value: 0},
					&gst.Int{Value: 1},
					&gst.Int{Value: 2}}},
				&gst.Int{Value: 2}},
			deleted: &gst.Array{Value: []gst.Object{&gst.Int{Value: 2}}},
			Array: &gst.Array{Value: []gst.Object{
				&gst.Int{Value: 0}, &gst.Int{Value: 1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := builtinSplice(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("builtinSplice() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.deleted) {
				t.Errorf("builtinSplice() = %v, want %v", got, tt.deleted)
			}
			if tt.wantErr && tt.wantedErr.Error() != err.Error() {
				t.Errorf("builtinSplice() error = %v, wantedErr %v",
					err, tt.wantedErr)
			}
			if tt.Array != nil && !reflect.DeepEqual(tt.Array, tt.args[0]) {
				t.Errorf("builtinSplice() arrays are not equal expected"+
					" %s, got %s", tt.Array, tt.args[0].(*gst.Array))
			}
		})
	}
}

func Test_builtinRange(t *testing.T) {
	var builtinRange func(args ...gst.Object) (gst.Object, error)
	for _, f := range gs.GetAllBuiltinFunctions() {
		if f.Name == "range" {
			builtinRange = f.Value
			break
		}
	}
	if builtinRange == nil {
		t.Fatal("builtin range not found")
	}
	tests := []struct {
		name      string
		args      []gst.Object
		result    *gst.Array
		wantErr   bool
		wantedErr error
	}{
		{name: "no args", args: []gst.Object{}, wantErr: true,
			wantedErr: gse.ErrWrongNumArguments,
		},
		{name: "single args", args: []gst.Object{&gst.Map{}},
			wantErr:   true,
			wantedErr: gse.ErrWrongNumArguments,
		},
		{name: "4 args", args: []gst.Object{&gst.Map{}, &gst.String{}, &gst.String{}, &gst.String{}},
			wantErr:   true,
			wantedErr: gse.ErrWrongNumArguments,
		},
		{name: "invalid start",
			args:    []gst.Object{&gst.String{}, &gst.String{}},
			wantErr: true,
			wantedErr: gse.ErrInvalidArgumentType{
				Name: "start", Expected: "int", Found: "string"},
		},
		{name: "invalid stop",
			args:    []gst.Object{&gst.Int{}, &gst.String{}},
			wantErr: true,
			wantedErr: gse.ErrInvalidArgumentType{
				Name: "stop", Expected: "int", Found: "string"},
		},
		{name: "invalid step",
			args:    []gst.Object{&gst.Int{}, &gst.Int{}, &gst.String{}},
			wantErr: true,
			wantedErr: gse.ErrInvalidArgumentType{
				Name: "step", Expected: "int", Found: "string"},
		},
		{name: "zero step",
			args:      []gst.Object{&gst.Int{}, &gst.Int{}, &gst.Int{}}, //must greate than 0
			wantErr:   true,
			wantedErr: gse.ErrInvalidRangeStep,
		},
		{name: "negative step",
			args:      []gst.Object{&gst.Int{}, &gst.Int{}, intObject(-2)}, //must greate than 0
			wantErr:   true,
			wantedErr: gse.ErrInvalidRangeStep,
		},
		{name: "same bound",
			args:    []gst.Object{&gst.Int{}, &gst.Int{}},
			wantErr: false,
			result: &gst.Array{
				Value: nil,
			},
		},
		{name: "positive range",
			args:    []gst.Object{&gst.Int{}, &gst.Int{Value: 5}},
			wantErr: false,
			result: &gst.Array{
				Value: []gst.Object{
					intObject(0),
					intObject(1),
					intObject(2),
					intObject(3),
					intObject(4),
				},
			},
		},
		{name: "negative range",
			args:    []gst.Object{&gst.Int{}, &gst.Int{Value: -5}},
			wantErr: false,
			result: &gst.Array{
				Value: []gst.Object{
					intObject(0),
					intObject(-1),
					intObject(-2),
					intObject(-3),
					intObject(-4),
				},
			},
		},

		{name: "positive with step",
			args:    []gst.Object{&gst.Int{}, &gst.Int{Value: 5}, &gst.Int{Value: 2}},
			wantErr: false,
			result: &gst.Array{
				Value: []gst.Object{
					intObject(0),
					intObject(2),
					intObject(4),
				},
			},
		},

		{name: "negative with step",
			args:    []gst.Object{&gst.Int{}, &gst.Int{Value: -10}, &gst.Int{Value: 2}},
			wantErr: false,
			result: &gst.Array{
				Value: []gst.Object{
					intObject(0),
					intObject(-2),
					intObject(-4),
					intObject(-6),
					intObject(-8),
				},
			},
		},

		{name: "large range",
			args:    []gst.Object{intObject(-10), intObject(10), &gst.Int{Value: 3}},
			wantErr: false,
			result: &gst.Array{
				Value: []gst.Object{
					intObject(-10),
					intObject(-7),
					intObject(-4),
					intObject(-1),
					intObject(2),
					intObject(5),
					intObject(8),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := builtinRange(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("builtinRange() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.wantedErr.Error() != err.Error() {
				t.Errorf("builtinRange() error = %v, wantedErr %v",
					err, tt.wantedErr)
			}
			if tt.result != nil && !reflect.DeepEqual(tt.result, got) {
				t.Errorf("builtinRange() arrays are not equal expected"+
					" %s, got %s", tt.result, got.(*gst.Array))
			}
		})
	}
}
