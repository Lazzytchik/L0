package models

import (
	"errors"
	"fmt"
)

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

func (d Delivery) Validate(err []error) bool {
	switch {
	case d.Name == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid Name given: %s", d.Name)))
		return false
	case d.Phone == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid Phone given: %s", d.Phone)))
		return false
	case d.Zip == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid Zip given: %s", d.Zip)))
		return false
	case d.City == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid City given: %s", d.City)))
		return false
	case d.Address == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid Address signature given: %s", d.Address)))
		return false
	case d.Region == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid Region given: %s", d.Region)))
		return false
	case d.Email == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid Email given: %s", d.Email)))
		return false
	}

	return true
}
