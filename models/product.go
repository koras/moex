package models

import (
	"database/sql"
	"log"
	"moex/repositories"

	_ "github.com/go-sql-driver/mysql"
)

type ProductDb struct {
	id       int64
	name     string
	fullname string
	price    float64
	quantity float64
	date     string
}

type HistoryDate struct {
	date string
}

// https://learningprogramming.net/golang/golang-and-mysql/update-entity-in-golang-and-mysql-database/

func GetHistory(db *sql.DB) []repositories.Product {
	result := []repositories.Product{}
	rows, err := db.Query("SELECT `id`, `name`, `fullname`, `price`,`quantity`,`date`  from `product` ORDER BY `name`, `date` ASC ;")
	if err != nil {
		return nil
	}

	log.Print("GetHistory")

	var fullname string = ""

	var percent int64 = 100
	var oldQuantity float64 = 0.0

	for rows.Next() {

		//	var prd repositories.Product
		var prd ProductDb
		err = rows.Scan(&prd.id, &prd.name, &prd.fullname, &prd.price, &prd.quantity, &prd.date)

		log.Print("prd.date== ", prd.date, " string == ", prd.fullname)

		if fullname != prd.fullname {
			percent = 100
		} else {
			percent = int64(prd.quantity) * 100 / int64(oldQuantity)
		}
		oldQuantity = prd.quantity

		product := repositories.Product{prd.id, prd.name, prd.fullname, prd.price, prd.quantity, prd.date, percent}
		result = append(result, product)
	}
	return result

}

/**
* Получаем дату от лейблов
 */
func GetDate(db *sql.DB) []repositories.HistoryDate {

	result := []repositories.HistoryDate{}
	rows, err := db.Query("SELECT `date`  from `product`  GROUP BY  `date`  ORDER BY  `date` ASC;")
	if err != nil {
		return nil
	}

	for rows.Next() {
		var dt HistoryDate
		err = rows.Scan(&dt.date)
		product := repositories.HistoryDate{dt.date}
		result = append(result, product)

	}
	return result

}

/**
* Обновляем данные от парсера которые приходят
 */
func Update(db *sql.DB, product repositories.Product) (int64, error) {
	result, err := db.Exec("INSERT INTO product (`name`,`price`,`quantity`,`date`,`fullname` ) VALUES (?, ?, ?, ?, ?);", product.Name, product.Price, product.Quantity, product.Date, product.Fullname)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
	return 0, err
}
