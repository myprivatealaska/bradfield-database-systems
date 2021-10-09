package physical_storage

import (
	"bytes"
	"io"
	"testing"
)

func TestFileScanner(t *testing.T) {

	//columns := []string{"name", "age"}
	//
	//t1 := Tuple{Attributes: []Attribute{
	//	{
	//		Name:  "name",
	//		Value: "Aya",
	//	},
	//	{
	//		Name:  "age",
	//		Value: 21,
	//	},
	//}}
	//
	//t2 := Tuple{Attributes: []Attribute{
	//	{
	//		Name:  "name",
	//		Value: "Jay",
	//	},
	//	{
	//		Name:  "age",
	//		Value: 24,
	//	},
	//}}
	//
	//tuples := []Tuple{t1, t2}

	//fileReader := NewMemoryTestFile(t, columns, tuples)

	//var (
	//	fileReader  = NewMemoryTestFile(t, columns, tuples)
	//	fileScanner = NewFileScanner(fileReader)
	//)
	//for _, tuple := range tuples {
	//	assertEq(t, true, fileScanner.Next())
	//	assertEq(t, tuple, fileScanner.Execute())
	//}
}

// NewMemoryTestFile creates a new Bradfield file format file with the provided tuples
// that is stored in memory.
func NewMemoryTestFile(t *testing.T, columns []string, tuples []Tuple) io.Reader {
	colMeta := map[string]interface{}{
		"name": "string",
		"age":  "int64",
	}
	var (
		buf    = bytes.NewBuffer(nil)
		writer = NewFileWriter(buf, colMeta, len(tuples))
	)

	for _, tuple := range tuples {
		if err := writer.WriteTuple(tuple); err != nil {
			t.Fatalf("error writing tuple: %v, err: %v", tuple, err)
		}
	}

	return buf
}
