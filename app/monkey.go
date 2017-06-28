package main

import (
	"fmt"
)

func main() {
	config := GetConfigFromCli()
	fmt.Printf("%+v\n",config)
	result := RunCorrupt(&config.Endpoints[0])
	fmt.Printf("%+v\n",result)
}
