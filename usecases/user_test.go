package usecases_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"nutri-plans-api/dto"
	"nutri-plans-api/entities"
	mockrepo "nutri-plans-api/mocks/repositories"
	mockpass "nutri-plans-api/mocks/utils/password"
	mocktoken "nutri-plans-api/mocks/utils/token"
	"nutri-plans-api/usecases"
)

type testCase struct {
	name string
	errs []error
}

func TestNewUserUsecase(t *testing.T) {
	assert.NotNil(
		t,
		usecases.NewUserUsecase(
			mockrepo.NewMockUserRepository(t),
			mockrepo.NewMockAuthRepository(t),
			mockrepo.NewMockUserPreferenceRepository(t),
			mockrepo.NewMockCountryRepository(t),
			mockpass.NewMockPasswordUtil(t),
			mocktoken.NewMockTokenUtil(t),
		),
	)
}

func TestRegister(t *testing.T) {
	var (
		countryID = uint(1)
		country   = &entities.Country{
			ID:   uint(1),
			Name: "Indonesia",
		}

		registerRequest = &dto.RegisterRequest{
			Email:     "some@example.com",
			Password:  "password",
			Username:  "username",
			FirstName: "first",
			LastName:  "last",
			Dob:       time.UnixMilli(1714757476909),
			Gender:    "M",
			CountryID: countryID,
		}

		auth = &entities.Auth{
			Email:    registerRequest.Email,
			Password: "hashedpasswordinhere",
			Username: registerRequest.Username,
		}

		user = &entities.User{
			AuthID:    uuid.MustParse("00000000-0000-0000-0000-000000000000"),
			FirstName: "first",
			LastName:  "last",
			Dob:       time.UnixMilli(1714757476909),
			Gender:    "M",
			Country:   *country,
		}
	)

	testCases := []testCase{
		{
			name: "success",
			errs: []error{nil, nil, nil, nil},
		},
		{
			name: "error get country",
			errs: []error{errors.New("country not found"), nil, nil, nil},
		},
		{
			name: "error hash password",
			errs: []error{nil, errors.New("failed hashing password"), nil, nil},
		},
		{
			name: "error create auth",
			errs: []error{nil, nil, errors.New("failed to create auth"), nil},
		},
		{
			name: "error create user",
			errs: []error{nil, nil, nil, errors.New("failed to create user")},
		},
	}

	for idx, tc := range testCases {
		mockUserRepo := new(mockrepo.MockUserRepository)
		mockAuthRepo := new(mockrepo.MockAuthRepository)
		mockUserPreferenceRepo := new(mockrepo.MockUserPreferenceRepository)
		mockCountryRepo := new(mockrepo.MockCountryRepository)
		mockPassUtil := new(mockpass.MockPasswordUtil)
		mockTokenUtil := new(mocktoken.MockTokenUtil)
		u := usecases.NewUserUsecase(
			mockUserRepo,
			mockAuthRepo,
			mockUserPreferenceRepo,
			mockCountryRepo,
			mockPassUtil,
			mockTokenUtil,
		)
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/register", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			ctx, cancel := context.WithCancel(c.Request().Context())
			defer cancel()
			mockCountryRepo.On("GetCountryByID", ctx, countryID).Return(country, tc.errs[0])
			mockPassUtil.On("HashPassword", registerRequest.Password).Return(
				auth.Password,
				tc.errs[1],
			)
			mockAuthRepo.On("CreateAuth", ctx, auth).Return(tc.errs[2])
			mockUserRepo.On("CreateUser", ctx, user).Return(tc.errs[3])
			mockUserPreferenceRepo.On(
				"CreateUserPreference",
				ctx,
				&entities.UserPreference{UserID: user.AuthID},
			).Return(nil)
			err := u.Register(c, registerRequest)

			if idx != 0 {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestLogin(t *testing.T) {
	var (
		loginRequest = &dto.LoginRequest{
			Email:    "some@example.com",
			Password: "password",
		}

		auth = &entities.Auth{
			Email:      loginRequest.Email,
			Password:   "hashedpasswordinhere",
			Username:   "testuser",
			RoleTypeID: uint(1),
			CreatedAt:  time.UnixMilli(1714757476909),
			UpdatedAt:  time.UnixMilli(1714757476909),
			DeletedAt:  gorm.DeletedAt{},
		}

		token = "jwttoken"

		loginResponse = &dto.LoginResponse{
			Token: token,
		}
	)

	testCases := []testCase{
		{
			name: "success",
			errs: []error{nil, nil, nil},
		},
		{
			name: "error get auth",
			errs: []error{errors.New("failed to get auth"), nil, nil},
		},
		{
			name: "error verify password",
			errs: []error{nil, errors.New("failed to verify password"), nil},
		},
		{
			name: "error generate token",
			errs: []error{nil, nil, errors.New("failed to generate token")},
		},
	}

	for idx, tc := range testCases {
		mockUserRepo := new(mockrepo.MockUserRepository)
		mockAuthRepo := new(mockrepo.MockAuthRepository)
		mockUserPreferenceRepo := new(mockrepo.MockUserPreferenceRepository)
		mockCountryRepo := new(mockrepo.MockCountryRepository)
		mockPassUtil := new(mockpass.MockPasswordUtil)
		mockTokenUtil := new(mocktoken.MockTokenUtil)
		u := usecases.NewUserUsecase(
			mockUserRepo,
			mockAuthRepo,
			mockUserPreferenceRepo,
			mockCountryRepo,
			mockPassUtil,
			mockTokenUtil,
		)
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/login", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			ctx, cancel := context.WithCancel(c.Request().Context())
			defer cancel()
			mockAuthRepo.On("GetAuthByEmail", ctx, loginRequest.Email).Return(auth, tc.errs[0])
			mockPassUtil.On("VerifyPassword", loginRequest.Password, auth.Password).Return(
				tc.errs[1],
			)
			mockTokenUtil.On("GenerateToken", auth.ID, auth.RoleTypeID).Return(
				loginResponse.Token,
				tc.errs[2],
			)

			_, err := u.Login(c, loginRequest)
			if idx != 0 {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, token, loginResponse.Token)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	var (
		id = uuid.New()
		r  = &dto.UpdateUserRequest{
			Email:     "some@example.com",
			Username:  "testuser",
			FirstName: "test",
			LastName:  "user",
			Dob:       time.UnixMilli(1714757476909),
			Gender:    "M",
			CountryID: 1,
		}
		user = &entities.User{
			AuthID: id,
			Auth: entities.Auth{
				ID:       id,
				Email:    r.Email,
				Username: r.Username,
			},
			FirstName: r.FirstName,
			LastName:  r.LastName,
			Dob:       r.Dob,
			Gender:    r.Gender,
			CountryID: r.CountryID,
		}
	)

	testCases := []testCase{
		{
			name: "success",
			errs: []error{nil},
		},
		{
			name: "error update",
			errs: []error{errors.New("failed to update")},
		},
	}

	for idx, tc := range testCases {
		mockUserRepo := new(mockrepo.MockUserRepository)
		mockAuthRepo := new(mockrepo.MockAuthRepository)
		mockUserPreferenceRepo := new(mockrepo.MockUserPreferenceRepository)
		mockCountryRepo := new(mockrepo.MockCountryRepository)
		mockPassUtil := new(mockpass.MockPasswordUtil)
		mockTokenUtil := new(mocktoken.MockTokenUtil)
		u := usecases.NewUserUsecase(
			mockUserRepo,
			mockAuthRepo,
			mockUserPreferenceRepo,
			mockCountryRepo,
			mockPassUtil,
			mockTokenUtil,
		)
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "/profiles", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			ctx, cancel := context.WithCancel(c.Request().Context())
			defer cancel()
			mockUserRepo.On("UpdateUser", ctx, user).Return(tc.errs[0])
			err := u.UpdateUser(c, id, r)
			if idx != 0 {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetUserByID(t *testing.T) {
	var (
		id      = uuid.New()
		example = &entities.User{AuthID: id}
	)
	mockUserRepo := new(mockrepo.MockUserRepository)
	mockAuthRepo := new(mockrepo.MockAuthRepository)
	mockUserPreferenceRepo := new(mockrepo.MockUserPreferenceRepository)
	mockCountryRepo := new(mockrepo.MockCountryRepository)
	mockPassUtil := new(mockpass.MockPasswordUtil)
	mockTokenUtil := new(mocktoken.MockTokenUtil)
	u := usecases.NewUserUsecase(
		mockUserRepo,
		mockAuthRepo,
		mockUserPreferenceRepo,
		mockCountryRepo,
		mockPassUtil,
		mockTokenUtil,
	)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/profiles", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()
	mockUserRepo.On("GetUserByID", ctx, id).Return(example, nil)
	res, err := u.GetUserByID(c, id)
	assert.NoError(t, err)
	assert.Equal(t, res, example)
}
