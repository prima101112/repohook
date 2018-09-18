package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	"github.com/prima101112/repohook/puller"
)

var BRANCH string
var PATH string
var POSTSCRIPT string

func init() {
	branchPtr := flag.String("branch", "master", "branch that gonna whatch")
	pathPtr := flag.String("path", "", "path where app will pull (default current dirrectory)")
	postscriptPtr := flag.String("script", "", "script that will execute after pull done")
	flag.Parse()
	BRANCH = *branchPtr
	PATH = *pathPtr
	POSTSCRIPT = *postscriptPtr
	_, err := exec.LookPath("git")
	if err != nil {
		log.Fatal("git not found")
	}
}

// GithubHandler handler for github
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
		log.Println("response body : ", string(b))
		http.Error(w, "failed to unmarshal", 200)
		return
	}
	push, err := puller.CheckRequest(BRANCH, res)
	err = puller.Pull(PATH, push)
	if err != nil {
		log.Println("failed to pull : ", err.Error())
		http.Error(w, "failed to pull", 200)
		return
	}

	err = puller.Postevent(POSTSCRIPT)
	if err != nil {
		log.Println("failed to execute postevent : ", err.Error())
		http.Error(w, "failed to execute postevent", 200)
		return
	}
	http.Error(w, "oke", 200)
	return
}
