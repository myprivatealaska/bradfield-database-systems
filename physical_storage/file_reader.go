package physical_storage

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/myprivatealaska/bradfield-database-systems/common"
)

type FileReader struct {
	Reader         io.Reader
	ColumnsNum     int
	ColumnMetadata map[string]common.SupportedAttributeTypes
	NumTuples      int
	TuplesRead     int
}

func (fr *FileReader) readHeader() error {
	buf := make([]byte, INT_SIZE_BYTES)
	if _, err := io.ReadFull(fr.Reader, buf); err != nil {
		return err
	}
	fr.ColumnsNum = int(binary.BigEndian.Uint64(buf))

	fmt.Printf("Column num: %v", fr.ColumnsNum)

	return nil
}
