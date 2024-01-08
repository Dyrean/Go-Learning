package main

import (
	"event-booking/internal/server"
	"fmt"
	"os"
	"strconv"
)

func main() {
	server := server.New()

	server.RegisterRoutes()

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}
