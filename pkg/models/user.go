package models

type User struct {
	ID            int
	Name          string
	DateCreated   string
	Illusions     int
	XCards        int
	Notifications bool
	RedeemedCodes []string
}
