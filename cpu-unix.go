// +build !windows

package cpu

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// Info -
func Info() CPU {
	count := fetchCPUInfo2("Socket(s):")
	countAsInt, _ := strconv.Atoi(count)

	coresPerCPUCount := fetchCPUInfo2("Core(s) per socket:")
	coresPerCPUCountAsInt, _ := strconv.Atoi(coresPerCPUCount)

	threadsPerCoreCount := fetchCPUInfo2("Thread(s) per core:")
	threadsPerCoreCountAsInt, _ := strconv.Atoi(threadsPerCoreCount)

	cpuInfo := CPU{
		Count:          countAsInt,
		CoresPerCPU:    coresPerCPUCountAsInt,
		ThreadsPerCore: threadsPerCoreCountAsInt,
		TotalThreads:   runtime.NumCPU(),
		Architecture:   getCPUArchitecture(),
		Variant: CPUVariant{
			Name:     getCPUVariant(),
			Detailed: getCPUVariantDetailed(),
		},
		Manufacturer: fetchCPUInfo2("Vendor ID:"),
		ByteOrder:    fetchCPUInfo2("Byte Order:"),
		Features: []string{
			fetchCPUInfo2("Flags:"),
		},
		AdditionalInfo: AdditionalInfo{
			"goos":          runtime.GOOS,
			"goarch":        runtime.GOARCH,
			"model name":    fetchCPUInfo("model name"),
			"Model name":    fetchCPUInfo2("Model name:"),
			"bugs":          fetchCPUInfo("bugs"),
			"Vulnerability": fetchCPUInfo2All("Vulnerability"),
		},
	}

	cpuInfo.setID()

	return cpuInfo
}

// On Linux
//		The kernel has already detected the ABI, ISA and Features.
//		We don't need to access the ARM registers to detect platform information.
//		We can just parse these information from /proc/cpuinfo
func fetchCPUInfo(pattern string) string {
	cpuinfo, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return ""
	}
	defer cpuinfo.Close()

	// Parse the Cpuinfo line by line. For SMP SoC, we parse the first core is enough.
	scanner := bufio.NewScanner(cpuinfo)
	for scanner.Scan() {
		newline := scanner.Text()
		list := strings.Split(newline, ":")

		if len(list) > 1 && strings.EqualFold(strings.TrimSpace(list[0]), pattern) {
			return strings.TrimSpace(list[1])
		}
	}

	// Check whether the scanner encountered errors
	err = scanner.Err()
	if err != nil {
		return ""
	}

	return ""
}

type LscpuFieldData struct {
	Field string `json:"field"`
	Data  string `json:"data"`
}

type LscpuOutput struct {
	Lscpu []LscpuFieldData `json:"lscpu"`
}

func fetchCPUInfo2(pattern string) string {
	cmd := exec.Command("lscpu", "--json")
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return ""
	}

	outputAsString := strings.TrimSpace(out.String())

	// fmt.Println(outputAsString)

	lscpuO := LscpuOutput{}
	err = json.Unmarshal([]byte(outputAsString), &lscpuO)
	if err != nil {
		return ""
	}

	for _, v := range lscpuO.Lscpu {
		if strings.Index(v.Field, pattern) != -1 {
			return v.Data
		}
	}

	return ""
}

func fetchCPUInfo2All(pattern string) []string {
	cmd := exec.Command("lscpu", "--json")
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return []string{}
	}

	type lscpuFieldData struct {
		field string `json:"field"`
		data  string `json:"data"`
	}

	type lscpuOutput struct {
		lscpu []lscpuFieldData `json:"lscpu"`
	}

	outputAsString := strings.TrimSpace(out.String())
	var lscpuO lscpuOutput
	err = json.Unmarshal([]byte(outputAsString), &lscpuO)
	if err != nil {
		return []string{}
	}

	theList := []string{}
	for _, v := range lscpuO.lscpu {
		if strings.Index(v.field, pattern) != -1 {
			theList = append(theList, fmt.Sprintf("%s : %s", v.field, v.data))
		}
	}

	return theList
}

func getPhysicalCPUCount() int {
	cpuinfo, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return 0
	}
	defer cpuinfo.Close()

	pattern := "physical id"
	scanner := bufio.NewScanner(cpuinfo)
	allLines := []string{}
	for scanner.Scan() {
		newline := scanner.Text()
		if strings.Index(newline, pattern) != -1 {
			list := strings.Split(newline, ":")
			if len(list) > 1 && strings.EqualFold(strings.TrimSpace(list[0]), pattern) {
				allLines = append(allLines, strings.TrimSpace(list[1]))
			}
		}
	}

	// Check whether the scanner encountered errors
	err = scanner.Err()
	if err != nil {
		return 0
	}

	if len(allLines) > 0 {
		cpuCountAsInt, err := strconv.Atoi(allLines[len(allLines)-1])
		if err == nil {
			return cpuCountAsInt + 1
		}
	}

	return 0
}

func getCPUVariant() CPUArchitectureVariant {
	variantAsString := fetchCPUInfo("Cpu architecture")
	if variantAsString == "" {
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
	variantAsString := fetchCPUInfo("Cpu architecture")
	if variantAsString == "" {
		return CPUVD_Uknown
	}

	switch strings.ToLower(variantAsString) {
	case "7m":
		return CPUVD_ARMv7_M
	case "?(12)", "?(13)", "?(14)", "?(15)", "?(16)", "?(17)":
		return CPUV_ARMv7

	case "6tej":
		return CPUVD_ARMv6TEJ

	case "5t":
		return CPUVD_ARMv5T
	case "5te":
		return CPUVD_ARMv5TE
	case "5tej":
		return CPUVD_ARMv5TEJ

	case "4t":
		return CPUVD_ARMv4T
	}

	return CPUVD_Uknown
}

func getKernelVersion() string {
	cmd := exec.Command("uname", "-r")
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(out.String())
}
