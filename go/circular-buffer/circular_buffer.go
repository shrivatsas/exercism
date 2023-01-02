package circular

import (
	"errors"
)

// Buffer is an implementation of a ring buffer
type Buffer struct {
	content []byte
	start   int
	filled  int
}

// NewBuffer creates a new ring buffer of required size
func NewBuffer(size int) *Buffer {
	var buf Buffer
	buf.content = make([]byte, size)
	buf.start = 0
	buf.filled = 0
	return &buf
}

// ReadByte reads a single byte from the buffer and moves forward
func (b *Buffer) ReadByte() (byte, error) {
	if b.filled > 0 {
		val := b.content[b.start]
		b.start = (b.start + 1) % len(b.content)
		b.filled--
		return val, nil
	}
	return byte(0), errors.New("Out of bounds")
}

// WriteByte writes a byte if space permits
func (b *Buffer) WriteByte(c byte) error {
	if b.filled == len(b.content) {
		return errors.New("Buffer is full")
	}
	idx := (b.start + b.filled) % len(b.content)
	b.filled++
	b.content[idx] = c
	return nil
}

// Overwrite writes a byte, and overwrites oldest if full
func (b *Buffer) Overwrite(c byte) {
	idx := (b.start + b.filled) % len(b.content)
	if b.filled == len(b.content) {
		b.start = (b.start + 1) % len(b.content)
	} else {
		b.filled++
	}
	b.content[idx] = c
}

// Reset clears the buffer
func (b *Buffer) Reset() {
	b.content = make([]byte, len(b.content))
	b.start = 0
	b.filled = 0
}
