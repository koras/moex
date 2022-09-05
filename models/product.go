package models

import (
	"database/sql"
	"fmt"
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

	sql := "SELECT `id`, `name`, `fullname`, `price`,`quantity`,`date`   from `product`  where  `date` between '2022-02-01' AND '2022-09-05' and fullname = 'Газпрнефть'  ORDER BY `name`, `date` ASC ;"
	fmt.Println("sql: ", sql)
	rows, err := db.Query(sql)
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

		//log.Print("prd.date== ", prd.date, " string == ", prd.fullname)
		//	x := prd.date
		//	log.Print("===pprd.id== ", prd.id, "===prd.prd.date== ", prd.date, " name == ", prd.fullname, " quantity== ", float64(prd.quantity))
		if fullname != prd.fullname {
			percent = 100
		} else {

			if oldQuantity != 0 {
				percent = int64(prd.quantity) * 100 / int64(oldQuantity)
			}
			if percent > 200 {
				log.Print("prd.date== ", percent, " quantity== ", int64(prd.quantity), " string == ", oldQuantity, " name == ", prd.fullname)
			}
		}
		//	y := percent

		// {{range $index, $element := .Products.DataMaps}} dataInfo.push([{{range $element}}
		//	{ x: "{{.Date}}", date:new Date({{.Date}}),  fullname:{{.Fullname}}, price:{{.Price}},y: "{{.Percent}}" }, {{end}} ,]);
		// {{end}}

		fullname = prd.fullname

		oldQuantity = prd.quantity
		//	log.Print("prd.date== ", percent, " string == ", prd.fullname)
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
	sql := "SELECT `date`  from `product` where  `date` between '2022-02-01' AND '2022-09-05'  GROUP BY  `date`  ORDER BY  `date` ASC;"

	rows, err := db.Query(sql)
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
