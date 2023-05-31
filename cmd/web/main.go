package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/simonwardle/bookings/pkg/config"
	"github.com/simonwardle/bookings/pkg/handlers"
	"github.com/simonwardle/bookings/pkg/render"
)

const portNumber = ":8080"

// declare here so it is available to entire main package
var app config.AppConfig
var session *scs.SessionManager

// main is the main application starting point
func main() {
	var err error
	//*********  Change this to true when going into Production **************
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	app.UseTemplateCache = false

	//populate app config template cache via render package
	app.TemplateCache, err = render.CreateTemplateCache()
	if err != nil {
		log.Fatal("error: cannot load template cache!")
	}

	//pass app config populated with template cache to render.
	render.NewTemplates(&app)

	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)

	_, _ = fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
