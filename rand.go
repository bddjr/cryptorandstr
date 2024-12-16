package cryptorandstr

import (
	"crypto/rand"
	"fmt"
	"unsafe"
)

const Chars64 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"

var Chars10 = Chars64[:10]
var Chars16 = Chars64[:16]
var Chars32 = Chars64[:32]

// Panic when error.
func MustRand10(strLen int) string {
	return MustRand(strLen, Chars10)
}

// Panic when error.
func MustRand16(strLen int) string {
	return MustRand(strLen, Chars16)
}

// Panic when error.
func MustRand32(strLen int) string {
	return MustRand(strLen, Chars32)
}

// Panic when error.
func MustRand64(strLen int) string {
	return MustRand(strLen, Chars64)
}

// Panic when error.
func MustRand(strLen int, chars string) string {
	s, err := Rand(strLen, chars)
	if err != nil {
		panic(err)
	}
	return s
}

func Rand10(strLen int) (string, error) {
	return Rand(strLen, Chars10)
}

func Rand16(strLen int) (string, error) {
	return Rand(strLen, Chars16)
}

func Rand32(strLen int) (string, error) {
	return Rand(strLen, Chars32)
}

func Rand64(strLen int) (string, error) {
	return Rand(strLen, Chars64)
}

func Rand(strLen int, chars string) (string, error) {
	if strLen < 1 {
		return "", fmt.Errorf("cryptorandstr: strLen %d < 1", strLen)
	}
	if len(chars) < 2 {
		return "", fmt.Errorf("cryptorandstr: chars length %d < 2", len(chars))
	}
	if len(chars) > 256 {
		return "", fmt.Errorf("cryptorandstr: chars length %d > 256", len(chars))
	}

	var (
		out = make([]byte, strLen)
		i   int

		chunk = make([]byte, 1)
		n     int
		err   error

		cache       uint16
		cacheBitLen byte

		maxB       = byte(len(chars) - 1)
		outBitLen  = bitLen(maxB)
		difference = 8 - outBitLen
		b          byte
	)

	for {
		if cacheBitLen >= outBitLen {
			// read cache
			b = byte(cache)
			cache >>= outBitLen
			cacheBitLen -= outBitLen
		} else {
			// random
			for {
				n, err = rand.Reader.Read(chunk)
				if err != nil {
					return "", err
				}
				if n == 1 {
					break
				}
			}
			b = chunk[0]
			// write cache
			cache <<= difference
			cache |= uint16(b >> outBitLen)
			cacheBitLen += difference
		}

		b <<= difference
		b >>= difference
		if b <= maxB {
			out[i] = chars[b]
			i++
			if i == strLen {
				return *(*string)(unsafe.Pointer(&out)), nil
			}
		}
	}
}

func bitLen(b byte) (l byte) {
	for b != 0 {
		b >>= 1
		l++
	}
	return
}
