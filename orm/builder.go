package orm

import (
	"JoeyOrm/orm/internal"
	reflect2 "JoeyOrm/orm/reflect"
	"strings"
)

type builder struct {
	sb    *strings.Builder
	args  []any
	where []Predicate
}

func (b *builder) BuilderPredicates(prefix string, model *reflect2.Model) error {
	if len(b.where) > 0 {
		b.sb.WriteString(" " + prefix + " ")

		err := b.BuilderPredicate(b.where, model)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *builder) BuilderPredicate(ps []Predicate, model *reflect2.Model) error {
	p := ps[0]
	for i := 1; i < len(ps); i++ {
		p = p.And(ps[i])
	}

	return b.BuilderExpression(p, model)
}
func (b *builder) BuilderExpression(e Expression, model *reflect2.Model) error {
	if e == nil {
		return nil
	}
	sb := b.sb

	switch exp := e.(type) {
	case Column:
		field, ok := model.FieldMap[exp.name]
		if !ok {
			return internal.NewNotFoundField(exp.name)
		}
		sb.WriteByte('`')
		sb.WriteString(field.ColName)
		sb.WriteByte('`')
	case value:
		sb.WriteString("?")
		b.addArg(exp.val)
	case Predicate:
		_, ok := exp.left.(Predicate)
		if ok {
			sb.WriteString("(")
		}
		if err := b.BuilderExpression(exp.left, model); err != nil {
			return err
		}
		if ok {
			sb.WriteString(")")
		}

		sb.WriteString(" ")
		sb.WriteString(exp.op.String())
		sb.WriteString(" ")

		_, ok = exp.right.(Predicate)
		if ok {
			sb.WriteString("(")
		}
		if err := b.BuilderExpression(exp.right, model); err != nil {
			return err
		}
		if ok {
			sb.WriteString(")")
		}

	default:
		return internal.NewUnSupportedExpression(exp)
	}

	return nil
}

func (b *builder) addArg(val any) {
	if b.args == nil {
		b.args = make([]any, 0, 8)
	}

	b.args = append(b.args, val)
}
