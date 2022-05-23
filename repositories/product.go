package repositories

type Product struct {
	Id       int64
	Name     string
	Fullname string
	Price    float64
	Quantity float64
	Date     string
}

type HistoryDate struct {
	Date string
}
