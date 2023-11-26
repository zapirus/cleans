package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	"clean/pkg/types"
)

func (a *Api) login(c echo.Context) error {
	login := c.FormValue("login")
	password := c.FormValue("password")
	val, err := a.app.Login(c.Request().Context(), login, password)
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

	user := &types.User{
		Guid:       c.FormValue("guid"),
		Login:      c.FormValue("login"),
		Password:   c.FormValue("password"),
		Name:       c.FormValue("name"),
		Email:      c.FormValue("email"),
		VerifyCode: c.FormValue("verify"),
	}
	regUser, err := a.app.Register(c.Request().Context(), user)
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

func (a *Api) verify(c echo.Context) error {
	guid := c.FormValue("guid")
	verify := c.FormValue("verify")
	err := a.app.Verify(c.Request().Context(), guid, verify)
	if err != nil {
		return c.JSON(http.StatusNotFound, response{
			Status:  statusError,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response{
		Status:  statusSuccess,
		Message: "",
	})

}

func (a *Api) reset(c echo.Context) error {
	login := c.FormValue("login")
	password := c.FormValue("password")
	retryPassword := c.FormValue("retryPassword")
	if err := a.app.Reset(c.Request().Context(), login, password, retryPassword); err != nil {
		c.JSON(http.StatusNotFound, response{
			Status:  statusError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response{
		Status:  statusSuccess,
		Message: "Пароль сброшен",
	})
}

func (a *Api) resend(c echo.Context) error {
	login := c.FormValue("login")
	password := c.FormValue("password")
	if err := a.app.Resend(c.Request().Context(), login, password); err != nil {
		return c.JSON(http.StatusBadRequest, response{Status: statusError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, response{Status: statusSuccess, Body: "Пароль успешно отправлен на почту"})

}
