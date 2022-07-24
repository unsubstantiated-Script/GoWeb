package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

//Setting up CSRF
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// Loading up the session information
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

//func WriteToConsole(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		fmt.Println("Hit the Page")
//		next.ServeHTTP(w, r)
//	})
//}
