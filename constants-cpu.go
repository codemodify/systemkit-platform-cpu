package platform

// CPUArchitecture -
type CPUArchitecture string

const (
	CPU_386         CPUArchitecture = "386"         //
	CPU_AMD64                       = "amd64"       //
	CPU_AMD64p32                    = "amd64p32"    //
	CPU_ARM                         = "arm"         //
	CPU_ARMbe                       = "armbe"       //
	CPU_ARM64                       = "arm64"       //
	CPU_ARM64be                     = "arm64be"     //
	CPU_PPC64                       = "ppc64"       //
	CPU_PPC64le                     = "ppc64le"     //
	CPU_MIPS                        = "mips"        //
	CPU_MIPSle                      = "mipsle"      //
	CPU_MIPS64                      = "mips64"      //
	CPU_MIPS64le                    = "mips64le"    //
	CPU_MIPS64p32                   = "mips64p32"   //
	CPU_MIPS64p32le                 = "mips64p32le" //
	CPU_PPC                         = "ppc"         //
	CPU_RISCv                       = "riscv"       //
	CPU_RISCv64                     = "riscv64"     //
	CPU_S390                        = "s390"        //
	CPU_S390x                       = "s390x"       //
	CPU_SPARC                       = "sparc"       //
	CPU_SPARC64                     = "sparc64"     //
	CPU_WASM                        = "wasm"        //
	CPU_Uknown                      = "uknown"
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
	CPUV_ARMv1  CPUArchitectureVariant = "armv1"
	CPUV_ARMv2                         = "armv2"
	CPUV_ARMv3                         = "armv3"
	CPUV_ARMv4                         = "armv4"
	CPUV_ARMv5                         = "armv5"
	CPUV_ARMv6                         = "armv6"
	CPUV_ARMv7                         = "armv7"
	CPUV_ARMv8                         = "armv8"
	CPUV_Uknown                        = "uknown"
)

// CPUArchitectureVariantDetailed -
type CPUArchitectureVariantDetailed string

const (
	CPUVD_ARMv2A    CPUArchitectureVariantDetailed = "armv2a"
	CPUVD_ARMv4T                                   = "armv4t"
	CPUVD_ARMv5T                                   = "armv5t"
	CPUVD_ARMv5TE                                  = "armv5te"
	CPUVD_ARMv5TEJ                                 = "armv5tej"
	CPUVD_ARMv6_M                                  = "armv6-m"
	CPUVD_ARMv6K                                   = "armv6k"
	CPUVD_ARMv6T2                                  = "armv6t2"
	CPUVD_ARMv6TEJ                                 = "armv6tej"
	CPUVD_ARMv6Z                                   = "armv6z"
	CPUVD_ARMv7_A                                  = "armv7-a"
	CPUVD_ARMv7_M                                  = "armv7-m"
	CPUVD_ARMv7_R                                  = "armv7-r"
	CPUVD_ARMv7E_M                                 = "armv7e-m"
	CPUVD_ARMv8_A                                  = "armv8-a"
	CPUVD_ARMv8_M                                  = "armv8-m"
	CPUVD_ARMv8_R                                  = "armv8-r"
	CPUVD_ARMv8_1_A                                = "armv8.1-a"
	CPUVD_ARMv8_2_A                                = "armv8.2-a"
	CPUVD_ARMv8_3_A                                = "armv8.3-a"
	CPUVD_ARMv8_4_A                                = "armv8.4-a"
	CPUVD_Uknown                                   = "uknown"
)
