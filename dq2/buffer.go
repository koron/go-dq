package dq2

import "bytes"

type Buffer struct {
	bytes.Buffer

	vbits byte
	nbits int
	err   error
}

var bitMasks = []int{0, 0x01, 0x03, 0x07, 0x0f, 0x1f, 0x3f, 0x7f, 0xff}

func (b *Buffer) WriteBits(v int, n int) error {
	if b.err != nil {
		return b.err
	}
	for n > 0 {
		o := 8 - b.nbits
		m := n
		if m > o {
			m = o
		}
		b.vbits |= byte((v>>uint(n-m))&bitMasks[m]) << uint(o-m)
		b.nbits += m
		if b.nbits == 8 {
			err := b.WriteByte(b.vbits)
			if err != nil {
				b.err = err
				return err
			}
			b.vbits = 0
			b.nbits = 0
		}
		n -= m
	}
	return nil
}

func (b *Buffer) Flush() error {
	if b.err != nil {
		return b.err
	}
	if b.nbits > 0 {
		err := b.WriteByte(b.vbits)
		if err != nil {
			b.err = err
			return err
		}
		b.vbits = 0
		b.nbits = 0
	}
	return nil
}
