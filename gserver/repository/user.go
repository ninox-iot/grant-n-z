package repository

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type UserRepository interface {
	FindById(id int) (*entity.User, *model.ErrorResBody)

	FindByEmail(email string) (*entity.User, *model.ErrorResBody)

	FindUserWithRoleByEmail(email string) (*model.UserOperatorPolicy, *model.ErrorResBody)

	Save(user entity.User) (*entity.User, *model.ErrorResBody)

	SaveUserWithUserService(user entity.User, userService *entity.UserService) (*entity.User, *model.ErrorResBody)

	Update(user entity.User) (*entity.User, *model.ErrorResBody)
}
