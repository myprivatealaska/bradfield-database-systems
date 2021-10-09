package query_executor

// Implement Scan, Limit, Selection, Projection, Sort Nodes

type RawTuple map[string]string

type Node interface {
	Next()
}

type Operation string

const (
	EqualsOp  Operation = "EQUALS"
	GreaterOp Operation = "GREATER"
)

type Query struct {
	Projection []string
	Selection  struct {
		Attribute string
		Operation Operation
		Predicate interface{}
	}
	Relation string
}

type QueryExecutor struct {
}
