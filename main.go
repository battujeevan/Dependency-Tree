package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Pkg struct {
	name   string
	childs []string
}

type PkgTree struct {
	mapPkg map[string]*Pkg
	root   string
	depth  int
}

type Artifact struct {
	Name        string //`json:"name"`
	Version     string //`json:"color"`
	Dependecies []Artifact
}

var Arr []Artifact
var count int

func NewPkgTree(depth int) *PkgTree {
	return &PkgTree{mapPkg: make(map[string]*Pkg), depth: depth}
}

func (p *PkgTree) Add(name string, child string) {
	pkg, ok := p.mapPkg[name]
	if !ok {
		pkg = &Pkg{name: name}
		p.mapPkg[name] = pkg
	}

	pkg.childs = append(pkg.childs, child)

	if len(p.root) == 0 {
		p.root = name
	}
}

func (p *PkgTree) GetPkg(name string) *Pkg {
	return p.mapPkg[name]
}

func (p *PkgTree) GetRootPkg() *Pkg {
	if len(p.root) == 0 {
		return nil
	}
	return p.mapPkg[p.root]
}

func (p *PkgTree) printTree(name string, Arr1 []Artifact) int {

	if count > p.depth {
		count = 0
		return -1
	}
	var s []Artifact
	contains := strings.Contains(name, "/")
	if contains {
		res := strings.Split(name, "/")
		st := Artifact{res[len(res)-2], res[len(res)-1], s}
		Arr = append(Arr, st)
	} else {
		res := strings.Split(name, "@")
		st := Artifact{res[0], res[1], s}
		Arr = append(Arr, st)
	}

	child, ok := p.mapPkg[name]
	if ok && len(child.childs) > 0 {
		for _, name := range child.childs {
			count = count + 1
			if len(Arr) > 0 {
				stparent := Arr[len(Arr)-1]
				var s []Artifact
				contains := strings.Contains(name, "/")
				if contains {
					res := strings.Split(name, "/")
					st := Artifact{res[len(res)-2], res[len(res)-1], s}
					stparent.Dependecies = append(stparent.Dependecies, st)
				} else {
					res := strings.Split(name, "@")
					st := Artifact{res[0], res[1], s}
					stparent.Dependecies = append(stparent.Dependecies, st)
				}
				Arr[len(Arr)-1].Dependecies = stparent.Dependecies
				if p.printTree(name, stparent.Dependecies) == -1 {
					break
				}
			}

		}
	}

	return 0
}

var pDepth = flag.Int("d", 18, "max depth of dependence")

func main() {

	graphFile := "C:\\Users\\jeevan_b1\\Desktop\\Input.txt"

	var err error
	var file *os.File
	if len(graphFile) == 0 {
		file = os.Stdin
	} else {
		file, err = os.Open(graphFile)
		if err != nil {
			log.Fatal(err)
		}
	}
	reader := bufio.NewReader(file)
	pkgTree := NewPkgTree(*pDepth)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		ss := strings.Split(string(line), " ")
		if len(ss) != 2 {
			if len(ss) == 1 {
				pkgTree.Add(ss[0], ss[0])
			} else {
				log.Fatal(errors.New("error input"))
			}
		} else {
			pkgTree.Add(ss[0], ss[1])
		}
	}

	root := pkgTree.GetRootPkg()
	if root == nil {
		return
	}

	fmt.Println("package:", root.name)
	fmt.Println("dependence tree:\n")
	for _, c := range root.childs {
		pkgTree.printTree(c, Arr)
	}

	b, err := json.MarshalIndent(Arr, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
