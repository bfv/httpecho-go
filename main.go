package main

import (
	"log"
	"net/http"

	"github.com/bfv/httpecho-go/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	setupServers()
}

func setupServers() {

	e := echo.New()

	s := http.Server{
		Addr:    ":1323",
		Handler: e,
	}

	e.GET("/get", handlers.GetHandler)
	e.POST("/post", handlers.PostHandler)

	e.GET("/quit", func(c echo.Context) error {
		defer s.Close()
		return c.String(http.StatusOK, "done")
	})

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
