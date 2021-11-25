package main

import (
	"fmt"
	logger "mygogw/log"
	"mygogw/vnic"
)

func main() {
	logger.InitLogger()
	err := vnic.CreateVrf("cjhVrf1", 174)
	if err != nil {
		fmt.Printf("CreateVrf err:%v\n", err)
	}

	err = vnic.CreateVlan("cjhVlan1", "123.123.123.2/24", "ens3")
	if err != nil {
		fmt.Printf("CreateVlan err:%v\n", err)
	}

	fmt.Printf("main exit")
}
