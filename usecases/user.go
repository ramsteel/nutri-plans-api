package usecases

import (
	"context"
	"nutri-plans-api/dto"
	"nutri-plans-api/entities"
	"nutri-plans-api/repositories"
	passutil "nutri-plans-api/utils/password"
	tokenutil "nutri-plans-api/utils/token"

	"github.com/labstack/echo/v4"
)

type UserUsecase interface {
	Register(c echo.Context, r *dto.RegisterRequest) error
	Login(c echo.Context, r *dto.LoginRequest) (*dto.LoginResponse, error)
}

type userUsecase struct {
	userRepo    repositories.UserRepository
	authRepo    repositories.AuthRepository
	countryRepo repositories.CountryRepository

	passUtil  passutil.PasswordUtil
	tokenUtil tokenutil.TokenUtil
}

func NewUserUsecase(
	userRepo repositories.UserRepository,
	authRepo repositories.AuthRepository,
	countryRepo repositories.CountryRepository,
	passUtil passutil.PasswordUtil,
	tokenUtil tokenutil.TokenUtil,
) *userUsecase {
	return &userUsecase{
		userRepo:    userRepo,
		authRepo:    authRepo,
		countryRepo: countryRepo,
		passUtil:    passUtil,
		tokenUtil:   tokenUtil,
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
	if err = u.authRepo.CreateAuth(ctx, auth); err != nil {
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

func (u *userUsecase) Login(c echo.Context, r *dto.LoginRequest) (*dto.LoginResponse, error) {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	auth, err := u.authRepo.GetAuthByEmail(ctx, r.Email)
	if err != nil {
		return nil, err
	}

	if err = u.passUtil.VerifyPassword(r.Password, auth.Password); err != nil {
		return nil, err
	}

	token, err := u.tokenUtil.GenerateToken(auth.ID, auth.RoleTypeID)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{Token: token}, nil
}
