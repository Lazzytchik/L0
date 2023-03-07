package db

type Storable interface {
	Insert() string
	Delete(id int) string
	TableName() string
}
