package models

type Item struct {
	ChrtId      int     `json:"chrt_id"`
	TrackNumber string  `json:"track_number"`
	Price       float32 `json:"price"`
	RID         string  `json:"rid"`
	Name        string  `json:"name"`
	Sale        float32 `json:"sale"`
	Size        string  `json:"size"`
	TotalPrice  float32 `json:"total_price"`
	NMID        int     `json:"nm_id"`
	Brand       string  `json:"brand"`
	Status      int     `json:"status"`
}
