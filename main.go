package main

import (
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "pong", 200)
	return
}

func main() {
	http.HandleFunc("/github", GithubHandler)
	http.HandleFunc("/ping", ping)
	log.Println("start server at port 9999")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("cant start server : ", err.Error())
	}
}
