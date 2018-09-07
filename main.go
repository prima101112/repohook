package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/github", GithubHandler)
	log.Println("start server at port 9999")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("cant start server : ", err.Error())
	}
}
