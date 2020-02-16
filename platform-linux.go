// +build linux
// +build !android

package platform

import (
	"bufio"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/containerd/containerd/errdefs"
	"github.com/pkg/errors"
)

// GetInfo -
func GetInfo() Info {
	return linuxInfo{}
}

// LinuxInfo -
type linuxInfo struct{}

// Platform -
func (thisRef linuxInfo) Platform() Platform {
	coresCount, _ := getCPUInfo("cpu cores")
	coresCountAsInt, _ := strconv.Atoi(coresCount)
	return Platform{
		OS: OS{
			Name:     getOS(),
			Version:  "",
			Features: []string{},
		},
		CPU: CPU{
			Count:        1,
			CoresCount:   coresCountAsInt,
			TotalThreads: runtime.NumCPU(),
			Architecture: getCPUArchitecture(),
			Variant: CPUVariant{
				Name:     getCPUVariant(),
				Detailed: getCPUVariantDetailed(),
			},
			Manufacturer: "intel",
			Features:     []string{},
		},
		Metadata: Metadata{
			"goos":   runtime.GOOS,
			"goarch": runtime.GOOS,
		},
	}
}

// On Linux
//		The kernel has already detected the ABI, ISA and Features.
//		We don't need to access the ARM registers to detect platform information.
//		We can just parse these information from /proc/cpuinfo
func getCPUInfo(pattern string) (info string, err error) {
	cpuinfo, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return "", err
	}
	defer cpuinfo.Close()

	// Parse the Cpuinfo line by line. For SMP SoC, we parse the first core is enough.
	scanner := bufio.NewScanner(cpuinfo)
	for scanner.Scan() {
		newline := scanner.Text()
		list := strings.Split(newline, ":")

		if len(list) > 1 && strings.EqualFold(strings.TrimSpace(list[0]), pattern) {
			return strings.TrimSpace(list[1]), nil
		}
	}

	// Check whether the scanner encountered errors
	err = scanner.Err()
	if err != nil {
		return "", err
	}

	return "", errors.Wrapf(errdefs.ErrNotFound, "getCPUInfo for pattern: %s", pattern)
}

func getCPUVariant() CPUArchitectureVariant {
	variantAsString, err := getCPUInfo("Cpu architecture")
	if err != nil {
		return ""
	}

	switch strings.ToLower(variantAsString) {
	case "8", "aarch64":
		// special case:
		//	if running a 32-bit userspace on aarch64,
		//	the variant should be "v7"
		if runtime.GOARCH == "arm" {
			return CPUV_ARMv7
		}

		return CPUV_ARMv8

	case "7", "7m", "?(12)", "?(13)", "?(14)", "?(15)", "?(16)", "?(17)":
		return CPUV_ARMv7

	case "6", "6tej":
		return CPUV_ARMv6

	case "5", "5t", "5te", "5tej":
		return CPUV_ARMv5

	case "4", "4t":
		return CPUV_ARMv4

	case "3":
		return CPUV_ARMv3
	}

	return CPUV_Uknown
}

func getCPUVariantDetailed() CPUArchitectureVariantDetailed {
	variantAsString, err := getCPUInfo("Cpu architecture")
	if err != nil {
		return ""
	}

	switch strings.ToLower(variantAsString) {
	case "8", "aarch64":
		// special case:
		//	if running a 32-bit userspace on aarch64,
		//	the variant should be "v7"
		if runtime.GOARCH == "arm" {
			return CPUV_ARMv7
		}

		return CPUV_ARMv8

	case "7":
		return CPUV_ARMv7
	case "7m":
		return CPUV_ARMv7_M
	case "?(12)", "?(13)", "?(14)", "?(15)", "?(16)", "?(17)":
		return CPUV_ARMv7

	case "6":
		return CPUV_ARMv6
	case "6tej":
		return CPUV_ARMv6TEJ

	case "5":
		return CPUV_ARMv5
	case "5t":
		return CPUV_ARMv5T
	case "5te":
		return CPUV_ARMv5TE
	case "5tej":
		return CPUV_ARMv5TEJ

	case "4":
		return CPUV_ARMv4
	case "4t":
		return CPUV_ARMv4T

	case "3":
		return CPUV_ARMv3
	}

	return CPUV_Uknown
}
