package arx

import (
	"encoding/binary"
)

// 128-bit 加算 (Little Endian)
func Add128(a, b []byte) []byte {
	res := make([]byte, 16)
	var carry uint16 = 0
	for i := 0; i < 16; i++ {
		sum := uint16(a[i]) + uint16(b[i]) + carry
		res[i] = byte(sum & 0xFF)
		carry = sum >> 8
	}
	return res
}

// 128-bit 減算
func Sub128(a, b []byte) []byte {
	res := make([]byte, 16)
	var borrow int16 = 0
	for i := 0; i < 16; i++ {
		sub := int16(a[i]) - int16(b[i]) - borrow
		if sub < 0 {
			sub += 0x100
			borrow = 1
		} else {
			borrow = 0
		}
		res[i] = byte(sub & 0xFF)
	}
	return res
}

// 128-bit XOR
func Xor128(a, b []byte) []byte {
	res := make([]byte, 16)
	for i := 0; i < 16; i++ {
		res[i] = a[i] ^ b[i]
	}
	return res
}

// Rn(4byte) -> 8byte 拡張
func ExpandRn8(rn []byte) []byte {
	res := make([]byte, 8)
	copy(res[0:4], rn)
	copy(res[4:8], rn)
	return res
}

// Rn(4byte) -> 16byte 拡張
func ExpandRn16(rn []byte) []byte {
	res := make([]byte, 16)
	for i := 0; i < 4; i++ {
		copy(res[i*4:], rn)
	}
	return res
}

// 32bit -> 128bit 拡張
func Expand32to128(val uint32) []byte {
	res := make([]byte, 16)
	binary.LittleEndian.PutUint32(res, val)
	return res
}

// CryptStream: マスクによるXOR暗号化/復号
func CryptStream(data, mask []byte) []byte {
	out := make([]byte, len(data))
	for i := range data {
		out[i] = data[i] ^ mask[i%16]
	}
	return out
}