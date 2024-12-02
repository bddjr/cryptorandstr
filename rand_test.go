package cryptorandstr_test

import (
	"fmt"
	"testing"

	"github.com/bddjr/cryptorandstr"
)

const strLen = 27
const n = 5

func TestMustRand10(t *testing.T) {
	const strLen = 6
	fmt.Println("MustRand10", strLen)
	for i := 0; i < n; i++ {
		fmt.Println(" ", cryptorandstr.MustRand10(strLen))
	}
	fmt.Println()
}

func TestMustRand16(t *testing.T) {
	fmt.Println("MustRand16", strLen)
	for i := 0; i < n; i++ {
		fmt.Println(" ", cryptorandstr.MustRand16(strLen))
	}
	fmt.Println()
}

func TestMustRand32(t *testing.T) {
	fmt.Println("MustRand32", strLen)
	for i := 0; i < n; i++ {
		fmt.Println(" ", cryptorandstr.MustRand32(strLen))
	}
	fmt.Println()
}

func TestMustRand64(t *testing.T) {
	fmt.Println("MustRand64", strLen)
	for i := 0; i < n; i++ {
		fmt.Println(" ", cryptorandstr.MustRand64(strLen))
	}
	fmt.Println()
}
