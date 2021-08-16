package main

import (
	"net/http"

	"github.com/mbaraa/ligma/controllers"
	"github.com/mbaraa/ligma/data"
)

func main() {
	formsStore := data.NewJSONSource("./res/jsons/")
	cont := controllers.NewFormsController(formsStore)
	println("running at http://localhost:8080")
	http.Handle("/forms/", cont)
	http.ListenAndServe(":8080", nil)
}
