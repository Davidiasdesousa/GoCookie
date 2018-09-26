package main

import (
	"fmt"
	"net/http"
)

func mains() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		http.SetCookie(res, &http.Cookie{
			Name:  "Semana",
			Value: "S39",
		})
	var cookie, err = req.Cookie("Semana")
	if err == nil {
		var cookievalue = cookie.Value
		fmt.Println(cookievalue)
	}
	})
	http.ListenAndServe(":9000", nil)
}
