package main

import (
	"fmt"

	"github.com/matkinhig/go-topup/server"
	_ "github.com/mattn/go-oci8"
)

func main() {
	fmt.Println("Start golang...")
	server.Run()
}
