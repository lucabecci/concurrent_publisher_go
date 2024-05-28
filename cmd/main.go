package main

import "github.com/lucabecci/concurrent_publisher_go/internal/domain/entities"

func main() {
	operation := entities.NewOperation("1")
	println(operation.GetStatus())
	operation.Start()
	println(operation.GetStatus())
	operation.Complete()
	println(operation.GetStatus())
	operation.Block()
	println(operation.GetStatus())
}
