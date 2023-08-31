package repository

import (
	"database/sql"
	"go-demo-app/models"
	"log"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (username, email) VALUES ($1, $2)"

	_, err := r.DB.Exec(query, user.Username, user.Email)
	if err != nil {
		log.Println("Error creating user:", err)
		return err
	}

	return nil
}

func (r *UserRepository) GetUserByID(userID int) (*models.User, error) {
	// Implement database select logic here
	return nil, nil
}

func (r *UserRepository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	rows, err := r.DB.Query("SELECT id, username, email FROM users")
	if err != nil {
		log.Println("Error retrieving users:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			log.Println("Error scanning user:", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		return nil, err
	}

	return users, nil
}
