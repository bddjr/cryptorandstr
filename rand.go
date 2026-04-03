// Quickly generate secure random string.
package cryptorandstr

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

const (
	Chars10 = "0123456789"

	// hex
	Chars16 = "0123456789abcdef"

	// Chars36 without "0123456789"
	Chars26 = "abcdefghijklmnopqrstuvwxyz"

	// Chars36 without "1ijl"
	Chars32 = "023456789abcdefghkmnopqrstuvwxyz"

	// Chars36 without "1l"
	Chars34 = "023456789abcdefghijkmnopqrstuvwxyz"

	Chars36 = "0123456789abcdefghijklmnopqrstuvwxyz"

	// Chars64 without "0123456789-_"
	Chars52 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Chars64 without "01lO-_"
	Chars58 = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	// Chars64 without "-_"
	Chars62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	Chars64 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"

	Chars64NoUpperCase = "!#$%&()*,-.:;<>?@[]^_`{|}~abcdefghijklmnopqrstuvwxyz0123456789+/"

	// Chars69 without " \"\\"
	Chars66 = "!#$%&'()*+,-./0123456789:;<=>?@[]^_`abcdefghijklmnopqrstuvwxyz{|}~"

	// Chars69 without " "
	Chars68 = "!\"#$%&'()*+,-./0123456789:;<=>?@[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"

	Chars69 = " !\"#$%&'()*+,-./0123456789:;<=>?@[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"

	// Chars95 without " \"\\"
	Chars92 = "!#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[]^_`abcdefghijklmnopqrstuvwxyz{|}~"

	// Chars95 without " "
	Chars94 = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"

	Chars95 = " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
)

// Panic when error.
func MustRand10(strLen int) string {
	return MustRand(strLen, Chars10)
}

// Panic when error.
func MustRand16(strLen int) string {
	s, err := Rand16(strLen)
	if err != nil {
		panic(err)
	}
	return s
}

// Panic when error.
func MustRand26(strLen int) string {
	return MustRand(strLen, Chars26)
}

// Panic when error.
func MustRand32(strLen int) string {
	return MustRand(strLen, Chars32)
}

// Panic when error.
func MustRand34(strLen int) string {
	return MustRand(strLen, Chars34)
}

// Panic when error.
func MustRand36(strLen int) string {
	return MustRand(strLen, Chars36)
}

// Panic when error.
func MustRand52(strLen int) string {
	return MustRand(strLen, Chars52)
}

// Panic when error.
func MustRand58(strLen int) string {
	return MustRand(strLen, Chars58)
}

// Panic when error.
func MustRand62(strLen int) string {
	return MustRand(strLen, Chars62)
}

// Panic when error.
func MustRand64(strLen int) string {
	return MustRand(strLen, Chars64)
}

// Panic when error.
func MustRand64NoUpperCase(strLen int) string {
	return MustRand(strLen, Chars64NoUpperCase)
}

// Panic when error.
func MustRand66(strLen int) string {
	return MustRand(strLen, Chars66)
}

// Panic when error.
func MustRand68(strLen int) string {
	return MustRand(strLen, Chars68)
}

// Panic when error.
func MustRand69(strLen int) string {
	return MustRand(strLen, Chars69)
}

// Panic when error.
func MustRand92(strLen int) string {
	return MustRand(strLen, Chars92)
}

// Panic when error.
func MustRand94(strLen int) string {
	return MustRand(strLen, Chars94)
}

// Panic when error.
func MustRand95(strLen int) string {
	return MustRand(strLen, Chars95)
}

// Panic when error.
func MustRand(strLen int, chars string) string {
	s, err := random(strLen, chars)
	if err != nil {
		panic(err)
	}
	return s
}

func Rand10(strLen int) (string, error) {
	return random(strLen, Chars10)
}

func Rand16(strLen int) (string, error) {
	if strLen < 1 {
		return "", fmt.Errorf("cryptorandstr: strLen %d < 1", strLen)
	}
	src := make([]byte, (strLen/2)+(strLen&1))
	dst := make([]byte, len(src)*2)
	rand.Read(src)
	hex.Encode(dst, src)
	return string(dst[:strLen]), nil
}

func Rand26(strLen int) (string, error) {
	return random(strLen, Chars26)
}

func Rand32(strLen int) (string, error) {
	return random(strLen, Chars32)
}

func Rand34(strLen int) (string, error) {
	return random(strLen, Chars34)
}

func Rand36(strLen int) (string, error) {
	return random(strLen, Chars36)
}

func Rand52(strLen int) (string, error) {
	return random(strLen, Chars52)
}

func Rand58(strLen int) (string, error) {
	return random(strLen, Chars58)
}

func Rand62(strLen int) (string, error) {
	return random(strLen, Chars62)
}

func Rand64(strLen int) (string, error) {
	return random(strLen, Chars64)
}

func Rand64NoUpperCase(strLen int) (string, error) {
	return random(strLen, Chars64NoUpperCase)
}

func Rand66(strLen int) (string, error) {
	return random(strLen, Chars66)
}

func Rand68(strLen int) (string, error) {
	return random(strLen, Chars68)
}

func Rand69(strLen int) (string, error) {
	return random(strLen, Chars69)
}

func Rand92(strLen int) (string, error) {
	return random(strLen, Chars92)
}

func Rand94(strLen int) (string, error) {
	return random(strLen, Chars94)
}

func Rand95(strLen int) (string, error) {
	return random(strLen, Chars95)
}

func Rand(strLen int, chars string) (string, error) {
	if chars == Chars16 {
		return Rand16(strLen)
	}
	if strLen < 1 {
		return "", fmt.Errorf("cryptorandstr: strLen %d < 1", strLen)
	}
	if len(chars) < 2 {
		return "", fmt.Errorf("cryptorandstr: chars length %d < 2", len(chars))
	}
	if len(chars) > 256 {
		return "", fmt.Errorf("cryptorandstr: chars length %d > 256", len(chars))
	}
	return random(strLen, chars)
}

func random(strLen int, chars string) (string, error) {
	var (
		out = make([]byte, strLen)
		i   int

		chunk = make([]byte, 1)

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
			rand.Read(chunk)
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
				return string(out), nil
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
