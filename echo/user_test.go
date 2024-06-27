package echo

import (
	"context"
	"errors"
	"example.com/demo"
	"example.com/demo/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type UserTestSuite struct {
	suite.Suite
	echo *echo.Echo
}

func TestUserTestSuite(t *testing.T) {
	t.Log("Creating the UserTestSuite")
	suite.Run(t, new(UserTestSuite))
}

func (suite *UserTestSuite) SetupTest() {
	suite.T().Log("Running the CardHandleTestSuite setup")
	suite.echo = echo.New()
}

func (suite *UserTestSuite) TestGetUser() {
	suite.T().Log("Running the TestGetUser")

	req := httptest.NewRequest(http.MethodGet, "/v1/getUser?id=1", nil)

	rec := httptest.NewRecorder()

	c := suite.echo.NewContext(req, rec)

	UserRepository := mocks.UserRepository{
		GetUserFunc: func(ctx context.Context, ID string) (*demo.User, error) {
			if ID != "1" {
				return nil, errors.New("error getting user")
			}

			return &demo.User{
				ID:       1,
				Name:     "John Doe",
				Email:    "john_doe@gmail.com",
				Birthday: demo.CustomTime(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
			}, nil
		},
	}

	userHandler := UserHandler{
		userRepo: &UserRepository,
	}

	err := userHandler.getUser(c)

	suite.NoError(err)

	suite.Equal(http.StatusOK, rec.Code)

	suite.Equal(`{"ID":1,"Name":"John Doe","Email":"john_doe@gmail.com","Birthday":"1990-01-01"}`, strings.TrimSuffix(rec.Body.String(), "\n"))
}
