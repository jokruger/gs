package gs_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/jokruger/gs"
)

func Test_builtinDelete(t *testing.T) {
	var builtinDelete func(args ...gs.Object) (gs.Object, error)
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
		args []gs.Object
	}
	tests := []struct {
		name      string
		args      args
		want      gs.Object
		wantErr   bool
		wantedErr error
		target    interface{}
	}{
		{name: "invalid-arg", args: args{[]gs.Object{&gs.String{},
			&gs.String{}}}, wantErr: true,
			wantedErr: gs.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "map",
				Found:    "string"},
		},
		{name: "no-args",
			wantErr: true, wantedErr: gs.ErrWrongNumArguments},
		{name: "empty-args", args: args{[]gs.Object{}}, wantErr: true,
			wantedErr: gs.ErrWrongNumArguments,
		},
		{name: "3-args", args: args{[]gs.Object{
			(*gs.Map)(nil), (*gs.String)(nil), (*gs.String)(nil)}},
			wantErr: true, wantedErr: gs.ErrWrongNumArguments,
		},
		{name: "nil-map-empty-key",
			args: args{[]gs.Object{&gs.Map{}, &gs.String{}}},
			want: gs.UndefinedValue,
		},
		{name: "nil-map-nonstr-key",
			args: args{[]gs.Object{
				&gs.Map{}, &gs.Int{}}}, wantErr: true,
			wantedErr: gs.ErrInvalidArgumentType{
				Name: "second", Expected: "string", Found: "int"},
		},
		{name: "nil-map-no-key",
			args: args{[]gs.Object{&gs.Map{}}}, wantErr: true,
			wantedErr: gs.ErrWrongNumArguments,
		},
		{name: "map-missing-key",
			args: args{
				[]gs.Object{
					&gs.Map{Value: map[string]gs.Object{
						"key": &gs.String{Value: "value"},
					}},
					&gs.String{Value: "key1"}}},
			want: gs.UndefinedValue,
			target: &gs.Map{
				Value: map[string]gs.Object{
					"key": &gs.String{
						Value: "value"}}},
		},
		{name: "map-emptied",
			args: args{
				[]gs.Object{
					&gs.Map{Value: map[string]gs.Object{
						"key": &gs.String{Value: "value"},
					}},
					&gs.String{Value: "key"}}},
			want:   gs.UndefinedValue,
			target: &gs.Map{Value: map[string]gs.Object{}},
		},
		{name: "map-multi-keys",
			args: args{
				[]gs.Object{
					&gs.Map{Value: map[string]gs.Object{
						"key1": &gs.String{Value: "value1"},
						"key2": &gs.Int{Value: 10},
					}},
					&gs.String{Value: "key1"}}},
			want: gs.UndefinedValue,
			target: &gs.Map{Value: map[string]gs.Object{
				"key2": &gs.Int{Value: 10}}},
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
				case *gs.Map, *gs.Array:
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
	var builtinSplice func(args ...gs.Object) (gs.Object, error)
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
		args      []gs.Object
		deleted   gs.Object
		Array     *gs.Array
		wantErr   bool
		wantedErr error
	}{
		{name: "no args", args: []gs.Object{}, wantErr: true,
			wantedErr: gs.ErrWrongNumArguments,
		},
		{name: "invalid args", args: []gs.Object{&gs.Map{}},
			wantErr: true,
			wantedErr: gs.ErrInvalidArgumentType{
				Name: "first", Expected: "array", Found: "map"},
		},
		{name: "invalid args",
			args:    []gs.Object{&gs.Array{}, &gs.String{}},
			wantErr: true,
			wantedErr: gs.ErrInvalidArgumentType{
				Name: "second", Expected: "int", Found: "string"},
		},
		{name: "negative index",
			args:      []gs.Object{&gs.Array{}, &gs.Int{Value: -1}},
			wantErr:   true,
			wantedErr: gs.ErrIndexOutOfBounds},
		{name: "non int count",
			args: []gs.Object{
				&gs.Array{}, &gs.Int{Value: 0},
				&gs.String{Value: ""}},
			wantErr: true,
			wantedErr: gs.ErrInvalidArgumentType{
				Name: "third", Expected: "int", Found: "string"},
		},
		{name: "negative count",
			args: []gs.Object{
				&gs.Array{Value: []gs.Object{
					&gs.Int{Value: 0},
					&gs.Int{Value: 1},
					&gs.Int{Value: 2}}},
				&gs.Int{Value: 0},
				&gs.Int{Value: -1}},
			wantErr:   true,
			wantedErr: gs.ErrIndexOutOfBounds,
		},
		{name: "insert with zero count",
			args: []gs.Object{
				&gs.Array{Value: []gs.Object{
					&gs.Int{Value: 0},
					&gs.Int{Value: 1},
					&gs.Int{Value: 2}}},
				&gs.Int{Value: 0},
				&gs.Int{Value: 0},
				&gs.String{Value: "b"}},
			deleted: &gs.Array{Value: []gs.Object{}},
			Array: &gs.Array{Value: []gs.Object{
				&gs.String{Value: "b"},
				&gs.Int{Value: 0},
				&gs.Int{Value: 1},
				&gs.Int{Value: 2}}},
		},
		{name: "insert",
			args: []gs.Object{
				&gs.Array{Value: []gs.Object{
					&gs.Int{Value: 0},
					&gs.Int{Value: 1},
					&gs.Int{Value: 2}}},
				&gs.Int{Value: 1},
				&gs.Int{Value: 0},
				&gs.String{Value: "c"},
				&gs.String{Value: "d"}},
			deleted: &gs.Array{Value: []gs.Object{}},
			Array: &gs.Array{Value: []gs.Object{
				&gs.Int{Value: 0},
				&gs.String{Value: "c"},
				&gs.String{Value: "d"},
				&gs.Int{Value: 1},
				&gs.Int{Value: 2}}},
		},
		{name: "insert with zero count",
			args: []gs.Object{
				&gs.Array{Value: []gs.Object{
					&gs.Int{Value: 0},
					&gs.Int{Value: 1},
					&gs.Int{Value: 2}}},
				&gs.Int{Value: 1},
				&gs.Int{Value: 0},
				&gs.String{Value: "c"},
				&gs.String{Value: "d"}},
			deleted: &gs.Array{Value: []gs.Object{}},
			Array: &gs.Array{Value: []gs.Object{
				&gs.Int{Value: 0},
				&gs.String{Value: "c"},
				&gs.String{Value: "d"},
				&gs.Int{Value: 1},
				&gs.Int{Value: 2}}},
		},
		{name: "insert with delete",
			args: []gs.Object{
				&gs.Array{Value: []gs.Object{
					&gs.Int{Value: 0},
					&gs.Int{Value: 1},
					&gs.Int{Value: 2}}},
				&gs.Int{Value: 1},
				&gs.Int{Value: 1},
				&gs.String{Value: "c"},
				&gs.String{Value: "d"}},
			deleted: &gs.Array{
				Value: []gs.Object{&gs.Int{Value: 1}}},
			Array: &gs.Array{Value: []gs.Object{
				&gs.Int{Value: 0},
				&gs.String{Value: "c"},
				&gs.String{Value: "d"},
				&gs.Int{Value: 2}}},
		},
		{name: "insert with delete multi",
			args: []gs.Object{
				&gs.Array{Value: []gs.Object{
					&gs.Int{Value: 0},
					&gs.Int{Value: 1},
					&gs.Int{Value: 2}}},
				&gs.Int{Value: 1},
				&gs.Int{Value: 2},
				&gs.String{Value: "c"},
				&gs.String{Value: "d"}},
			deleted: &gs.Array{Value: []gs.Object{
				&gs.Int{Value: 1},
				&gs.Int{Value: 2}}},
			Array: &gs.Array{
				Value: []gs.Object{
					&gs.Int{Value: 0},
					&gs.String{Value: "c"},
					&gs.String{Value: "d"}}},
		},
		{name: "delete all with positive count",
			args: []gs.Object{
				&gs.Array{Value: []gs.Object{
					&gs.Int{Value: 0},
					&gs.Int{Value: 1},
					&gs.Int{Value: 2}}},
				&gs.Int{Value: 0},
				&gs.Int{Value: 3}},
			deleted: &gs.Array{Value: []gs.Object{
				&gs.Int{Value: 0},
				&gs.Int{Value: 1},
				&gs.Int{Value: 2}}},
			Array: &gs.Array{Value: []gs.Object{}},
		},
		{name: "delete all with big count",
			args: []gs.Object{
				&gs.Array{Value: []gs.Object{
					&gs.Int{Value: 0},
					&gs.Int{Value: 1},
					&gs.Int{Value: 2}}},
				&gs.Int{Value: 0},
				&gs.Int{Value: 5}},
			deleted: &gs.Array{Value: []gs.Object{
				&gs.Int{Value: 0},
				&gs.Int{Value: 1},
				&gs.Int{Value: 2}}},
			Array: &gs.Array{Value: []gs.Object{}},
		},
		{name: "nothing2",
			args: []gs.Object{
				&gs.Array{Value: []gs.Object{
					&gs.Int{Value: 0},
					&gs.Int{Value: 1},
					&gs.Int{Value: 2}}}},
			Array: &gs.Array{Value: []gs.Object{}},
			deleted: &gs.Array{Value: []gs.Object{
				&gs.Int{Value: 0},
				&gs.Int{Value: 1},
				&gs.Int{Value: 2}}},
		},
		{name: "pop without count",
			args: []gs.Object{
				&gs.Array{Value: []gs.Object{
					&gs.Int{Value: 0},
					&gs.Int{Value: 1},
					&gs.Int{Value: 2}}},
				&gs.Int{Value: 2}},
			deleted: &gs.Array{Value: []gs.Object{&gs.Int{Value: 2}}},
			Array: &gs.Array{Value: []gs.Object{
				&gs.Int{Value: 0}, &gs.Int{Value: 1}}},
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
					" %s, got %s", tt.Array, tt.args[0].(*gs.Array))
			}
		})
	}
}

