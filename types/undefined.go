package types

type Undefined struct {
	ObjectImpl
}

func (o *Undefined) TypeName() string {
	return "undefined"
}

func (o *Undefined) String() string {
	return "<undefined>"
}

func (o *Undefined) Copy() Object {
	return o
}

func (o *Undefined) IsFalsy() bool {
	return true
}

func (o *Undefined) Equals(x Object) bool {
	return o == x
}

func (o *Undefined) IndexGet(Object) (Object, error) {
	return UndefinedValue, nil
}

func (o *Undefined) Iterate() Iterator {
	return o
}

func (o *Undefined) CanIterate() bool {
	return true
}

func (o *Undefined) Next() bool {
	return false
}

func (o *Undefined) Key() Object {
	return o
}

func (o *Undefined) Value() Object {
	return o
}

func (o *Undefined) ToBool() (bool, bool) {
	return false, true
}
