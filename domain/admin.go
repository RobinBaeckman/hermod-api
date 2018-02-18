package domain

// Product is a product that is being sold
type Admin struct {
	ID        string
	Email     string
	Password  []byte
	FirstName string
	LastName  string
}
