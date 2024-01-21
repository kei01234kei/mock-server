package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORSのミドルウェアを全許可の設定で追加
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	// すべてのルートに対してハンドラを登録
	e.Any("/*", func(c echo.Context) error {
		return c.String(http.StatusOK, "status code OK: 200")
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
