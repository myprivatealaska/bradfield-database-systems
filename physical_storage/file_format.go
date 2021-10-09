package physical_storage

import (
	"encoding/binary"
	"errors"
	"io"
)

const OFFSET_SIZE_BYTES = 2
const INT_SIZE_BYTES = 8

type FileWriter struct {
	Writer         io.Writer
	ColumnMetadata map[string]SupportedAttributeTypes
	NumTuples      int
	TuplesWritten  int
}

func NewFileWriter(w io.Writer, colMeta map[string]SupportedAttributeTypes, numTuples int) *FileWriter {
	return &FileWriter{
		Writer:         w,
		ColumnMetadata: colMeta,
		NumTuples:      numTuples,
	}
}

func (fw *FileWriter) writeHeader() error {
	err := binary.Write(fw.Writer, binary.BigEndian, len(fw.ColumnMetadata))
	if err != nil {
		return err
	}
	for k, v := range fw.ColumnMetadata {
		if err = binary.Write(fw.Writer, binary.BigEndian, k); err != nil {
			return err
		}
		if err = binary.Write(fw.Writer, binary.BigEndian, v); err != nil {
			return err
		}
	}
	err = binary.Write(fw.Writer, binary.BigEndian, fw.NumTuples)
	if err != nil {
		return err
	}

	return nil
}

func (fw *FileWriter) WriteTuple(t Tuple) error {

	if fw.TuplesWritten == 0 {
		err := fw.writeHeader()
		if err != nil {
			return err
		}
	}

	offsetBytes := make([]byte, OFFSET_SIZE_BYTES)
	intBytes := make([]byte, INT_SIZE_BYTES)

	var offset uint16

	for _, attr := range t.Attributes {
		switch attr.Value.(type) {
		case uint64:
			offset = 8
			binary.BigEndian.PutUint16(offsetBytes, offset)
			if _, err := fw.Writer.Write(offsetBytes); err != nil {
				return err
			}
			binary.BigEndian.PutUint64(intBytes, attr.Value.(uint64))
			if _, err := fw.Writer.Write(intBytes); err != nil {
				return err
			}
		case string:
			stringBytes := []byte(attr.Value.(string))
			offset = uint16(len(stringBytes))
			binary.BigEndian.PutUint16(offsetBytes, offset)
			if _, err := fw.Writer.Write(offsetBytes); err != nil {
				return err
			}
			if _, err := fw.Writer.Write(stringBytes); err != nil {
				return err
			}
		default:
			return errors.New("attribute type not supported")
		}

	}

	fw.TuplesWritten += 1

	return nil
}