func Test_builtinRange(t *testing.T) {
	var builtinRange func(args ...gs.Object) (gs.Object, error)
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
		args      []gs.Object
		result    *gs.Array
		wantErr   bool
		wantedErr error
	}{
		{name: "no args", args: []gs.Object{}, wantErr: true,
			wantedErr: gs.ErrWrongNumArguments,
		},
		{name: "single args", args: []gs.Object{&gs.Map{}},
			wantErr:   true,
			wantedErr: gs.ErrWrongNumArguments,
		},
		{name: "4 args", args: []gs.Object{&gs.Map{}, &gs.String{}, &gs.String{}, &gs.String{}},
			wantErr:   true,
			wantedErr: gs.ErrWrongNumArguments,
		},
		{name: "invalid start",
			args:    []gs.Object{&gs.String{}, &gs.String{}},
			wantErr: true,
			wantedErr: gs.ErrInvalidArgumentType{
				Name: "start", Expected: "int", Found: "string"},
		},
		{name: "invalid stop",
			args:    []gs.Object{&gs.Int{}, &gs.String{}},
			wantErr: true,
			wantedErr: gs.ErrInvalidArgumentType{
				Name: "stop", Expected: "int", Found: "string"},
		},
		{name: "invalid step",
			args:    []gs.Object{&gs.Int{}, &gs.Int{}, &gs.String{}},
			wantErr: true,
			wantedErr: gs.ErrInvalidArgumentType{
				Name: "step", Expected: "int", Found: "string"},
		},
		{name: "zero step",
			args:      []gs.Object{&gs.Int{}, &gs.Int{}, &gs.Int{}}, //must greate than 0
			wantErr:   true,
			wantedErr: gs.ErrInvalidRangeStep,
		},
		{name: "negative step",
			args:      []gs.Object{&gs.Int{}, &gs.Int{}, intObject(-2)}, //must greate than 0
			wantErr:   true,
			wantedErr: gs.ErrInvalidRangeStep,
		},
		{name: "same bound",
			args:    []gs.Object{&gs.Int{}, &gs.Int{}},
			wantErr: false,
			result: &gs.Array{
				Value: nil,
			},
		},
		{name: "positive range",
			args:    []gs.Object{&gs.Int{}, &gs.Int{Value: 5}},
			wantErr: false,
			result: &gs.Array{
				Value: []gs.Object{
					intObject(0),
					intObject(1),
					intObject(2),
					intObject(3),
					intObject(4),
				},
			},
		},
		{name: "negative range",
			args:    []gs.Object{&gs.Int{}, &gs.Int{Value: -5}},
			wantErr: false,
			result: &gs.Array{
				Value: []gs.Object{
					intObject(0),
					intObject(-1),
					intObject(-2),
					intObject(-3),
					intObject(-4),
				},
			},
		},

		{name: "positive with step",
			args:    []gs.Object{&gs.Int{}, &gs.Int{Value: 5}, &gs.Int{Value: 2}},
			wantErr: false,
			result: &gs.Array{
				Value: []gs.Object{
					intObject(0),
					intObject(2),
					intObject(4),
				},
			},
		},

		{name: "negative with step",
			args:    []gs.Object{&gs.Int{}, &gs.Int{Value: -10}, &gs.Int{Value: 2}},
			wantErr: false,
			result: &gs.Array{
				Value: []gs.Object{
					intObject(0),
					intObject(-2),
					intObject(-4),
					intObject(-6),
					intObject(-8),
				},
			},
		},

		{name: "large range",
			args:    []gs.Object{intObject(-10), intObject(10), &gs.Int{Value: 3}},
			wantErr: false,
			result: &gs.Array{
				Value: []gs.Object{
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
					" %s, got %s", tt.result, got.(*gs.Array))
			}
		})
	}
}
