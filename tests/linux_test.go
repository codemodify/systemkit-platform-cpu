// +build linux
// +build !android

package tests

import (
	"fmt"
	"testing"

	platforms "github.com/codemodify/systemkit-platform"
)

func TestCPUVariant(t *testing.T) {
	info := platforms.GetInfo()

	fmt.Println(info.Platform())
}
