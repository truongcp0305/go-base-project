package main

import app "go-project/init"

// @title Base Go API
// @version 1.0
// @host localhost:1234
// @description This is a sample server for Base go API.
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	app.StartApp()
}
