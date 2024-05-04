package usecases

import (
	"context"
	"nutri-plans-api/dto"
	"nutri-plans-api/entities"
	"nutri-plans-api/repositories"
	passutil "nutri-plans-api/utils/password"

	"github.com/labstack/echo/v4"
)

type UserUsecase interface {
	Register(c echo.Context, r *dto.RegisterRequest) error
}

type userUsecase struct {
	userRepo    repositories.UserRepository
	authRepo    repositories.AuthRepository
	countryRepo repositories.CountryRepository
	passUtil    passutil.PasswordUtil
}

func NewUserUsecase(
	userRepo repositories.UserRepository,
	authRepo repositories.AuthRepository,
	countryRepo repositories.CountryRepository,
	passUtil passutil.PasswordUtil,
) *userUsecase {
	return &userUsecase{
		userRepo:    userRepo,
		authRepo:    authRepo,
		countryRepo: countryRepo,
		passUtil:    passUtil,
	}
}

func (u *userUsecase) Register(c echo.Context, r *dto.RegisterRequest) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	country, err := u.countryRepo.GetCountryByID(ctx, r.CountryID)
	if err != nil {
		return err
	}

	hashedPassword, err := u.passUtil.HashPassword(r.Password)
	if err != nil {
		return err
	}

	auth := &entities.Auth{
		Email:    r.Email,
		Password: hashedPassword,
		Username: r.Username,
	}
	if err := u.authRepo.CreateAuth(ctx, auth); err != nil {
		return err
	}

	user := &entities.User{
		AuthID:    auth.ID,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Dob:       r.Dob,
		Gender:    r.Gender,
		Country:   *country,
	}
	return u.userRepo.CreateUser(ctx, user)
}
