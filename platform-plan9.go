// +build windows

package platform

import (
	"runtime"
)

func getCPUVariant() string {
	// Windows only supports v7 for ARM32 and v8 for ARM64
	// We can use runtime.GOARCH to determine the variants

	var variant = "unknown"
	switch runtime.GOARCH {
	case "arm64":
		variant = "v8"
	case "arm":
		variant = "v7"
	}

	return variant

}
