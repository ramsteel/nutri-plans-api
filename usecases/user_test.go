package usecases_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"nutri-plans-api/dto"
	"nutri-plans-api/entities"
	mockrepo "nutri-plans-api/mocks/repositories"
	mockpass "nutri-plans-api/mocks/utils/password"
	"nutri-plans-api/usecases"
)

type testCase struct {
	name string
	errs []error
}

var testCases = []testCase{
	{
		name: "success",
		errs: []error{nil, nil, nil},
	},
	{
		name: "error get country",
		errs: []error{errors.New("country not found"), nil, nil},
	},
	{
		name: "error hash password",
		errs: []error{nil, errors.New("failed hashing password"), nil},
	},
	{
		name: "error create auth",
		errs: []error{nil, nil, errors.New("failed to create auth")},
	},
}

func TestNewUserUsecase(t *testing.T) {
	assert.NotNil(
		t,
		usecases.NewUserUsecase(
			mockrepo.NewMockUserRepository(t),
			mockrepo.NewMockAuthRepository(t),
			mockrepo.NewMockCountryRepository(t),
			mockpass.NewMockPasswordUtil(t),
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

	for idx, tc := range testCases {
		mockUserRepo := new(mockrepo.MockUserRepository)
		mockAuthRepo := new(mockrepo.MockAuthRepository)
		mockCountryRepo := new(mockrepo.MockCountryRepository)
		mockPassUtil := new(mockpass.MockPasswordUtil)
		u := usecases.NewUserUsecase(mockUserRepo, mockAuthRepo, mockCountryRepo, mockPassUtil)
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
			mockUserRepo.On("CreateUser", ctx, user).Return(nil)
			err := u.Register(c, registerRequest)
			fmt.Println(tc.errs[1])

			if idx != 0 {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
