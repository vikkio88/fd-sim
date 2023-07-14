package main

import (
	"fdsim/scripts"
	"time"
)

func main() {
	s := scripts.DBGen{}

	s.Run(time.Now().Unix(), 20)
}
