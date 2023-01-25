package GoTemplates

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type Dashboard struct{}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles("./Templates/index.html", "./Templates/Head/head.html", "./Templates/navbar.html")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	view.ExecuteTemplate(w, "index", nil)
}
