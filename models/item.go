package models

import (
	"errors"
	"fmt"
)

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

func (d Item) Validate(err []error) bool {
	switch {
	case d.ChrtId == 0:
		err = append(err, errors.New(fmt.Sprintf("Invalid ChrtID given: %d", d.ChrtId)))
		return false
	case d.TrackNumber == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid TrackNumber given: %s", d.TrackNumber)))
		return false
	case d.Price == 0:
		err = append(err, errors.New(fmt.Sprintf("Invalid Price given: %f", d.Price)))
		return false
	case d.RID == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid RID given: %s", d.RID)))
		return false
	case d.Name == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid Name signature given: %s", d.Name)))
		return false
	case d.Sale == 0:
		err = append(err, errors.New(fmt.Sprintf("Invalid Sale given: %f", d.Sale)))
		return false
	case d.Size == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid Size given: %s", d.Size)))
		return false
	case d.TotalPrice == 0:
		err = append(err, errors.New(fmt.Sprintf("Invalid TotalPrice given: %f", d.TotalPrice)))
		return false
	case d.NMID == 0:
		err = append(err, errors.New(fmt.Sprintf("Invalid NMID given: %d", d.NMID)))
		return false
	case d.Brand == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid Brand given: %s", d.Brand)))
		return false
	case d.Status == 0:
		err = append(err, errors.New(fmt.Sprintf("Invalid Status given: %d", d.Status)))
		return false
	}

	return true
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
