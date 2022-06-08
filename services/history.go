package services

import (
	"fmt"
	"html/template"
	"moex/config"
	"moex/models"
	"moex/repositories"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type ViewData struct {
	Title    string
	Labels   []string
	Products Charts
}

type Charts struct {
	Label       []string
	BorderColor []string
	DataMaps    map[string][]repositories.Product
}

func GetCharts(w http.ResponseWriter, r *http.Request) {
	// подключаем базу, рак, но потом переделаю
	db, err := config.GetMySQLDB()
	if err != nil {
		panic(err)
	}

	// подключаем шаблоны
	tmpl := template.Must(template.ParseFiles("html/layout.html", "html/chart.html"))
	// получаем данные из базы
	history := models.GetHistory(db)
	historyDate := models.GetDate(db)

	var data ViewData
	var name string
	var color string
	var labels []string
	var labelName []string
	var labelNameDate []string
	var chart Charts

	chart.DataMaps = make(map[string][]repositories.Product)
	for indx := 0; indx < len(historyDate); indx++ {
		fmt.Sprintln("dd== %v", historyDate[indx].Date)
		labelNameDate = append(labelNameDate, historyDate[indx].Date)
	}

	for i := 0; i < len(history); i++ {
		if name != history[i].Fullname {
			name = history[i].Fullname
			color = "#3e95cd"
			labels = nil
			//  устанавливаем данные
			labelName = append(labelName, history[i].Date)
			// потом доделаю, пока заглушка
			chart.BorderColor = append(chart.BorderColor, color)
			chart.Label = append(chart.Label, name)
		}
		// устанавливаем данные
		chart.DataMaps[name] = append(chart.DataMaps[name], history[i])
		labels = append(labels, history[i].Date)
	}

	data = ViewData{
		Title:    "Market volume",
		Labels:   labelNameDate,
		Products: chart,
	}

	tmpl.Execute(w, data)
}
