package controller

import (
	"fmt"
	"html/template"
	httprouter "httprouter-master"
	"net/http"
	"strings"
)

type Controller struct {
	TPL *template.Template
}

func NewController(TPL *template.Template) *Controller {
	return &Controller{
		TPL: TPL,
	}
}

func (c *Controller) Index(w http.ResponseWriter, req *http.Request, par httprouter.Params) {
	var semana string
	var cookienew, err = req.Cookie("Semana")
	if err == nil {
		semana = cookienew.Value
	}

	conteudo := struct {
		Titulo string
		Semana string
	}{
		Titulo: "Cookie",
		Semana: semana,
	}
	c.TPL.ExecuteTemplate(w, "index.html", conteudo)
}

func (c *Controller) Cookie(w http.ResponseWriter, req *http.Request, par httprouter.Params) {
	req.ParseForm()

	nome := "S" + strings.Join(req.Form["cookiename"], "")
	http.SetCookie(w, &http.Cookie{
		Name:  "Semana",
		Value: nome,
	})
	var cookienew, err = req.Cookie("Semana")
	if err == nil {
		var cookienewvalue = cookienew.Value
		fmt.Println(cookienewvalue)
	}

	conteudo := struct {
		Titulo string
		Semana string
	}{
		Titulo: "Cookie",
		Semana: nome,
	}
	c.TPL.ExecuteTemplate(w, "index.html", conteudo)
}
