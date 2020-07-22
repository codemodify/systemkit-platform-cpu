# ![](https://fonts.gstatic.com/s/i/materialiconsoutlined/flare/v4/24px.svg) Detailed CPU Detection at Runtime
[![](https://img.shields.io/github/v/release/codemodify/systemkit-platform-cpu?style=flat-square)](https://github.com/codemodify/systemkit-platform-cpu/releases/latest)
![](https://img.shields.io/github/languages/code-size/codemodify/systemkit-platform-cpu?style=flat-square)
![](https://img.shields.io/github/last-commit/codemodify/systemkit-platform-cpu?style=flat-square)
[![](https://img.shields.io/badge/license-0--license-brightgreen?style=flat-square)](https://github.com/codemodify/TheFreeLicense)

![](https://img.shields.io/github/workflow/status/codemodify/systemkit-platform-cpu/qa?style=flat-square)
![](https://img.shields.io/github/issues/codemodify/systemkit-platform-cpu?style=flat-square)
[![](https://goreportcard.com/badge/github.com/codemodify/systemkit-platform-cpu?style=flat-square)](https://goreportcard.com/report/github.com/codemodify/systemkit-platform-cpu)

[![](https://img.shields.io/badge/godoc-reference-brightgreen?style=flat-square)](https://godoc.org/github.com/codemodify/systemkit-platform-cpu)
![](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)
![](https://img.shields.io/gitter/room/codemodify/systemkit-platform-cpu?style=flat-square)

![](https://img.shields.io/github/contributors/codemodify/systemkit-platform-cpu?style=flat-square)
![](https://img.shields.io/github/stars/codemodify/systemkit-platform-cpu?style=flat-square)
![](https://img.shields.io/github/watchers/codemodify/systemkit-platform-cpu?style=flat-square)
![](https://img.shields.io/github/forks/codemodify/systemkit-platform-cpu?style=flat-square)


# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Usage
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
