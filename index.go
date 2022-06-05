package main

import (
	"fmt"
	"log"
	"moex/services"
	"net/http"

	"github.com/joho/godotenv"
)

type Todo struct {
	Title string
	Done  bool
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load("onee.env"); err != nil {
		log.Print("No .env file found. my bug")
	}
}
func main() {
	log.Printf("start")
	// route
	http.HandleFunc("/", services.Home)

	err := http.ListenAndServe(":9990", nil)
	if err != nil {
		fmt.Println("ListendAndServe doesn't work : ", err)
	}
}
