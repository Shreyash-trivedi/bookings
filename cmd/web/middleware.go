package main

import (
	"net/http"
	"github.com/justinas/nosurf"
)

//NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler{
	csrfHanddler := nosurf.New(next)

	csrfHanddler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHanddler 
}

//SessionLoad loads and saves the session on every event
func SessionLoad(next http.Handler)http.Handler{
	return session.LoadAndSave(next)
}