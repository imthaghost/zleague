package main

import "zleague/api/server"

func main() {
	// start the server
	server := server.New(nil)
	server.Start(":8080")
}
