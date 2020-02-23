// +build windows

package cpu

import (
	"runtime"

	"github.com/StackExchange/wmi"
)

func Info() CPU {
	// fetch info from WMI
	var win32descriptions []win32Processor
	if err := wmi.Query(wmqlProcessor, &win32descriptions); err != nil {
		return CPU{
			Architecture: CPU_Uknown,
		}
	}

	if len(win32descriptions) <= 0 {
		return CPU{
			Architecture: CPU_Uknown,
		}
	}

	return CPU{
		Count:          len(win32descriptions),
		CoresPerCPU:    int(win32descriptions[0].NumberOfCores),
		ThreadsPerCore: int(win32descriptions[0].NumberOfLogicalProcessors),
		TotalThreads:   runtime.NumCPU(),
		Architecture:   getCPUArchitecture(),
		Variant: CPUVariant{
			Name: getCPUArchitectureVariantFromString(win32descriptions[0].Name),
		},
		Manufacturer: win32descriptions[0].Manufacturer,
		AdditionalInfo: AdditionalInfo{
			"goos":   runtime.GOOS,
			"goarch": runtime.GOARCH,
		},
	}
}

const wmqlProcessor = "SELECT Manufacturer, Name, NumberOfLogicalProcessors, NumberOfCores FROM Win32_Processor"

type win32Processor struct {
	Manufacturer              string
	Name                      string
	NumberOfLogicalProcessors uint32
	NumberOfCores             uint32
}
