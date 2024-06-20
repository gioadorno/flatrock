package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Page struct {
  c echo.Context
  page int
}

func RenderPage(p Page) error {
  pageNumber := p.c.Param("page")
  return p.c.Render(http.StatusOK, fmt.Sprintf("page%s", pageNumber), nil)
}
