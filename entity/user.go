package entity

type User struct {
	ID         ID
	Email      string
	Firstname  string
	Lastname   string
	Attributes map[string]string
	Timestamps
}
