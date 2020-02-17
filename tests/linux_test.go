// +build linux
// +build !android

package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	platform "github.com/codemodify/systemkit-platform"
)

func TestCPUVariant(t *testing.T) {
	data, err := json.Marshal(platform.GetInfo())
	if err != nil {
		fmt.Println(err.Error())
		t.Fatal(err.Error())
	}

	dataAsString := string(data)
	fmt.Println(dataAsString)
}
