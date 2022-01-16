package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
)

type fs_node struct {
	name     string
	parent   *fs_node
	children []*fs_node
	dir      bool
	exec     bool
}

var max_level = -1

func build_tree(parent *fs_node, path string, level int) {
	if parent == nil {
		return
	}

	if level >= max_level && max_level > 0 {
		return
	}

	entries, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, f := range entries {
		if f.Name()[0] == '.' {
			continue
		}

		info, err := os.Lstat(path + "/" + f.Name())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		mode := info.Mode()

		v := fs_node{name: f.Name(), parent: parent, children: []*fs_node{}, dir: f.IsDir(), exec: mode&0100 != 0}
		parent.children = append(parent.children, &v)

		if v.dir {
			build_tree(&v, path+"/"+v.name, level+1)
		}
	}
}

func traverse_tree(parent *fs_node, level int) {
	if parent == nil {
		return
	}

	if level >= max_level && max_level > 0 {
		return
	}

	for idx, v := range parent.children {

		for i := 0; i < level; i++ {
			fmt.Print("    ")
		}

		if idx == 0 {
			if level == 0 {
				fmt.Print("\u250c\u2500")
			} else {
				fmt.Print("\u252c\u2500")
			}
		} else if idx == len(parent.children)-1 {
			fmt.Print("\u2514\u2500")
		} else {
			fmt.Print("\u251c\u2500")
		}

		fmt.Print(" ")
		if v.dir {
			color.Blue(v.name)
			traverse_tree(v, level+1)
		} else if v.exec {
			color.Red(v.name)
		} else {
			fmt.Print(v.name + "\n")
		}
	}
}

func main() {

	var path = "."

	argc := len(os.Args)

	if argc >= 2 {
		path = os.Args[argc-1]

		help := flag.Bool("help", false, "Display help")
		level := flag.Int("level", -1, "The depth of folder structure to print")
		flag.Parse()

		max_level = *level
		if *help {
			fmt.Println("Usage for: gotree\n")
			flag.PrintDefaults()

			os.Exit(0)
		}
	}

	// do the traversal
	root := fs_node{name: ".", parent: nil, children: []*fs_node{}, dir: true, exec: false}
	build_tree(&root, path, 0)
	traverse_tree(&root, 0)

}
