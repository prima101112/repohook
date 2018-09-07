package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GithubHandler(w http.ResponseWriter, r *http.Request) {
	var res map[string]interface{}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("failed to read body : ", err.Error())
		http.Error(w, "failed to read body", 200)
		return
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		log.Println("failed to unmarshal : ", err.Error())
		http.Error(w, "failed to unmarshal", 200)
		return
	}

	log.Print(res)
}
