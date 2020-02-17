# Detailed Platform Detection at Runtime
[![GoDoc](https://godoc.org/github.com/codemodify/SystemKit?status.svg)](https://godoc.org/github.com/codemodify/systemkit-platform)
[![0-License](https://img.shields.io/badge/license-0--license-brightgreen)](https://github.com/codemodify/TheFreeLicense)
[![Go Report Card](https://goreportcard.com/badge/github.com/codemodify/SystemKit)](https://goreportcard.com/report/github.com/codemodify/SystemKit)
[![Test Status](https://github.com/danawoodman/systemservice/workflows/Test/badge.svg)](https://github.com/danawoodman/systemservice/actions)
![code size](https://img.shields.io/github/languages/code-size/codemodify/SystemKit?style=flat-square)

# Usages
- Intented use is as a library in Go code
	```go
	packlage main

	import platform "github.com/codemodify/systemkit-platform"

	func main() {

		fmt.Println(platform.GetInfo().CPU.Architecture)     // => ex: arm OR amd64
		fmt.Println(platform.GetInfo().CPU.Variant.Name)     // => ex: armv5 OR armv6 OR armv8, etc
		fmt.Println(platform.GetInfo().CPU.Variant.Detailed) // => ex: armv5te

	}
	```

-

# References
- `go tool dist list`
- https://github.com/golang/go/blob/master/src/go/build/syslist.go
- https://github.com/containerd/containerd
- https://github.com/mackerelio/go-osstat/blob/master/cpu/cpu_linux.go
