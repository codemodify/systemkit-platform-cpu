// +build linux
// +build !android

package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	platform "github.com/codemodify/systemkit-platform-cpu"
)

func TestJSON(t *testing.T) {
	var testValue = `{
   "lscpu": [
      {"field":"Architecture:", "data":"x86_64"},
      {"field":"CPU op-mode(s):", "data":"32-bit, 64-bit"},
      {"field":"Byte Order:", "data":"Little Endian"},
      {"field":"Address sizes:", "data":"39 bits physical, 48 bits virtual"},
      {"field":"CPU(s):", "data":"12"},
      {"field":"On-line CPU(s) list:", "data":"0-11"},
      {"field":"Thread(s) per core:", "data":"2"},
      {"field":"Core(s) per socket:", "data":"6"},
      {"field":"Socket(s):", "data":"1"},
      {"field":"NUMA node(s):", "data":"1"},
      {"field":"Vendor ID:", "data":"GenuineIntel"},
      {"field":"CPU family:", "data":"6"},
      {"field":"Model:", "data":"158"},
      {"field":"Model name:", "data":"Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz"},
      {"field":"Stepping:", "data":"10"},
      {"field":"CPU MHz:", "data":"3999.895"},
      {"field":"CPU max MHz:", "data":"4300.0000"},
      {"field":"CPU min MHz:", "data":"800.0000"},
      {"field":"BogoMIPS:", "data":"5202.65"},
      {"field":"Virtualization:", "data":"VT-x"},
      {"field":"L1d cache:", "data":"192 KiB"},
      {"field":"L1i cache:", "data":"192 KiB"},
      {"field":"L2 cache:", "data":"1.5 MiB"},
      {"field":"L3 cache:", "data":"9 MiB"},
      {"field":"NUMA node0 CPU(s):", "data":"0-11"},
      {"field":"Vulnerability Itlb multihit:", "data":"KVM: Mitigation: Split huge pages"},
      {"field":"Vulnerability L1tf:", "data":"Mitigation; PTE Inversion; VMX conditional cache flushes, SMT vulnerable"},
      {"field":"Vulnerability Mds:", "data":"Mitigation; Clear CPU buffers; SMT vulnerable"},
      {"field":"Vulnerability Meltdown:", "data":"Mitigation; PTI"},
      {"field":"Vulnerability Spec store bypass:", "data":"Mitigation; Speculative Store Bypass disabled via prctl and seccomp"},
      {"field":"Vulnerability Spectre v1:", "data":"Mitigation; usercopy/swapgs barriers and __user pointer sanitization"},
      {"field":"Vulnerability Spectre v2:", "data":"Mitigation; Full generic retpoline, IBPB conditional, IBRS_FW, STIBP conditional, RSB filling"},
      {"field":"Vulnerability Tsx async abort:", "data":"Mitigation; Clear CPU buffers; SMT vulnerable"},
      {"field":"Flags:", "data":"fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx pdpe1gb rdtscp lm constant_tsc art arch_perfmon pebs bts rep_good nopl xtopology nonstop_tsc cpuid aperfmperf pni pclmulqdq dtes64 monitor ds_cpl vmx smx est tm2 ssse3 sdbg fma cx16 xtpr pdcm pcid sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand lahf_lm abm 3dnowprefetch cpuid_fault epb invpcid_single pti ssbd ibrs ibpb stibp tpr_shadow vnmi flexpriority ept vpid ept_ad fsgsbase tsc_adjust bmi1 hle avx2 smep bmi2 erms invpcid rtm mpx rdseed adx smap clflushopt intel_pt xsaveopt xsavec xgetbv1 xsaves dtherm ida arat pln pts hwp hwp_notify hwp_act_window hwp_epp md_clear flush_l1d"}
   ]
}`

	lscpuO := platform.LscpuOutput{}
	err := json.Unmarshal([]byte(testValue), &lscpuO)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(lscpuO)
}
