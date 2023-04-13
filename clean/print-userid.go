package main

import "fmt"

// User struct
type User struct {
	ID int64
}

// UserRepository interface
type UserRepository interface {
	FindByID(id int64) (*User, error)
}

// UserGateway struct (implements UserRepository)
type UserGateway struct{}

func (g *UserGateway) FindByID(id int64) (*User, error) {
	return &User{ID: id}, nil
}

// UserMockGateway struct (mock implementation of UserRepository)
type UserMockGateway struct{}

func (m *UserMockGateway) FindByID(id int64) (*User, error) {
	return &User{ID: 2}, nil
}

// UserUsecase struct
type UserUsecase struct {
	userRepository UserRepository
}

// NewUserUsecase creates a new UserUsecase with dependency injection
func NewUserUsecase(userRepository UserRepository) *UserUsecase {
	return &UserUsecase{userRepository: userRepository}
}

// GetUser gets a user by ID
func (u *UserUsecase) GetUser(id int64) (*User, error) {
	return u.userRepository.FindByID(id)
}

func main() {
	userGateway := &UserGateway{}
	userUsecase := NewUserUsecase(userGateway)

	user, err := userUsecase.GetUser(1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Actual User ID:", user.ID)

	userMockGateway := &UserMockGateway{}
	userUsecaseWithMock := NewUserUsecase(userMockGateway)

	mockUser, err := userUsecaseWithMock.GetUser(1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Mock User ID:", mockUser.ID)
}
