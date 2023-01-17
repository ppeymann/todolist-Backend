package main

import (
	"fmt"
	"log"
	"net/http"
	"todoList/router"
)

func main() {
	r := router.Router()
	fmt.Println("Server is Getting Start ...")
	log.Fatal(http.ListenAndServe(":4000",r))
	fmt.Println("Listening at port 4000 ...")
}