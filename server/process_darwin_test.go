package server

import (
	"fmt"
	"testing"
)

func TestDarwinProce(t *testing.T) {
	list, err := processes()

	if err != nil {

	}

	for _, itm := range list {
		fmt.Printf("%v \t %v \t %s \n", itm.Pid(), itm.PPid(), itm.Executable())
	}
}
