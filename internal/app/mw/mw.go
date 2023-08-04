package mw

import (
	"github.com/labstack/echo/v4"
	"log"
)

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		val := ctx.Request().Header.Get("User-Role")

		if val == "admin" {
			log.Println("You are logged as admin.")
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}
