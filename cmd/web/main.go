package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Shreyash-trivedi/bookings/pkg/config"
	"github.com/Shreyash-trivedi/bookings/pkg/handlers"
	"github.com/Shreyash-trivedi/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const pN = ":8080"
var app config.AppConfig
var session *scs.SessionManager


func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	
	tc, err := render.CreateTemplatecache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	_, _ = fmt.Printf("starting app on %s", pN)
	// http.ListenAndServe(pN, nil)
	srv := &http.Server{
		Addr:    pN,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
