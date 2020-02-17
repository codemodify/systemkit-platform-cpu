package platform

// Metadata -
type Metadata map[string]interface{}

// Platform -
type Platform struct {
	OS       OS       `json:"os"`       //
	CPU      CPU      `json:"cpu"`      //
	Metadata Metadata `json:"metadata"` //
}

// OS -
type OS struct {
	Name     OSName   `json:"name"`     // ex: Linux
	Version  string   `json:"version"`  // ex: `10.0.14393.1066` on Windows
	Features []string `json:"features"` // ex: `win32k` on Windows
}

// CPU -
type CPU struct {
	Count          int             `json:"count"`          // nr of physical CPUs
	CoresPerCPU    int             `json:"coresPerCPU"`    // nr of Cores per CPU
	ThreadsPerCore int             `json:"threadsPerCore"` // nr of Threads per Core
	TotalThreads   int             `json:"totalThreads"`   // nr of CPU Threads from all Cores and all CPUs
	Architecture   CPUArchitecture `json:"architecture"`   // ex: `amd64` or `ppc64`
	Variant        CPUVariant      `json:"variant"`        //
	Manufacturer   string          `json:"manufacturer"`   // intel, amd, arm
	ByteOrder      string          `json:"byteOrder"`      //
	Features       []string        `json:"features"`       // ex: `SWPB (swap) instructions` for ARM
}

// CPUVariant -
type CPUVariant struct {
	Name     CPUArchitectureVariant         `json:"name"`     // ex: ARMv8
	Detailed CPUArchitectureVariantDetailed `json:"detailed"` // ex: ARMv8.2-A
}
