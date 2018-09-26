package main

import (
	_ "fmt"
	"html/template"
	"httprouter-master"
	"log"
	"net/http"

	ctrl "github.com/GoSession/controller"
)

var Tpl *template.Template

func init() {
	Tpl = template.Must(template.New("Coockie").Delims("[[", "]]").ParseGlob("view/*.html"))
}
func main() {
	rota := ctrl.NewController(Tpl)
	httprouter := httprouter.New()
	httprouter.ServeFiles("/view/*filepath", http.Dir("view"))

	httprouter.GET("/", rota.Index)
	httprouter.POST("/cookie", rota.Cookie)

	log.Fatal(http.ListenAndServe(":3000", httprouter))
}
