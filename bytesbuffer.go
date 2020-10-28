package bytesbuffer

import (
	"errors"
)

type Buffer struct {
	buffer []byte
	index  int
	length int
}

var (
	ErrorOutOfBounds = errors.New("out of bounds")
	ErrorEmpty       = errors.New("empty string")
)

func New(length int) *Buffer {
	return &Buffer{
		buffer: make([]byte, length),
		length: length,
		index:  0,
	}
}

func (buffer *Buffer) Write(s []byte) error {
	var length = len(s)
	if length == 0 {
		return ErrorEmpty
	}
	if (buffer.index + length) <= buffer.length {
		copy(buffer.buffer[buffer.index:], s)
		buffer.index += length
		return nil
	}
	return ErrorOutOfBounds
}

func (buffer *Buffer) WriteWithIndex(index int, s []byte) error {
	var length = len(s)
	if length == 0 {
		return ErrorEmpty
	}
	length += index
	if length <= buffer.length {
		copy(buffer.buffer[index:], s)
		buffer.index = length
		return nil
	}
	return ErrorOutOfBounds
}

func (buffer *Buffer) Get() []byte {
	return buffer.buffer[:buffer.index]
}

func (buffer *Buffer) GetPtr() *[]byte {
	return &buffer.buffer
}

func (buffer *Buffer) SetIndex(index int) error {
	if index >= 0 && index <= buffer.length {
		buffer.index = index
		return nil
	}
	return ErrorOutOfBounds
}

func (buffer *Buffer) GetLength() int {
	return buffer.index
}

func (buffer *Buffer) GetMaxLength() int {
	return buffer.length
}

func (buffer *Buffer) Reset() {
	buffer.index = 0
}

func (buffer *Buffer) Close() {
	buffer.index = 0
	buffer.length = 0
	buffer.buffer = buffer.buffer[:0]
}
