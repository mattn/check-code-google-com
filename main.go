package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/daviddengcn/go-colortext"
	"os"
	"os/exec"
	"strings"
)

var verbose = flag.Bool("v", false, "vebose")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage of %s: [package]\n", os.Args[0])
	}
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	name := flag.Arg(0)

	walk := map[string]bool{}
	walk[name] = false

	code_google_com := map[string]bool{}
	found := false

	b, err := exec.Command("go", "list", "std").CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, v := range strings.Split(string(b), "\n") {
		walk[v] = true
	}

	for {
		updated := false
		for k, v := range walk {
			if v {
				continue
			}
			if *verbose {
				ct.ChangeColor(ct.Yellow, false, ct.None, true)
				fmt.Println("checking:", k)
				ct.ResetColor()
			}
			b, err := exec.Command("go", "list", "-json", name).CombinedOutput()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				walk[k] = true
				continue
			}
			var deps struct {
				Deps []string
			}
			err = json.Unmarshal(b, &deps)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				walk[k] = true
				continue
			}
			for _, dep := range deps.Deps {
				if _, ok := walk[dep]; !ok {
					walk[dep] = false
					if strings.HasPrefix(dep, "code.google.com/") {
						code_google_com[dep] = true
						found = true
						if *verbose {
							ct.ChangeColor(ct.Magenta, false, ct.None, true)
							fmt.Println("  " + dep)
							ct.ResetColor()
						}
					}
				}
			}
			walk[k] = true
			updated = true
		}
		if !updated {
			break
		}
	}

	if found {
		ct.ChangeColor(ct.White, false, ct.Red, true)
		fmt.Printf("%s is depend on below's packages on code.google.com:", name)
		ct.ResetColor()
		fmt.Println()
		for k := range code_google_com {
			fmt.Println(k)
		}
	} else {
		ct.ChangeColor(ct.Black, false, ct.Green, true)
		fmt.Printf("%s is not depend on packages on code.google.com", name)
		ct.ResetColor()
		fmt.Println()
	}
}
