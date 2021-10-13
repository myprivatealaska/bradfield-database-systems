package common

// Node - each node (Scan, Sort, Limit, etc.) will produce the next tuple in its output sequence or nil if no more tuples are available
type Node interface {
	Next() *Tuple
}

// SupportedAttributeValues null Int, null String, null Float
type SupportedAttributeTypes int

//const (
//	NullInt SupportedAttributeTypes = iota
//	NullString
//	NullFloat
//)

type Tuple map[string]interface{}

type Operation string

const (
	EqualsOp  Operation = "EQUALS"
	GreaterOp Operation = "GREATER"
)

type Query struct {
	Projection []string
	Selection  string
	Limit      int
	Relation   string
}
