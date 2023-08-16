package main

import (
	"fmt"
	"go-mongo/routes"
	"net/http"
)

func main() {
	fmt.Println("server running on port 3000 ...")
	r := routes.Router()
	http.ListenAndServe(":3000", r)
}
