package GoTemplates

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gomovie/Webscraping"
	"html/template"
	"net/http"
	"strings"
)

type Search struct {
	Content []string
	Headers []string
}

func StringFormat(str string) string {
	str = strings.ToLower(str)
	str = strings.Replace(str, " ", "_", -1)
	return str
}

func (search Search) SearchMovie(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	views, err := template.ParseFiles("./Templates/Movies/search.html", "./Templates/Head/head.html", "./Templates/navbar.html")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	movie_name := r.FormValue("search")

	movie_name_format := StringFormat(movie_name)

	if len(movie_name_format) > 1 {
		content := Webscraping.ScrapInfo(movie_name_format)
		headers := []string{"Rating", "Genre", "Language", "Director", "Producer", "Writer", "Release Date(T)", "Release Date(S)", "Box Office", "Runtime", "Distributor", "Sound Mix", "Aspect Ratio"}
		content_str := Search{Content: content, Headers: headers}
		views.ExecuteTemplate(w, "search", content_str)
		return
	}

	views.ExecuteTemplate(w, "search", nil)
}
