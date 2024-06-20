package main

import (
	"github.com/gioadorno/flatrock/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	p1 := &Page{c: e.Context(), page: 1}
	p2 := &Page{c: e.Context(), page: 2}
	p3 := &Page{c: e.Context(), page: 3}

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "home", nil)
	})

	e.GET("/page-1", func(c echo.Context) error { return p1.RenderPage() })
	e.GET("/page-2", func(c echo.Context) error { return p2.Render() })
	e.GET("/page-3", func(c echo.Context) error { return p3.Render() })

	// hx-get and hx-push-state are used to enable htmx routing
	e.GET("/page-{page}", func(c echo.Context) error {
		page := c.Param("page")
		switch page {
		case "1":
			return p1.Render()
		case "2":
			return p2.Render()
		case "3":
			return p3.Render()
		default:
			return c.String(404, "Page not found")
		}
	})

	e.GET("/page-{page}/prev", func(c echo.Context) error {
		page := c.Param("page")
		switch page {
		case "1":
			return p3.Render()
		default:
			return p2.Render()
		}
	})

	e.GET("/page-{page}/next", func(c echo.Context) error {
		page := c.Param("page")
		switch page {
		case "3":
			return p1.Render()
		default:
			return p2.Render()
		}
	})

	e.Start(":8080")
}
