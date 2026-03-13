package types

type ObjectPtr struct {
	ObjectImpl
	Value *Object
}

func (o *ObjectPtr) TypeName() string {
	return "<free-var>"
}

func (o *ObjectPtr) String() string {
	return "free-var"
}

func (o *ObjectPtr) Copy() Object {
	return o
}

func (o *ObjectPtr) IsFalsy() bool {
	return o.Value == nil
}

func (o *ObjectPtr) Equals(x Object) bool {
	return o == x
}

func (o *ObjectPtr) ToBool() (bool, bool) {
	return !o.IsFalsy(), true
}
