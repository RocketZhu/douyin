package main

import (
	"douyin/router"
	"log"
)

func main() {
	r := router.Init()
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
