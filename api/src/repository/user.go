package repository

import (
	"database/sql"
	"fmt"
	"social-car/src/models"
)

// UserRepo represents a user repository
type UserRepo struct {
	db *sql.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

// Create creates a new user
func (repo *UserRepo) Create(user models.User) (uint64, error) {
	stmt, err := repo.db.Prepare("INSERT INTO user (name, nick, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastIDInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastIDInserted), nil
}

// Search searches for users
func (repo *UserRepo) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // returns %nameOrNick%

	lines, err := repo.db.Query(`
		SELECT id, name, nick, email, created_at from user WHERE name LIKE ? OR nick LIKE ?
		`, nameOrNick, nameOrNick)

	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User
		if err = lines.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUserByID returns a user
func (repo *UserRepo) GetUserByID(ID uint64) (models.User, error) {
	line, err := repo.db.Query(`
		SELECT id, name, nick, email, created_at from user WHERE id = ?
		`, ID)
	if err != nil {
		return models.User{}, err
	}
	defer line.Close()

	var user models.User
	if line.Next() {
		if err = line.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

// UpdateUser updates a user
func (repo *UserRepo) UpdateUser(userID uint64, user models.User) error {
	stmt, err := repo.db.Prepare("UPDATE user SET name = ?, nick = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Nick, user.Email, userID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser deletes a user
func (repo *UserRepo) DeleteUser(userID uint64) error {
	stmt, err := repo.db.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(userID); err != nil {
		return err
	}

	return nil
}

// GetUserByEmail returns a user based on given e-mail
func (repo *UserRepo) GetUserByEmail(email string) (models.User, error) {
	line, err := repo.db.Query(`
		SELECT id, password from user WHERE email = ?
		`, email)
	if err != nil {
		return models.User{}, err
	}
	defer line.Close()

	var user models.User
	if line.Next() {
		if err = line.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

// FollowUser inserts a new follower relationship into the database.
func (repo *UserRepo) FollowUser(userID, followerID uint64) error {
	stmt, err := repo.db.Prepare(
		"INSERT OR IGNORE INTO follower (user_id, follower_id) VALUES (?, ?)",
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(userID, followerID); err != nil {
		return err
	}
	return nil
}

// UnFollowUser removes a follower relationship into the database.
func (repo *UserRepo) UnFollowUser(userID, followerID uint64) error {
	stmt, err := repo.db.Prepare(
		"DELETE FROM follower WHERE user_id = ? AND follower_id = ?",
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(userID, followerID); err != nil {
		return err
	}
	return nil
}
