package query_executor

import (
	"github.com/myprivatealaska/bradfield-database-systems/common"
)

type projection struct {
	Columns []string
	Child   common.Node
}

func NewProjection(columns []string, child common.Node) *projection {
	return &projection{
		Columns: columns,
		Child:   child,
	}
}

func (p *projection) Next() *common.Tuple {
	t := p.Child.Next()
	if t == nil {
		return t
	}
	proj := make(map[string]interface{})
	for _, key := range p.Columns {
		proj[key] = (*t)[key]
	}

	return (*common.Tuple)(&proj)
}
