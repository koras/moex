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
	"strconv"

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
	// month
	for month := 2; month <= 9; month++ {
		// days
		for day := 1; day <= 31; day++ {

			for p := 0; p <= 4; p++ {

				pageStart := p * 100

				dayString := strconv.Itoa(day)
				if day < 10 {
					dayString = "0" + strconv.Itoa(day)
				}

				monthString := strconv.Itoa(month)
				if month < 10 {
					monthString = "0" + strconv.Itoa(month)
				}

				//	url := fmt.Sprintf("https://iss.moex.com/iss/history/engines/stock/markets/shares/boardgroups/57/securities.jsonp?iss.meta=off&iss.json=extended&callback=JSON_CALLBACK&lang=ru&security_collection=3&date=2022-0%v-%v&start=0&limit=20&sort_column=VALUE&sort_order=desc", i, v)
				url := fmt.Sprintf("https://iss.moex.com/iss/history/engines/stock/markets/shares/boardgroups/57/securities.jsonp?iss.meta=off&iss.json=extended&date=2022-%v-%v&start=%v&limit=100&sort_order=desc", monthString, dayString, pageStart)

				dt := fmt.Sprintf("2022-%v-%v", monthString, dayString)
				fmt.Printf("%s \n\r", url)
				lastDay(db, url, dt)
			}
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
		fmt.Printf("fail %s ", url)
		panic(err)
	}
	content := string(body)

	contentReplace := services.ClearText(content)
	if contentReplace != "" {
		fmt.Print("\r")
		var moex config.LASTVOLUME

		if err := json.Unmarshal([]byte(contentReplace), &moex); err != nil { // Parse []byte to go struct pointer
			fmt.Println("contentReplace")
			fmt.Println(contentReplace)
			fmt.Printf("fail 2 %s ", url)
			panic(err)
		}

		for _, rec := range moex.History {

			//		fmt.Print("\r\n")
			//	fmt.Sprintf(" %v ", rec.Shortname)
			fmt.Sprintf(" %v ", rec.Shortname)
			if rec.Close != 0 {
				product := repositories.Product{
					Name:     rec.Secid,
					Price:    rec.Close,
					Quantity: rec.Value,
					Fullname: rec.Shortname,
					Date:     dt,
				}
				models.Update(db, product)
			}
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
