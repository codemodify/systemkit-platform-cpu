package main

import (
	"encoding/json"
	"fmt"
	"os"

	platformCPU "github.com/codemodify/systemkit-platform-cpu"
)

func main() {
	cpuInfo := platformCPU.Info()

	data, err := json.Marshal(cpuInfo)
	if err != nil {
		fmt.Println(err.Error())
	}

	if len(os.Args) > 1 && os.Args[1] == "-p" {
		data, err = json.MarshalIndent(cpuInfo, "", "    ")
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Println(string(data))
}
