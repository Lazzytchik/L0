package models

import "fmt"

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

func (d Item) Insert() string {
	return fmt.Sprintf(
		"INSERT INTO %s (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES ('%d', '%s', '%f', '%s', '%s', '%f', '%s', '%f', '%d', '%s', '%d') RETURNING id",
		d.TableName(),
		d.ChrtId,
		d.TrackNumber,
		d.Price,
		d.RID,
		d.Name,
		d.Sale,
		d.Size,
		d.TotalPrice,
		d.NMID,
		d.Brand,
		d.Status,
	)
}

func (d Item) Delete(id int) string {
	return fmt.Sprintf(
		"DELETE FROM %s WHERE %s = %d",
		d.TableName(),
		d.PrimaryColumn(),
		id,
	)
}

func (d Item) TableName() string {
	return "items"
}

func (d Item) PrimaryColumn() string {
	return "id"
}
