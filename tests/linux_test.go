// +build linux
// +build !android

package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	platformCPU "github.com/codemodify/systemkit-platform-cpu"
)

func TestCPUVariant(t *testing.T) {
	data, err := json.Marshal(platformCPU.Info())
	if err != nil {
		fmt.Println(err.Error())
		t.Fatal(err.Error())
	}

	dataAsString := string(data)
	fmt.Println(dataAsString)
}
