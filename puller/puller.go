package puller

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

//Repo Contains repo fields
type Push struct {
	Branch string
	Repo   Repo
}

type Repo struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Url      string `json:"url"`
}

// ChecRequest checking request will return repo and error
func CheckRequest(branch string, req map[string]interface{}) (Push, error) {
	var push Push
	if val, ok := req["ref"]; ok {
		ref := strings.Split(val.(string), "/")
		if len(ref) == 3 {
			if ref[2] != branch {
				return push, errors.New("branch is not" + branch)
			}
			push.Branch = ref[2]
		}
		er := errors.New("failed get branch")
		return push, er
	}

	if val, ok := req["repository"]; ok {
		err := json.Unmarshal([]byte(val.(string)), &push.Repo)
		if err != nil {
			er := errors.New("failed get repo detail")
			return push, er
		}
	}

	return push, nil
}

// Pull execute git pull to path and base on repo
func Pull(path string, push Push) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", "pull", "--no-ff", "origin", push.Branch)
	cmd.Dir = path
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	log.Println("running command..")
	log.Println(cmd.Args)
	log.Println("=================")
	err := cmd.Run()
	if err != nil {
		er := errors.New("failed execute command : " + err.Error())
		log.Println(er.Error())
		return er
	}
	log.Println("=================")
	log.Println("finish command..")
	return nil
}

func Postevent(postscript string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if postscript == "" {
		return errors.New("no postscript defined")
	} else {
		cmd := exec.CommandContext(ctx, "/bin/sh", "-c", postscript)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		log.Println("running command..")
		log.Println(cmd.Args)
		log.Println("=================")
		err := cmd.Run()
		if err != nil {
			return errors.New("failed execute command : " + err.Error())
		}
		log.Println("=================")
		log.Println("finish command..")
	}
	return nil
}
