package main

import (
	"prakerja2/config"
	"prakerja2/route"
)

func main() {
	config.InitDB()
	e := route.InitRoute()
	e.Logger.Fatal(e.Start(":8080"))
}

// fundamental golang
// variable, kondisi, perulangan, slice, struct

// rest api
// response code, response body, design rest api, method, database, CRUD Rest API
// jwt, oauth, unit test, clean architecture, microservices, kafka, handler, message queue

// deployment
// free cloud
// docker, container, AWS, GCP
