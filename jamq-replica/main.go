package main

import (
	"jamq-replica/router"
)

func main() {
	e := router.New()
	v1 = e.Group("/api/v1")
	e.Start(":5000")
}
