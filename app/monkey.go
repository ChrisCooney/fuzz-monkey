package main

import (
	"fmt"
)

func main() {
	config := GetConfigFromCli()
	result := RunHttpSpam(&config.Endpoints[0])
	fmt.Printf("%+v\n",result)
}
