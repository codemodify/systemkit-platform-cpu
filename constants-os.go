package platform

// OSName -
type OSName string

const (
	OS_AIX       OSName = "aix"       // https://en.wikipedia.org/wiki/IBM_AIX
	OS_Android          = "android"   // https://en.wikipedia.org/wiki/Android_app
	OS_Darwin           = "darwin"    // https://en.wikipedia.org/wiki/Darwin_(operating_system)
	OS_Dragonfly        = "dragonfly" // https://www.dragonflybsd.org/
	OS_FreeBSD          = "freebsd"   // https://www.freebsd.org/
	OS_HURD             = "hurd"      // https://www.gnu.org/software/hurd
	OS_Illumos          = "illumos"   // https://www.illumos.org
	OS_JS               = "js"        // https://github.com/os-js
	OS_Linux            = "linux"     // https://www.kernel.org
	OS_NACL             = "nacl"      // https://en.wikipedia.org/wiki/Google_Native_Client
	OS_NetBSD           = "netbsd"    // https://netbsd.org
	OS_OpenBSD          = "openbsd"   // https://www.openbsd.org
	OS_Plan9            = "plan9"     // https://9p.io/plan9
	OS_Solaris          = "solaris"   // https://www.oracle.com/solaris/solaris11
	OS_Windows          = "windows"   // https://www.microsoft.com/en-us/windows
	OS_ZOS              = "zos"       // https://en.wikipedia.org/wiki/Z/OS, https://www.ibm.com/it-infrastructure/z/zos
	OS_Uknown           = "uknown"
)

var supportedOses = []OSName{
	OS_AIX,
	OS_Android,
	OS_Darwin,
	OS_Dragonfly,
	OS_FreeBSD,
	OS_HURD,
	OS_Illumos,
	OS_JS,
	OS_Linux,
	OS_NACL,
	OS_NetBSD,
	OS_OpenBSD,
	OS_Plan9,
	OS_Solaris,
	OS_Windows,
	OS_ZOS,
}
