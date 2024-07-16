package main

import "github.com/andrew-nino/ATM/internal/app"

const pathConfig = "config/config.yaml"

func main() {
	app.Run(pathConfig)
}
