package lib

import (
	"context"
	"errors"
	"log"
	"os"
	"os/exec"
	"time"
)

//Repo Contains repo fields
type Push struct {
	Branch string
	Repo   Repository
}

type Repository struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Url      string `json:"url"`
}

// Pull execute git pull to path and base on repo
func Pull(path string, branch string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", "reset", "--hard", "HEAD")
	cmd.Dir = path
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		er := errors.New("failed execute command : " + err.Error())
		log.Println(er.Error())
		return er
	}
	cmd = exec.CommandContext(ctx, "git", "pull", "--no-ff", "origin", branch)
	cmd.Dir = path
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		er := errors.New("failed execute command : " + err.Error())
		log.Println(er.Error())
		return er
	}
	return nil
}

func Clone(reponame, rootdir string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := os.Stat(rootdir + "repohook"); os.IsNotExist(err) {
		cmd := exec.CommandContext(ctx, "git", "clone", Cfg.Repo)
		cmd.Dir = rootdir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}
	} else {
		cmd := exec.CommandContext(ctx, "git", "fetch", "--all", "--tags", "--prune")
		cmd.Dir = rootdir + reponame
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
