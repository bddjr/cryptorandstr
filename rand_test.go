package cryptorandstr_test

import (
	"fmt"
	"testing"

	"github.com/bddjr/cryptorandstr"
)

const strLen = 17
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

func TestMustRand26(t *testing.T) {
	fmt.Println("MustRand26", strLen)
	for i := 0; i < n; i++ {
		fmt.Println(" ", cryptorandstr.MustRand26(strLen))
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

func TestMustRand34(t *testing.T) {
	fmt.Println("MustRand34", strLen)
	for i := 0; i < n; i++ {
		fmt.Println(" ", cryptorandstr.MustRand34(strLen))
	}
	fmt.Println()
}

func TestMustRand36(t *testing.T) {
	fmt.Println("MustRand36", strLen)
	for i := 0; i < n; i++ {
		fmt.Println(" ", cryptorandstr.MustRand36(strLen))
	}
	fmt.Println()
}

func TestMustRand52(t *testing.T) {
	fmt.Println("MustRand52", strLen)
	for i := 0; i < n; i++ {
		fmt.Println(" ", cryptorandstr.MustRand52(strLen))
	}
	fmt.Println()
}

func TestMustRand58(t *testing.T) {
	fmt.Println("MustRand58", strLen)
	for i := 0; i < n; i++ {
		fmt.Println(" ", cryptorandstr.MustRand58(strLen))
	}
	fmt.Println()
}

func TestMustRand62(t *testing.T) {
	fmt.Println("MustRand62", strLen)
	for i := 0; i < n; i++ {
		fmt.Println(" ", cryptorandstr.MustRand62(strLen))
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

func TestMustRand64NoUpperCase(t *testing.T) {
	fmt.Println("MustRand64NoUpperCase", strLen)
	for i := 0; i < n; i++ {
		fmt.Println(" ", cryptorandstr.MustRand64NoUpperCase(strLen))
	}
	fmt.Println()
}

func TestMustRand66(t *testing.T) {
	fmt.Println("MustRand66", strLen)
	for i := 0; i < n; i++ {
		fmt.Println(" ", cryptorandstr.MustRand66(strLen))
	}
	fmt.Println()
}

func TestMustRand68(t *testing.T) {
	fmt.Println("MustRand68", strLen)
	for i := 0; i < n; i++ {
		fmt.Println(" ", cryptorandstr.MustRand68(strLen))
	}
	fmt.Println()
}
func TestMustRand69(t *testing.T) {
	fmt.Println("MustRand69", strLen)
	for i := 0; i < n; i++ {
		fmt.Println(" ", cryptorandstr.MustRand69(strLen))
	}
	fmt.Println()
}

func TestMustRand92(t *testing.T) {
	fmt.Println("MustRand92", strLen)
	for i := 0; i < n; i++ {
		fmt.Println(" ", cryptorandstr.MustRand92(strLen))
	}
	fmt.Println()
}

func TestMustRand94(t *testing.T) {
	fmt.Println("MustRand94", strLen)
	for i := 0; i < n; i++ {
		fmt.Println(" ", cryptorandstr.MustRand94(strLen))
	}
	fmt.Println()
}

func TestMustRand95(t *testing.T) {
	fmt.Println("MustRand95", strLen)
	for i := 0; i < n; i++ {
		fmt.Println(" ", cryptorandstr.MustRand95(strLen))
	}
	fmt.Println()
}
