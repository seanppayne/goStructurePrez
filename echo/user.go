package echo

import (
	"example.com/demo"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	userRepo demo.UserRepository
}

func (s *echoServer) RegisterUserHandlers() {
	userHandler := UserHandler{
		userRepo: s.UserRepo,
	}

	s.server.GET("/getUser", userHandler.getUser)
	s.server.POST("/addUser", userHandler.addUser)
}

func (u *UserHandler) getUser(c echo.Context) error {
	id := c.QueryParam("id")

	user, err := u.userRepo.Get(c.Request().Context(), id)

	if err != nil {
		c.String(http.StatusInternalServerError, "error getting user")
		return err
	}

	err = c.JSON(200, user)

	if err != nil {
		c.String(http.StatusInternalServerError, "error parsing user")
		return err
	}

	return nil
}

func (u *UserHandler) addUser(c echo.Context) error {
	user := &demo.User{}

	err := c.Bind(user)

	if err != nil {
		c.String(http.StatusInternalServerError, "error binding user")
		return err
	}

	err = u.userRepo.Add(c.Request().Context(), user)

	if err != nil {
		c.String(http.StatusInternalServerError, "error adding user")
		return err
	}

	err = c.JSON(200, user)

	if err != nil {
		c.String(http.StatusInternalServerError, "error parsing user")
		return err
	}

	return nil
}
