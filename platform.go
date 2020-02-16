package platform

// Info -
type Info interface {
	Platform() Platform
}

// Metadata -
type Metadata map[string]interface{}

// Platform -
type Platform struct {
	OS       OS       `json:"os,omitempty"`       //
	CPU      CPU      `json:"cpu,omitempty"`      //
	Metadata Metadata `json:"metadata,omitempty"` //
}

// OS -
type OS struct {
	Name     OSName   `json:"name,omitempty"`     // ex: Linux
	Version  string   `json:"version,omitempty"`  // ex: `10.0.14393.1066` on Windows
	Features []string `json:"features,omitempty"` // ex: `win32k` on Windows
}

// CPU -
type CPU struct {
	Count        int             `json:"count,omitempty"`        // nr of physical CPUs
	CoresCount   int             `json:"coresCount,omitempty"`   // nr of CPU Cores from all CPUs
	TotalThreads int             `json:"totalThreads,omitempty"` // nr of CPU Threads from all Cores
	Architecture CPUArchitecture `json:"architecture,omitempty"` // ex: `amd64` or `ppc64`
	Variant      CPUVariant      `json:"variant,omitempty"`      //
	Manufacturer string          `json:"manufacturer,omitempty"` // intel, amd, arm
	Features     []string        `json:"features,omitempty"`     // ex: `SWPB (swap) instructions` for ARM
}

// CPUVariant -
type CPUVariant struct {
	Name     CPUArchitectureVariant         `json:"name,omitempty"`     // ex: ARMv8
	Detailed CPUArchitectureVariantDetailed `json:"detailed,omitempty"` // ex: ARMv8.2-A
}
