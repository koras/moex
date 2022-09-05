package main

import (
	"fmt"
	"log"
	"moex/services"
	"net/http"
	"os"

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

	port := os.Getenv("PORT_SERVER")

	http.HandleFunc("/", services.Home)

	http.HandleFunc("/chart", services.GetCharts)
	http.HandleFunc("/getInfo", services.GetData)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("ListendAndServe doesn't work : ", err)
	}
}
