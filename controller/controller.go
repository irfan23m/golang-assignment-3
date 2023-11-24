package controller

import (
	"assignment-3/config"
	"assignment-3/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddDataEnvironment(ctx echo.Context) error {
	db := config.GetDB()

	environment := models.Environment{}

	if err := ctx.Bind(&environment); err != nil {
		return err
	}

	db.Create(&environment)

	// resString, _ := json.Marshal(environment)
	// fmt.Println(string(resString))

	return ctx.JSON(http.StatusOK, environment)
}
