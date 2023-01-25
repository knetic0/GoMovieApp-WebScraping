package Routes

import (
	"github.com/julienschmidt/httprouter"
	"gomovie/GoTemplates"
)

func Router() *httprouter.Router {
	app := httprouter.New()

	app.GET("/", GoTemplates.Dashboard{}.Index)
	app.GET("/movies", GoTemplates.Movies{}.MoviesPage)
	app.GET("/search", GoTemplates.Search{}.SearchMovie)

	return app
}
