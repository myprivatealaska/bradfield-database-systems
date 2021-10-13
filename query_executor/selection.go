package query_executor

import (
	"github.com/myprivatealaska/bradfield-database-systems/common"
)

type selection struct {
	predicateFunc func(t common.Tuple) bool
	child         common.Node
}

func NewSelection(predicateFunc func(t common.Tuple) bool, child common.Node) *selection {
	return &selection{
		predicateFunc: predicateFunc,
		child:         child,
	}
}

func (s *selection) Next() *common.Tuple {
	for t := s.child.Next(); t != nil; t = s.child.Next() {
		result := s.predicateFunc(*t)
		if result == true {
			return t
		}
	}
	return nil
}
