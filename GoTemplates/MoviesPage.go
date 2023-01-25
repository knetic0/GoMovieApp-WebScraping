package GoTemplates

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gomovie/Webscraping"
	"html/template"
	"net/http"
)

type Movies struct {
	Titles []string
	Stars  []string
}

func (movies Movies) MoviesPage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles("./Templates/Movies/movies.html", "./Templates/Head/head.html", "./Templates/navbar.html")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var content_str []string
	var stars_str []string

	content_str, stars_str = Webscraping.Scrap()

	datas := Movies{Titles: content_str, Stars: stars_str}
	view.ExecuteTemplate(w, "movies", datas)
}
