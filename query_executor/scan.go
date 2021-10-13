package query_executor

import (
	"github.com/myprivatealaska/bradfield-database-systems/common"
)

// Scan which yields each row for the table as needed.
// In this initial implementation your Scan operator can return rows from a predefined list in memory.
type scan struct {
	BufferSize  int
	Buffer      []common.Tuple
	CurrIndex   int
	BatchNumber int
	InputPtr    *[]common.Tuple
}

// @TODO replace inputPtr with relation
func NewScan(inputPtr *[]common.Tuple, bufSize int) *scan {
	s := &scan{
		BufferSize: bufSize,
		InputPtr:   inputPtr,
	}

	s.loadBuffer()

	return s
}

// Next returns nil when there is nothing left to read
func (s *scan) Next() *common.Tuple {
	if s.CurrIndex >= s.BufferSize {
		s.loadBuffer()
	}
	if len(s.Buffer) <= s.CurrIndex {
		return nil
	}
	res := &s.Buffer[s.CurrIndex]
	s.CurrIndex++
	return res
}

// in real world, we would be accessing disk here
func (s *scan) loadBuffer() {
	start := s.BatchNumber * s.BufferSize
	end := s.BatchNumber*s.BufferSize + s.BufferSize

	inputVal := *s.InputPtr

	veryEnd := len(inputVal)
	if start >= veryEnd {
		s.Buffer = []common.Tuple{}
		return
	}
	if end > veryEnd {
		end = veryEnd
	}

	s.Buffer = inputVal[start:end]
	s.CurrIndex = 0
	s.BatchNumber += 1
}
