package query_executor

import (
	"github.com/myprivatealaska/bradfield-database-systems/common"
)

type limit struct {
	LimitValue int
	Count      int
	Child      common.Node
}

func NewLimit(limitVal int, child common.Node) *limit {
	return &limit{
		LimitValue: limitVal,
		Child:      child,
	}
}

func (l *limit) Next() *common.Tuple {
	if l.Count >= l.LimitValue {
		return nil
	}
	res := l.Child.Next()
	l.Count++
	return res
}
