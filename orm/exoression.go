package orm

type op string

const (
	opEQ  = "="
	opLT  = "<"
	opGT  = ">"
	opAND = "AND"
	opOR  = "OR"
	opNOT = "NOT"
)

func (o op) String() string {
	return string(o)
}

type Predicate struct {
	left  Expression
	op    op
	right Expression
}

func (p Predicate) expr() {
	//TODO implement me
	panic("implement me")
}

type Column struct {
	name string
}

func C(name string) Column {
	return Column{name: name}
}

func (c Column) expr() {
	//TODO implement me
	panic("implement me")
}

func (c Column) Eq(arg any) Predicate {
	return Predicate{
		left:  c,
		op:    opEQ,
		right: value{val: arg},
	}
}

func Not(p Predicate) Predicate {
	return Predicate{
		op:    opNOT,
		right: p,
	}
}

func (p Predicate) And(r Predicate) Predicate {
	return Predicate{
		left:  p,
		op:    opAND,
		right: r,
	}
}

func (p Predicate) Or(r Predicate) Predicate {
	return Predicate{
		left:  p,
		op:    opOR,
		right: r,
	}
}

type value struct {
	val any
}

func (v value) expr() {
	//TODO implement me
	panic("implement me")
}

type Expression interface {
	expr()
}
