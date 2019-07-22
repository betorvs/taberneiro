package controller

import (
	"net/http"

	"github.com/betorvs/taberneiro/domain"
	"github.com/betorvs/taberneiro/usecase"
	"github.com/labstack/echo"
)

// ReceiveMessages func
func ReceiveMessages(c echo.Context) (err error) {
	message := new(domain.Message)
	if err = c.Bind(message); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	magicHeader := c.Request().Header.Get("X-Magic-Header")
	verifier := usecase.ValidateHeader(magicHeader)
	if verifier != true {
		return c.JSON(http.StatusForbidden, nil)
	}
	res, err := usecase.PostMessage(message)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, res)
}
