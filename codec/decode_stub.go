// +build amd64
// +build !noasm

package codec

//go:noescape
func Decode4(dst []uint32, src []byte)
