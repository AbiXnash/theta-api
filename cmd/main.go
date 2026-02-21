package main

import "github.com/AbiXnash/theta-api/internals/router"

func main() {
	r := router.GetRouter()
	r.Run()
}
