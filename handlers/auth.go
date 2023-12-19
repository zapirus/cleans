package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"

	"clean/pkg/types"
)

type RequestUser struct {
	Guid       uuid.UUID
	Login      string
	Password   string
	Name       string
	Email      string
	VerifyCode string
	CreatedAt  string
	UpdatedAt  string
}

func (a *Api) login(c echo.Context) error {
	user := new(RequestUser)
	if err := c.Bind(user); err != nil {
		return err
	}
	val, err := a.app.Login(c.Request().Context(), user.Login, user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  statusError,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response{
		Status: statusSuccess,
		Body:   val,
	})
}

func (a *Api) register(c echo.Context) error {
	user := new(RequestUser)
	if err := c.Bind(user); err != nil {
		return err
	}

	fmt.Println("user", user)
	regUser, err := a.app.Register(c.Request().Context(), *user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  statusError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response{
		Status: statusSuccess,
		Body:   regUser,
	})
}

func (a *Api) reset(c echo.Context) error {
	user := new(types.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	resetUser, err := a.app.Reset(c.Request().Context(), user.Login)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  statusError,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response{
		Status: statusSuccess,
		Body:   "Мы отправили вам код на почту: " + *resetUser,
	})
}

func (a *Api) verify(c echo.Context) error {
	user := new(types.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	verifyUser, err := a.app.Verify(c.Request().Context(), user.Login, user.VerifyCode)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  statusError,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response{
		Status: statusSuccess,
		Body:   verifyUser,
	})

}

func (a *Api) resend(c echo.Context) error {
	user := new(RequestUser)
	if err := c.Bind(user); err != nil {
		return err
	}

	resendCode, err := a.app.Resend(c.Request().Context(), user.Login)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  statusError,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response{
		Status: statusSuccess,
		Body:   resendCode,
	})

}
