package query_executor

import (
	"github.com/myprivatealaska/bradfield-database-systems/common"
)

type selection struct {
	predicateFunc func(t common.Tuple) bool
	child         common.Node
}

func NewSelection() *selection {
	return &selection{
		predicateFunc: nil,
		child:         nil,
	}
}

func (s *selection) Next() *common.Tuple {
	return s.getFirstMatch()
}

func (s *selection) getFirstMatch() *common.Tuple {
	t := s.child.Next()
	if t == nil {
		return nil
	}
	result := s.predicateFunc(*t)
	if result == true {
		return t
	} else {
		return s.getFirstMatch()
	}
}
