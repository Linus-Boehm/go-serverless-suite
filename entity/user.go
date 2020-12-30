package entity

type User struct {
	ID            ID
	Email         string
	Firstname     string
	Lastname      string
	EmailVerified bool
	Attributes    map[string]string
	Timestamps
}
