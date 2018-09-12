package puller

import (
	"context"
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

//Repo Contains repo fields
type Repo struct {
	Branch string
}

// ChecRequest checking request will return repo and error
func CheckRequest(branch string, req map[string]interface{}) (Repo, error) {
	var rep Repo
	if val, ok := req["ref"]; ok {
		ref := strings.Split(val.(string), "/")
		if len(ref) == 3 {
			rep.Branch = ref[2]
			if rep.Branch != branch {
				return rep, errors.New("branch is not" + branch)
			}
			return rep, nil
		}
		er := errors.New("failed get branch")
		return rep, er
	}
	er := errors.New("failed get branch")
	return rep, er
}

// Pull execute git pull to path and base on repo
func Pull(path string, repo Repo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", "pull", "--no-ff", "origin", repo.Branch)
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if postscript == "" {
		return errors.New("no postscript defined")
	} else {
		cmd := exec.CommandContext(ctx, "sh", postscript)
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

func PostFunction(postfunction func()) {
	postfunction()
}
