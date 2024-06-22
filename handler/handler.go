package handler

import (
	"fmt"
	"net/http"

	"github.com/gioadorno/flatrock/models"
)

func RenderPage(p models.Page) error {
  pageNumber := p.c.Param("page")
  return p.c.Render(http.StatusOK, fmt.Sprintf("page%s", pageNumber), nil)
}
