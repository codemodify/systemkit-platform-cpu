package platform

import (
	"fmt"
	"runtime"
	"strings"
)

func getOS() OSName {
	runtimeOSName := strings.ToLower(runtime.GOOS)
	runtimeOSName = normalizeOS(runtimeOSName)

	for _, osName := range supportedOses {
		osNameAsLowerCase := strings.ToLower(fmt.Sprintf("%v", osName))
		if strings.Index(osNameAsLowerCase, runtimeOSName) != -1 {
			return osName
		}
	}

	return OS_Uknown
}

func normalizeOS(os string) string {
	os = strings.ToLower(os)

	switch os {
	case "macos":
		os = OS_Darwin
	}

	return os
}

func getCPUArchitecture() CPUArchitecture {
	runtimeArchName := strings.ToLower(runtime.GOARCH)
	runtimeArchName, _ = normalizeArch(runtimeArchName, "")

	for _, cpuArchitecture := range supportedCPUArchitectures {
		cpuArchitectureAsLowerCase := strings.ToLower(fmt.Sprintf("%v", cpuArchitecture))
		if strings.Index(cpuArchitectureAsLowerCase, runtimeArchName) != -1 {
			return cpuArchitecture
		}
	}

	return CPU_Uknown
}

func normalizeArch(arch, variant string) (string, string) {
	arch, variant = strings.ToLower(arch), strings.ToLower(variant)
	switch arch {
	case "i386":
		arch = "386"
		variant = ""
	case "x86_64", "x86-64":
		arch = "amd64"
		variant = ""
	case "aarch64", "arm64":
		arch = "arm64"
		switch variant {
		case "8", "v8":
			variant = ""
		}
	case "armhf":
		arch = "arm"
		variant = "v7"
	case "armel":
		arch = "arm"
		variant = "v6"
	case "arm":
		switch variant {
		case "", "7":
			variant = "v7"
		case "5", "6", "8":
			variant = "v" + variant
		}
	}

	return arch, variant
}
