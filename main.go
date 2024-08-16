package main

import (
	"log"
	"luwa-server/utils"
	"net/http"

	"luwa-server/controllers"
)

func main() {
	db := utils.ConnectDB()
	log.Println("Connected to the database")
	defer db.Close()

	utils.InitDB(db)

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		utils.Render(w, "pong.page.gohtml", nil)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.Render(w, "index.page.gohtml", nil)
	})

	mux.HandleFunc("/users/create", controllers.CreateUser)
	mux.HandleFunc("/user/", controllers.GetUserByEmail)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
