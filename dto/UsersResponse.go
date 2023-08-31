package dto

import "go-demo-app/models"

type UsersResponse struct {
	Data []*models.User `json:"data"`
}
