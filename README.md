# Detailed CPU Detection at Runtime

# Usage
- as a library in Go code
	```go
	package main

	import platformCPU "github.com/codemodify/systemkit-platform-cpu"

	func main() {
		cpuInfo := platformCPU.Info()

		// Example 1
		fmt.Println(cpuInfo.Architecture)     // => ex: arm OR amd64
		fmt.Println(cpuInfo.Variant.Name)     // => ex: armv5 OR armv6 OR armv8, etc
		fmt.Println(cpuInfo.Variant.Detailed) // => ex: armv5te

		// Example 2
		if platformCPU.IsArm(cpuInfo.Architecture) {
			if cpuInfo.Variant.Name == platformCPU.CPUV_ARMv8 {
				// we got AMRv8
			} else if cpuInfo.Variant.Name == platformCPU.CPUV_ARMv5 {
				if cpuInfo.Variant.Detailed == platformCPU.CPUVD_ARMv5T {
					// we got ARMv5t
				} else if cpuInfo.Variant.Detailed == platformCPU.CPUVD_ARMv5TE {
					// we got ARMv5te
				} else if cpuInfo.Variant.Detailed == platformCPU.CPUVD_ARMv5TEJ {
					// we got ARMv5tej
				}
			}
		}
	}
	```

- as a binary on a a bunch of platforms
	- `https://github.com/codemodify/systemkit-platform-cpu/releases/latest`
	![](https://raw.githubusercontent.com/codemodify/systemkit-platform-cpu/master/.helper-files/dox/sample.png)
