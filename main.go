package main

import (
	"github.com/shikhar1996/Covid19/src/server"
)

func main() {

	data := server.Getdata()
	server.Updatedata(data)
}
