package main

import (
	"net/http"
	"fmt"
	"github.com/labstack/echo/v4"
)

func addHelmController() {
	// Routes
	e.GET("/public/helm", getPodsCount)
}

// Handler
func getPodsCount(c echo.Context) error {
	return c.JSON(http.StatusOK, &Result{Result: countPods()})
  }