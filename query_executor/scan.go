package query_executor

// Scan which yields each row for the table as needed.
// In this initial implementation your Scan operator can return rows from a predefined list in memory.
type Scan struct {
	BufferSize  int
	Buffer      []Movie
	CurrIndex   int
	BatchNumber int
	InputPtr    *[]Movie
}

func (s *Scan) Init(inputPtr *[]Movie, bufSize int) {
	s.BufferSize = bufSize
	s.InputPtr = inputPtr
	s.loadBuffer()
}

// Next returns nil when there is nothing left to read
func (s *Scan) Next() *Movie {
	if s.CurrIndex >= s.BufferSize {
		s.loadBuffer()
	}
	if len(s.Buffer) < s.CurrIndex {
		return nil
	}
	res := &s.Buffer[s.CurrIndex]
	s.CurrIndex++
	return res
}

// in real world, we would be accessing disk here
func (s *Scan) loadBuffer() {
	start := s.BatchNumber * s.BufferSize
	end := s.BatchNumber*s.BufferSize + s.BufferSize

	inputVal := *s.InputPtr

	veryEnd := len(inputVal)
	if start >= veryEnd {
		s.Buffer = []Movie{}
		return
	}
	if end > veryEnd {
		end = veryEnd
	}

	s.Buffer = inputVal[start:end]
	s.CurrIndex = 0
	s.BatchNumber += 1
}
