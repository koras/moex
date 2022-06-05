package main

import (
	"log"
	"moex/services"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

func init() {
	// loads values from .env into the system
	//if err := godotenv.Load(); err != nil {
	//	log.Print("No .env file found. my bug")
	//}
}
func main() {
	log.Printf("start")
	// route
	http.HandleFunc("/", services.Home)
	http.HandleFunc("/chart", services.GetCharts)

	http.ListenAndServe(":80", nil)
}
