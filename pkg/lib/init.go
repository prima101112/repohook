package lib

import (
	"flag"
	"log"
	"os/exec"
	"strings"
)

type Config struct {
	Branch   string
	Path     string
	Repo     string
	RepoName string
	Interval int
	Restart  bool
}

var Cfg Config

func init() {
	branchPtr := flag.String("branch", "master", "branch that gonna whatch")
	pathPtr := flag.String("path", "", "path where app will pull (default current dirrectory)")
	repoPtr := flag.String("repo", "", "repository full path")
	intervalPtr := flag.Int("interval", 5, "interval puller will pull")
	flag.Parse()

	Cfg.Branch = *branchPtr
	Cfg.Path = *pathPtr
	Cfg.Repo = *repoPtr
	Cfg.Interval = *intervalPtr
	Cfg.RepoName = getRepoName(Cfg.Repo)

	//check if git on mechine exist
	_, err := exec.LookPath("git")
	if err != nil {
		log.Fatal("git not found")
	}
}

func GetConfig() Config {
	return Cfg
}

func getRepoName(a string) string {
	repostr := strings.Split(a, "/")
	return strings.Split(repostr[len(repostr)-1], ".")[0]
}
