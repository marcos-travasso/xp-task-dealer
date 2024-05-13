package main

import (
	"log"
	"net/http"
	"xp-task-dealer/core"
	"xp-task-dealer/core/openai_dealer"
	"xp-task-dealer/core/sqlite_store"
)

var dbStore core.Storer
var dealer core.Dealer

func main() {
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/", fs)

	dbStore = sqlite_store.Init("./xp-task-dealer.db")
	dealer = openai_dealer.Init()

	setupDevelopersRoutes()
	setupTasksRoutes()
	setupSuggestionsRoutes()

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
