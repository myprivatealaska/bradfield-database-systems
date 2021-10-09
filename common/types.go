package common

// SupportedAttributeValues null Int, null String, null Float
type SupportedAttributeTypes int

//const (
//	NullInt SupportedAttributeTypes = iota
//	NullString
//	NullFloat
//)

type Attribute struct {
	Name  string
	Value interface{}
}

type Tuple struct {
	Attributes []Attribute
}
