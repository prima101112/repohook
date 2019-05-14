package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prima101112/repohook/pkg/lib"
)

func ping(w http.ResponseWriter, r *http.Request) {
	if lib.GetConfig().Restart {
		http.Error(w, "restart", 503)
	} else {
		http.Error(w, "pong", 200)
	}
	return
}

func main() {
	http.HandleFunc("/ping", ping)
	log.Println("start server at port 9999")
	err := lib.Clone(lib.GetConfig().Repo, lib.GetConfig().Path)
	log.Println(err)
	path := lib.GetConfig().Path + "/" + lib.GetConfig().RepoName
	branch := lib.GetConfig().Branch
	log.Println(path + " :: " + branch)
	for {
		lib.Pull(path, branch)
		time.Sleep(time.Duration(lib.GetConfig().Interval) * time.Second)
	}
	err = http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("cant start server : ", err.Error())
	}

}
