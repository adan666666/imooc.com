package main

import (
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println(grpc.Version)
}
