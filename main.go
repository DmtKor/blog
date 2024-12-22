package main

import (
	"blog/src/commander"
	"blog/src/middleware"
	"net/http"
	"testing"
)

func TestCommander(t *testing.T) {
	c := commander.Commander{}
	c.InitServer(initstr_temp, ":50050", []func(http.Handler)http.Handler { middleware.LoggerMiddleware, middleware.RedirectMiddleware })
}