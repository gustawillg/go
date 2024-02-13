package main

import "github.com/gustawillg/go/9/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
