/*
 * @Author: kingford
 * @Date: 2022-09-20 11:49:21
 * @LastEditTime: 2022-09-20 11:56:58
 */
package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/skingford/gin-server/internal/dto"
	"github.com/skingford/gin-server/internal/entity"
	"github.com/skingford/gin-server/internal/repository"
)

// UserService is a contract.....
type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}
