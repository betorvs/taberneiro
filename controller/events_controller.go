package controller

import (
	"encoding/json"
	"net/http"

	"github.com/betorvs/taberneiro/usecase"
	"github.com/labstack/echo"
	"github.com/nlopes/slack"
)

// ReceiveEvents func
func ReceiveEvents(c echo.Context) (err error) {
	data := new(slack.InteractionCallback)
	payload := c.FormValue("payload")
	json.Unmarshal([]byte(payload), &data)
	// log.Println(data)
	res, err := usecase.ActionEvent(data, c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, res)
}
