package main

import (
	"ZiranServer/route"
)

// var a int

func main() {

	r := route.Route()

	//3.监听端口，默认8080
	r.Run(":8080")
}
