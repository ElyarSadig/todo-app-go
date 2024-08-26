package database

import "time"

type User struct {
	ID       int  `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Token    string `db:"token"`
	Password string `db:"password"`
}

type Todo struct {
	ID        int     `db:"id"`
	Title     string    `db:"title"`
	Completed bool      `db:"completed"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

const user string = `
CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255),
    token VARCHAR(255),
    password VARCHAR(255)
);
`

const todo string = `
CREATE TABLE IF NOT EXISTS todos (
    id BIGINT PRIMARY KEY,
    title VARCHAR(255),
    completed BOOLEAN,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
`
