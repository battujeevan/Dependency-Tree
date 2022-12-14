package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func main() {

	repoURL := "https://github.com/etcd-io/etcd"
	clonePath := "D:\\Test"
	Tag := "v3.5.5"

	_, err := git.PlainClone(clonePath, false, &git.CloneOptions{
		URL:               repoURL,
		Progress:          os.Stdout,
		ReferenceName:     plumbing.ReferenceName("refs/tags/" + Tag),
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	if err != nil {
		panic(err)
	}
	println("cloned")
	os.Chdir("D:\\Test")
	out, err := exec.Command("go", "mod", "graph").Output()

	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)
}
