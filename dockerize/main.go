package main

import (
	"fmt"
	"os"
)

func main() {
	env, ok := os.LookupEnv("NAME")

	if !ok {
		fmt.Println(`ENV Variable not found`)
		os.Exit(1)
	}

	fmt.Println(env, "hello, this is working..")
}
