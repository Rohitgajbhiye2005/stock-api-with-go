package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rohitgajbhiye2005/stock-api/router"
)

func main() {
	r := router.Router()
	fmt.Print("Starting the server at 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
