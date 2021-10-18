package main

import (
	"net/http"

	"github.com/mbaraa/sheev/controllers"
	"github.com/mbaraa/sheev/data"
)

func main() {
	formsStore := data.NewJSONSource("./res/jsons/")
	cont := controllers.NewFormsController(formsStore)
	println("running at http://localhost:4200")

	http.Handle("/forms/", cont)
	http.Handle("/", http.FileServer(http.Dir("./client/")))
	http.ListenAndServe(":4200", nil)
}
