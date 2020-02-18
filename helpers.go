package cpu

import (
	"fmt"
	"runtime"
	"strings"
)

func getCPUArchitectureFromString(archNameToCheck string) CPUArchitecture {
	archNameToCheck = strings.ToLower(archNameToCheck)
	archNameToCheck, _ = normalizeArch(archNameToCheck, "")

	for _, cpuArchitecture := range supportedCPUArchitectures {
		cpuArchitectureAsLowerCase := strings.ToLower(fmt.Sprintf("%v", cpuArchitecture))
		if strings.Index(cpuArchitectureAsLowerCase, archNameToCheck) != -1 {
			return cpuArchitecture
		}
	}

	return CPU_Uknown
}

func getCPUArchitecture() CPUArchitecture {
	return getCPUArchitectureFromString(runtime.GOARCH)
}

func getCPUArchitectureVariantFromString(archNameToCheck string) CPUArchitectureVariant {
	archNameToCheck = strings.ToLower(archNameToCheck)
	archNameToCheck, _ = normalizeArch(archNameToCheck, "")

	for _, cpuArchitectureV := range supportedCPUArchitectureVariants {
		cpuArchitectureVAsLowerCase := strings.ToLower(fmt.Sprintf("%v", cpuArchitectureV))
		if strings.Index(cpuArchitectureVAsLowerCase, archNameToCheck) != -1 {
			return cpuArchitectureV
		}
	}

	return CPUV_Uknown
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
