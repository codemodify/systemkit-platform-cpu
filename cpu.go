package cpu

// AdditionalInfo -
type AdditionalInfo map[string]interface{}

// CPU -
type CPU struct {
	Count          int             `json:"count,omitempty"`          // nr of physical CPUs
	CoresPerCPU    int             `json:"coresPerCPU,omitempty"`    // nr of Cores per CPU
	ThreadsPerCore int             `json:"threadsPerCore,omitempty"` // nr of Threads per Core
	TotalThreads   int             `json:"totalThreads,omitempty"`   // nr of CPU Threads from all Cores and all CPUs
	Architecture   CPUArchitecture `json:"architecture,omitempty"`   // ex: `amd64` or `ppc64`
	Variant        CPUVariant      `json:"variant,omitempty"`        //
	Manufacturer   string          `json:"manufacturer,omitempty"`   // intel, amd, arm
	ByteOrder      string          `json:"byteOrder,omitempty"`      //
	Features       []string        `json:"features,omitempty"`       // ex: `SWPB (swap) instructions` for ARM
	AdditionalInfo AdditionalInfo  `json:"additionalInfo,omitempty"` //
}

// CPUVariant -
type CPUVariant struct {
	Name     CPUArchitectureVariant         `json:"name,omitempty"`     // ex: ARMv8
	Detailed CPUArchitectureVariantDetailed `json:"detailed,omitempty"` // ex: ARMv8.2-A
}

// CPUArchitecture -
type CPUArchitecture string

const (
	CPU_386         CPUArchitecture = "386"         //
	CPU_AMD64                       = "AMD64"       //
	CPU_AMD64p32                    = "AMD64p32"    //
	CPU_ARM                         = "ARM"         //
	CPU_ARMbe                       = "ARMbe"       //
	CPU_ARM64                       = "ARM64"       //
	CPU_ARM64be                     = "ARM64be"     //
	CPU_PPC64                       = "PPC64"       //
	CPU_PPC64le                     = "PPC64le"     //
	CPU_MIPS                        = "MIPS"        //
	CPU_MIPSle                      = "MIPSle"      //
	CPU_MIPS64                      = "MIPS64"      //
	CPU_MIPS64le                    = "MIPS64le"    //
	CPU_MIPS64p32                   = "MIPS64p32"   //
	CPU_MIPS64p32le                 = "MIPS64p32le" //
	CPU_PPC                         = "PPC"         //
	CPU_RISCv                       = "RISCv"       //
	CPU_RISCv64                     = "RISCv64"     //
	CPU_S390                        = "S390"        //
	CPU_S390x                       = "S390x"       //
	CPU_SPARC                       = "SPARC"       //
	CPU_SPARC64                     = "SPARC64"     //
	CPU_WASM                        = "WASM"        //
	CPU_Uknown                      = "Uknown"      //
)

var supportedCPUArchitectures = []CPUArchitecture{
	CPU_386,
	CPU_AMD64,
	CPU_AMD64p32,
	CPU_ARM,
	CPU_ARMbe,
	CPU_ARM64,
	CPU_ARM64be,
	CPU_PPC64,
	CPU_PPC64le,
	CPU_MIPS,
	CPU_MIPSle,
	CPU_MIPS64,
	CPU_MIPS64le,
	CPU_MIPS64p32,
	CPU_MIPS64p32le,
	CPU_PPC,
	CPU_RISCv,
	CPU_RISCv64,
	CPU_S390,
	CPU_S390x,
	CPU_SPARC,
	CPU_SPARC64,
	CPU_WASM,
}

// CPUArchitectureVariant -
type CPUArchitectureVariant string

// https://en.wikipedia.org/wiki/List_of_ARM_microarchitectures
const (
	CPUV_ARMv1  CPUArchitectureVariant = "ARM1"
	CPUV_ARMv2                         = "ARM2"
	CPUV_ARMv3                         = "ARM3"
	CPUV_ARMv4                         = "ARM4"
	CPUV_ARMv5                         = "ARM5"
	CPUV_ARMv6                         = "ARM6"
	CPUV_ARMv7                         = "ARM7"
	CPUV_ARMv8                         = "ARM8"
	CPUV_Uknown                        = "Uknown"
)

var supportedCPUArchitectureVariants = []CPUArchitectureVariant{
	CPUV_ARMv1,
	CPUV_ARMv2,
	CPUV_ARMv3,
	CPUV_ARMv4,
	CPUV_ARMv5,
	CPUV_ARMv6,
	CPUV_ARMv7,
	CPUV_ARMv8,
}

// CPUArchitectureVariantDetailed -
type CPUArchitectureVariantDetailed string

const (
	CPUVD_ARMv2A    CPUArchitectureVariantDetailed = "ARMv2a"
	CPUVD_ARMv4T                                   = "ARMv4t"
	CPUVD_ARMv5T                                   = "ARMv5t"
	CPUVD_ARMv5TE                                  = "ARMv5te"
	CPUVD_ARMv5TEJ                                 = "ARMv5tej"
	CPUVD_ARMv6_M                                  = "ARMv6-m"
	CPUVD_ARMv6K                                   = "ARMv6k"
	CPUVD_ARMv6T2                                  = "ARMv6t2"
	CPUVD_ARMv6TEJ                                 = "ARMv6tej"
	CPUVD_ARMv6Z                                   = "ARMv6z"
	CPUVD_ARMv7_A                                  = "ARMv7-a"
	CPUVD_ARMv7_M                                  = "ARMv7-m"
	CPUVD_ARMv7_R                                  = "ARMv7-r"
	CPUVD_ARMv7E_M                                 = "ARMv7e-m"
	CPUVD_ARMv8_A                                  = "ARMv8-a"
	CPUVD_ARMv8_M                                  = "ARMv8-m"
	CPUVD_ARMv8_R                                  = "ARMv8-r"
	CPUVD_ARMv8_1_A                                = "ARMv8.1-a"
	CPUVD_ARMv8_2_A                                = "ARMv8.2-a"
	CPUVD_ARMv8_3_A                                = "ARMv8.3-a"
	CPUVD_ARMv8_4_A                                = "ARMv8.4-a"
	CPUVD_Uknown                                   = "Uknown"
)
