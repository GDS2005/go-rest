package repositories

import (
	"database/sql"
	"example/models"
)

type UserRepository struct {
	db *sql.DB
}

// NewUserRepository initializes a new UserRepository instance
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetAll retrieves all users from the database
func (r *UserRepository) GetAll() ([]models.User, error) {
	// Execute SQL query to retrieve all users
	rows, err := r.db.Query("SELECT id, username, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the rows and scan the data into User structs
	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetByID retrieves a user by ID from the database
func (r *UserRepository) GetByID(id int) (*models.User, error) {
	// Execute SQL query to retrieve the user with the given ID
	row := r.db.QueryRow("SELECT id, username, email FROM users WHERE id = $1", id)

	// Scan the row data into a User struct
	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Create inserts a new user into the database
func (r *UserRepository) Create(user *models.User) error {
	// Execute SQL query to insert the new user
	_, err := r.db.Exec("INSERT INTO users (username, email) VALUES ($1, $2)", user.Username, user.Email)
	return err
}

// Update updates an existing user in the database
func (r *UserRepository) Update(id int, user *models.User) error {
	// Execute SQL query to update the user with the given ID
	_, err := r.db.Exec("UPDATE users SET username = $1, email = $2 WHERE id = $3", user.Username, user.Email, id)
	return err
}

// Delete deletes a user from the database
func (r *UserRepository) Delete(id int) error {
	// Execute SQL query to delete the user with the given ID
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
