package repositories

type Product struct {
	//	x        string
	Id       int64
	Name     string
	Fullname string
	Price    float64
	Quantity float64
	Date     string
	Percent  int64
}

type HistoryDate struct {
	Date string
}
