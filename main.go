package main

import (
	"WonderServer/route"
	"log/slog"
	"time"
)

// var a int

func main() {

	r := route.Route()
	//3.监听端口，默认8080
	go func() {
		time.Sleep(5 * time.Second)
		slog.Info("server started")
		slog.Info("\033]8;;http://localhost:8080\033\\Ctr + Click Here to Open the Link\033]8;;\033\\")
	}()
	r.Run(":8080")
}
