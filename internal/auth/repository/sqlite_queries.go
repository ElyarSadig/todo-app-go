package repository

const (
	createUserQuery     = `INSERT INTO users(name, email, token, password) VALUES (?, ?, ?, ?) RETURNING id`
	updateUserToken     = `UPDATE users SET token = ? WHERE email = ?`
	getUserByEmailQuery = `SELECT id, name, email, token, password FROM users WHERE email = ?`
	deleteUserToken     = `UPDATE users SET token = null WHERE token = ?`
)
