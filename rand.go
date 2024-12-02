package cryptorandstr

import (
	"crypto/rand"
	"fmt"
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

// panic when error.
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

	chunk := make([]byte, 1)
	maxB := byte(len(chars) - 1)
	r := randWithGC{
		outBitLen: bitLen(maxB),
	}
	difference := 8 - r.outBitLen

	out := make([]byte, strLen)
	i := 0
	for i < strLen {
		b, err := r.byte(chunk, difference)
		if err != nil {
			return "", err
		}
		if b <= maxB {
			out[i] = chars[b]
			i++
		}
	}

	return string(out), nil
}

func bitLen(b byte) (l byte) {
	for b != 0 {
		b >>= 1
		l++
	}
	return
}

type randWithGC struct {
	outBitLen   byte
	cacheBitLen byte
	cache       uint16
}

func (r *randWithGC) byte(chunk []byte, difference byte) (byte, error) {
	// read cache
	if r.cacheBitLen >= r.outBitLen {
		b := (byte(r.cache) << difference) >> difference
		r.cache >>= r.outBitLen
		r.cacheBitLen -= r.outBitLen
		return b, nil
	}

	// random
	for {
		n, err := rand.Reader.Read(chunk)
		if err != nil {
			return 0, err
		}
		if n == 1 {
			break
		}
	}

	// write cache
	r.cache <<= difference
	r.cache |= uint16(chunk[0] >> r.outBitLen)
	r.cacheBitLen += difference

	return (chunk[0] << difference) >> difference, nil
}
