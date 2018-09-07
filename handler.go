package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
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

	if res["ref"] == "refs/heads/"+BRANCH {
		pull()
		postevent()
		log.Println("recheive git push : ok")
		http.Error(w, "ok", 200)
		return
	}
}

func pull() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", "pull", "--no-ff", "origin", BRANCH)
	cmd.Dir = PATH
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	log.Println("running command..")
	log.Println(cmd.Args)
	log.Println("=================")
	err := cmd.Run()
	if err != nil {
		log.Println("failed execute command : ", err.Error())
		return
	}
	log.Println("=================")
	log.Println("finish command..")
}

func postevent() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if POSTSCRIPT == "" {
		POSTSCRIPT = "script/postevent.sh"
	}

	cmd := exec.CommandContext(ctx, "sh", POSTSCRIPT)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	log.Println("running command..")
	log.Println(cmd.Args)
	log.Println("=================")
	err := cmd.Run()
	if err != nil {
		log.Println("failed execute command : ", err.Error())
		return
	}
	log.Println("=================")
	log.Println("finish command..")
}
