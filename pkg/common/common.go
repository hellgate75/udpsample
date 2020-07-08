package common

import (
	"fmt"
	"os"
)

const Network = "udp"

const LookupNetwork = ""

func CheckError(err error) {
	if err != nil {
		fmt.Printf("Fatal error %v\n", err.Error())
		os.Exit(1)
	}
}