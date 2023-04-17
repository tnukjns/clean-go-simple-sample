package main

import "fmt"

// User struct
type User struct {
	ID int64
}

// UserGateway struct
type UserGateway struct{}

func (g *UserGateway) FindByID(id int64) (*User, error) {
	return &User{ID: id}, nil
}

// UserUsecase struct
type UserUsecase struct {
	userGateway *UserGateway
}

// NewUserUsecase creates a new UserUsecase without dependency injection
func NewUserUsecase(userGateway *UserGateway) *UserUsecase {
	return &UserUsecase{userGateway: userGateway}
}

// GetUser gets a user by ID
func (u *UserUsecase) GetUser(id int64) (*User, error) {
	return u.userGateway.FindByID(id)
}

func main() {
	userGateway := &UserGateway{}
	userUsecase := NewUserUsecase(userGateway)

	user, err := userUsecase.GetUser(1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Actual User ID:", user.ID)
}
