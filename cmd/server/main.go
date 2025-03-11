package main

import "github.com/philippeberto/go-api/configs"

func main() {
	config := configs.LoadConfig()
	println("running at ", config.WebServerPort)
}
