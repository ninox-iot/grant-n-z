package main

import (
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/domain/repository"
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/domain"
	"gopkg.in/go-playground/validator.v9"
	"github.com/tomoyane/grant-n-z/controller"
	"github.com/tomoyane/grant-n-z/infra"
)

func main() {
	infra.InitDB()
	infra.DbMigration()

	di.InitUserService(repository.UserRepositoryImpl{})
	di.InitTokenService(repository.TokenRepositoryImpl{})

	e := echo.New()
	e.Validator = &domain.GrantValidator{Validator: validator.New()}
	e.POST("/v1/users", controller.PostUser)
	e.POST("/v1/tokens", controller.PostToken)
	e.Logger.Fatal(e.Start(":8080"))
}