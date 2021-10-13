package query_executor

import (
	"testing"

	"github.com/myprivatealaska/bradfield-database-systems/common"
	"go.uber.org/zap"
)

var data = []common.Tuple{
	map[string]interface{}{
		"id":     1,
		"name":   "Toy Story (1995)",
		"genres": "Adventure|Animation|Children|Comedy|Fantasy",
	},
	map[string]interface{}{
		"id":     2,
		"name":   "Jumanji (1995)",
		"genres": "Adventure|Children|Fantasy",
	},
	map[string]interface{}{
		"id":     3,
		"name":   "Grumpier Old Men (1995)",
		"genres": "Comedy|Romance",
	},
	map[string]interface{}{
		"id":     4,
		"name":   "Waiting to Exhale (1995)",
		"genres": "Comedy|Drama|Romance",
	},
}

func Execute(n common.Node) []common.Tuple {
	tuples := make([]common.Tuple, 0)
	for t := n.Next(); t != nil; t = n.Next() {
		tuples = append(tuples, *t)
	}
	return tuples
}

func TestLimitQuery(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	logger.Info("SELECT name FROM movies LIMIT 3")
	scanNode := NewScan(&data, 3)
	limitNode := NewLimit(3, scanNode)
	projectNode := NewProjection([]string{"name"}, limitNode)

	results := Execute(projectNode)
	logger.Info("Result", zap.Any("Tuple array", results))
}
