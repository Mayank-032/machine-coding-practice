package stock

type Stock struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	CurrentPrice float64 `json:"current_price"`
}

func New() *Stock {
	return &Stock{}
}

func (s *Stock) UpdateCurrentPrice(newPrice float64) {
	s.CurrentPrice = newPrice
}
