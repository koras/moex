package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"moex/config"
	"moex/models"
	"moex/repositories"
	"moex/services"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load("onee.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	db, err := config.GetMySQLDB()
	if err != nil {
		panic(err)
	}

	for v := 1; v < 32; v++ {
		for i := 5; i <= 9; i++ {
			//	url := fmt.Sprintf("https://iss.moex.com/iss/history/engines/stock/markets/shares/boardgroups/57/securities.jsonp?iss.meta=off&iss.json=extended&callback=JSON_CALLBACK&lang=ru&security_collection=3&date=2022-0%v-%v&start=0&limit=20&sort_column=VALUE&sort_order=desc", i, v)
			url := fmt.Sprintf("https://iss.moex.com/iss/history/engines/stock/markets/shares/boardgroups/57/securities.jsonp?iss.meta=off&iss.json=extended&date=2022-0%v-%v&start=0&limit=100&sort_order=desc", i, v)

			dt := fmt.Sprintf("2022-0%v-%v", i, v)
			//	fmt.Printf("%s \n\r", url)
			lastDay(db, url, dt)
		}
	}
	fmt.Println("done")
}

/**
* Парсер за каждй день
 */
func lastDay(db *sql.DB, url string, dt string) {

	res, err := http.Get(url)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	content := string(body)

	contentReplace := services.ClearText(content)

	fmt.Println(contentReplace)
	fmt.Print("\r\n")
	var moex config.LASTVOLUME

	if err := json.Unmarshal([]byte(contentReplace), &moex); err != nil { // Parse []byte to go struct pointer
		panic(err)
	}

	for _, rec := range moex.History {
		if rec.Close != 0 {
			product := repositories.Product{
				Name:     rec.Secid,
				Price:    rec.Close,
				Quantity: rec.Marketprice3Tradesvalue,
				Fullname: rec.Shortname,
				Date:     dt,
			}
			models.Update(db, product)
		}
	}
}

/**
* онлайн парсер
 */
func curentDay() {
	// url := "https://iss.moex.com/iss/engines/stock/markets/shares/boardgroups/57/securities.jsonp?iss.meta=off&iss.json=extended&callback=JSON_CALLBACK&lang=ru&security_collection=160&sort_column=VALTODAY&sort_order=desc"

	// db, err := modul.GetMySQLDB()

	// productModel := models.ProductModel{
	// 	Db: db,
	// }

	// res, err := http.Get(url)
	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// content := string(body)

	// contentReplace := modul.ClearText(content)

	// fmt.Println(contentReplace)
	// var moex modul.MOEXCURRENT

	// if err := json.Unmarshal([]byte(contentReplace), &moex); err != nil { // Parse []byte to go struct pointer
	// 	panic(err)
	// }

	// for _, rec := range moex.Marketdata {

	// 	product := modul.Product{
	// 		Name:     rec.Secid,
	// 		Price:    rec.Last,
	// 		Quantity: rec.ValtodayRur,
	// 		Fullname: rec.Boardid,
	// 		Date:     "2022-12-13",
	// 	}
	// 	rowsAffected, err2 := productModel.Update(product)
	// 	if err2 != nil {
	// 		fmt.Println(err2)
	// 	} else {
	// 		fmt.Println("Rows Affected:", rowsAffected)
	// 	}
	// }
}

// models
