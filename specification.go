package serification

type Specification interface {
	IsSatisfied(interface{}) bool

	And(Specification) Specification
	Or(Specification) Specification
	Not() Specification
}

type AndSpecification struct {
	Left  Specification
	Right Specification
}

func (a AndSpecification) IsSatisfied(value interface{}) bool {
	return a.Left.IsSatisfied(value) && a.Right.IsSatisfied(value)
}

func (a AndSpecification) And(other Specification) Specification {
	return AndSpecification{Left: a, Right: other}
}

func (a AndSpecification) Or(other Specification) Specification {
	return OrSpecification{Left: a, Right: other}
}

func (a AndSpecification) Not() Specification {
	return NotSpecification{Subject: a}
}

type OrSpecification struct {
	Left  Specification
	Right Specification
}

func (o OrSpecification) IsSatisfied(value interface{}) bool {
	return o.Left.IsSatisfied(value) || o.Right.IsSatisfied(value)
}

func (o OrSpecification) And(other Specification) Specification {
	return AndSpecification{Left: o, Right: other}
}

func (o OrSpecification) Or(other Specification) Specification {
	return OrSpecification{Left: o, Right: other}
}

func (o OrSpecification) Not() Specification {
	return NotSpecification{Subject: o}
}

type NotSpecification struct {
	Subject Specification
}

func (n NotSpecification) IsSatisfied(value interface{}) bool {
	return !n.Subject.IsSatisfied(value)
}

func (n NotSpecification) And(other Specification) Specification {
	return AndSpecification{Left: n, Right: other}
}

func (n NotSpecification) Or(other Specification) Specification {
	return OrSpecification{Left: n, Right: other}
}

func (n NotSpecification) Not() Specification {
	return NotSpecification{Subject: n}
}
