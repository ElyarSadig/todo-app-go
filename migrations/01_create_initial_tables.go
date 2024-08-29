package migrations

import "database/sql"

const (
	userTable = `
	CREATE TABLE IF NOT EXISTS users (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	name VARCHAR(255) NOT NULL,
    	email VARCHAR(255) UNIQUE NOT NULL,
    	token VARCHAR(255),
    	password VARCHAR(255) NOT NULL
	);
`
	todoTable = `
	CREATE TABLE IF NOT EXISTS todos (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	user_id INT,
    	title VARCHAR(255) NOT NULL,
    	completed BOOLEAN DEFAULT FALSE,
    	created_at TIMESTAMP,
    	updated_at TIMESTAMP,
    	FOREIGN KEY (user_id) REFERENCES users(id)
	);
`
)

func RunMigrationsV1(db *sql.DB) error {
	_, err := db.Exec(userTable)
	if err != nil {
		return err
	}
	_, err = db.Exec(todoTable)
	if err != nil {
		return err
	}
	return nil
}
