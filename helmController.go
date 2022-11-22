package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func addHelmController() {
	// Routes
	e.GET("/public/helm", getReleases)
}

func getReleases(c echo.Context) error {
	doGetReleases()
	return c.JSON(http.StatusOK, &Result{Result: 1})
}

// Handler
func getPodsCount(c echo.Context) error {
	return c.JSON(http.StatusOK, &Result{Result: countPods()})
  }