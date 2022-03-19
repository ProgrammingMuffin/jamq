package main

import (
	"jamq-replica/router"
	v1 "jamq-replica/v1"
)

func main() {
	e := router.New()
	v1Group := e.Group("/api/v1")
	v1.HandleV1Routes(v1Group)
	e.Start(":5000")
}
