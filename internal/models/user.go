package models

type User struct {
	ID       int     `db:"id"`
	Name     string  `db:"name"`
	Email    string  `db:"email"`
	Token    *string `db:"token"`
	Password string  `db:"password"`
}
