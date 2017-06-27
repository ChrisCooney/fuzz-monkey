package main

import (
	"fmt"
	"fuzz-monkey/app/cli"
)

func main() {
	config := cli.GetConfig()
	fmt.Printf("%+v\n",config)
}
