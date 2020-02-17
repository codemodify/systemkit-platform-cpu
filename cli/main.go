package main

import (
	"encoding/json"
	"fmt"
	"os"

	platform "github.com/codemodify/systemkit-platform"
)

func main() {
	data, err := json.Marshal(platform.GetInfo())
	if err != nil {
		fmt.Println(err.Error())
	}

	if len(os.Args) > 1 && os.Args[1] == "-p" {
		data, err = json.MarshalIndent(platform.GetInfo(), "", "    ")
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Println(string(data))
}
