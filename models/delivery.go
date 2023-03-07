package models

import "fmt"

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

func (d Delivery) Insert() string {
	return fmt.Sprintf(
		"INSERT INTO %s (name, phone, zip, city, address, region, email) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s') RETURNING id",
		d.TableName(),
		d.Name,
		d.Phone,
		d.Zip,
		d.City,
		d.Address,
		d.Region,
		d.Email,
	)
}

func (d Delivery) Delete(id int) string {
	return fmt.Sprintf(
		"DELETE FROM %s WHERE %s = %d",
		d.TableName(),
		d.PrimaryColumn(),
		id,
	)
}

func (d Delivery) TableName() string {
	return "deliveries"
}

func (d Delivery) PrimaryColumn() string {
	return "id"
}
