package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [directory]\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	for _, root := range os.Args[1:] {
		filepath.Walk(root, func(p string, i os.FileInfo, err error) error {
			if err != nil {
				fmt.Fprintf(os.Stderr, "error at %s: %s\n", p, err)
				return nil
			}
			if !i.IsDir() {
				return nil
			}
			if strings.HasPrefix(i.Name(), ".") {
				return filepath.SkipDir
			}

			g := filepath.Join(p, ".git")
			if stat, err := os.Stat(g); err == nil && stat.IsDir() {
				fmt.Println(p)
				return filepath.SkipDir
			}

			return nil
		})
	}
}
