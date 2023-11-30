package bytereader

import (
	"io"
)

var _ io.ByteReader = &internalByteReaderFromString{}

type internalByteReaderFromString struct {
	value string
	index int
}

// NewByteReaderFromString returns a new io.ByteReader from a string.
//
// Example usage:
//
//	var s string = "ONCE 🙂 TWICE 🙂 THRICE 🙂 FOURCE"
//
//	var byteReader io.ByteReader = bytereader.NewByteReaderFromString(s)
func NewByteReaderFromString(value string) io.ByteReader {
	return &internalByteReaderFromString{
		value:value,
	}
}

func (receiver *internalByteReaderFromString) ReadByte() (byte, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	var value string = receiver.value
	var index int    = receiver.index

	if len(value) <= index {
		return 0, io.EOF
	}

	var b byte = receiver.value[index]
	receiver.index++

	return b, nil
}
