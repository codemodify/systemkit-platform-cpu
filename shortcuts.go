package platform

import (
	"fmt"
	"strings"
)

// IsArm -
func IsArm(v CPUArchitecture) bool {
	return strings.Index(fmt.Sprint("", v), "arm") != -1
}

// Is64bit -
func Is64bit(v CPUArchitecture) bool {
	return strings.Index(fmt.Sprint("", v), "64") != -1
}

// IsMips -
func IsMips(v CPUArchitecture) bool {
	return strings.Index(fmt.Sprint("", v), "mips") != -1
}

// IsPpc -
func IsPpc(v CPUArchitecture) bool {
	return strings.Index(fmt.Sprint("", v), "ppc") != -1
}

// IsSparc -
func IsSparc(v CPUArchitecture) bool {
	return strings.Index(fmt.Sprint("", v), "sparc") != -1
}
