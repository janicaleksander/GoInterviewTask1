package main

import "github.com/janicaleksander/GoInterviewTask1/api"

func main() {
	serv := api.NewServer(":8080")
	serv.RunServer()

}
