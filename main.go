package main

import "api-gemini-go/application/routers"

func main() {
	routers.SetupRouter().Run(":8080")
}
