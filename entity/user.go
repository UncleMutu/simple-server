package entity

type User struct {
	ID    int64
	Name  string
	Email string
}

type Location struct {
	ID       int64
	Address  string
	Postcode string
	Lat      float64
	Lng      float64
}
