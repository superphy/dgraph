package codec

import (
	"encoding/binary"
	mathbits "math/bits"
)

// Encode4 is used by group variant function.
func Encode4(dst []byte, src []uint32) []byte {

	var bits uint8
	var n, b uint32

	offs := uint32(1)

	n = src[0]
	binary.LittleEndian.PutUint32(dst[offs:], n)
	b = 3 - uint32(mathbits.LeadingZeros32(n|1)/8)
	bits |= byte(b)
	offs += b + 1

	n = src[1]
	binary.LittleEndian.PutUint32(dst[offs:], n)
	b = 3 - uint32(mathbits.LeadingZeros32(n|1)/8)
	bits |= byte(b) << 2
	offs += b + 1

	n = src[2]
	binary.LittleEndian.PutUint32(dst[offs:], n)
	b = 3 - uint32(mathbits.LeadingZeros32(n|1)/8)
	bits |= byte(b) << 4
	offs += b + 1

	n = src[3]
	binary.LittleEndian.PutUint32(dst[offs:], n)
	b = 3 - uint32(mathbits.LeadingZeros32(n|1)/8)
	bits |= byte(b) << 6
	offs += b + 1

	dst[0] = bits

	return dst[:offs]
}

// encodeGroupVariant for encoding []uint to byte array
func encodeGroupVarint(input []uint32) []byte {

	var r []byte

	var padding int
	for len(input) > 0 {
		var dst [17]byte

		d := Encode4(dst[:], input)

		padding = 17 - len(d)

		r = append(r, d...)

		input = input[4:]
	}

	// must be able to load 17 bytes from start of final block
	for i := 0; i < padding; i++ {
		r = append(r, 0)
	}

	return r
}
