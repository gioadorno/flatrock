package models

import "github.com/labstack/echo/v4"

type Page struct {
  c echo.Context
  page int
}

