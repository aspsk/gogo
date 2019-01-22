package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"strconv"
	"io/ioutil"
	"strings"
	"os"
	"math/rand"
)

func debug(format string, args ...interface{}) {
	log.Printf("[debug] " + format + "\n", args...)
}

func fatal(format string, args ...interface{}) {
	log.Fatalf(format + "\n", args...)
}

func permute(names []string, perm []int) error {
	prefix := "wokka__" // "unique" prefix

	for _, name := range names {
		err := os.Rename(name, prefix + name)
		if err != nil {
			// XXX: restore
			return fmt.Errorf("os.Rename: %s -> %s: %v", name, prefix+name, err)
		}
	}

	for i, j := range perm {
		err := os.Rename(prefix + names[i], names[j])
		if err != nil {
			// XXX: restore
			return fmt.Errorf("os.Rename: %s -> %s: %v", prefix + names[i], names[j], err)
		}
	}

	return nil
}

func inversePerm(perm []int) []int {
	N := len(perm)
	ret := make([]int, N)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if perm[j] == i {
				ret[i] = j
			}
		}
	}
	return ret
}

func savePerm(path string, names []string, perm []int) error {

	if len(names) != len(perm) {
		return fmt.Errorf("len(names)[%d] != len(perm)[%d]", len(names), len(perm))
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("os.Create: saved.perm: %v", err)
	}
	defer f.Close()

	fmt.Fprintln(f, len(names))
	for _, name := range names {
		fmt.Fprintln(f, name)
	}
	for _, x := range perm {
		fmt.Fprintln(f, x)
	}

	return nil
}

func restorePerm(path string) ([]string, []int, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, nil, fmt.Errorf("os.Open: %s: %v", path, err)
	}
	defer f.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("scanner: %s: %v", path, err)
	}

	debug("lines=%v", lines)

	N, err := strconv.Atoi(lines[0])
	if err != nil {
		return nil, nil, fmt.Errorf("strconv.Atoi: %s: %v", lines[0], err)
	}
	if len(lines) != 1 + 2 * N {
		return nil, nil, fmt.Errorf("wrong number of lines: %d [%d expected]", len(lines), 1 + 2 * N)
	}

	names := make([]string, N)
	perm := make([]int, N)
	for i := 0; i < N; i++ {
		names[i] = lines[i + 1]
		var err error
		perm[i], err = strconv.Atoi(lines[i + N + 1])
		if err != nil {
			return nil, nil, fmt.Errorf("strconv.Atoi: %s: %v", lines[i + N + 1], err)
		}
	}

	return names, perm, nil
}

func gimmeNames(target string, permuteFiles bool) ([]string, error) {

	if err := os.Chdir(target); err != nil {
		return nil, fmt.Errorf("os.Chdir: %s: %v", target, err)
	}

	names, err := ioutil.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadDir: %v", err)
	}

	var ret []string
	for _, f := range names {
		name := f.Name()

		// special case: we are skipping files which we, presumably, created
		if strings.HasSuffix(name, ".perm") {
			debug("skipping permutation %s", name)
			continue
		}

		// xor
		if f.IsDir() != permuteFiles {
			ret = append(ret, name)
		}
	}

	// ternary operator, Go style
	var what string
	if permuteFiles {
		what = "files"
	} else {
		what = "directories"
	}
	debug("Going to permute the following %s in the directory '%s': %v", what, target, ret)

	return ret, nil
}

func gimmePerm(permName string, names []string, inverse bool) ([]int, []string, error) {

	var perm []int

	if permName == "" {
		perm = rand.Perm(len(names))
	} else {
		var err error
		names, perm, err = restorePerm(permName)
		if err != nil {
			return nil, nil, fmt.Errorf("restorePerm: %v", err)
		}
	}
	debug("perm = %v", perm)

	if inverse {
		perm = inversePerm(perm)
		debug("inverse perm = %v", perm)
	}

	return perm, names, nil
}

func main() {

	var target = flag.String("target", "", "target directory")
	var permuteDirs = flag.Bool("dirs", false, "permute directories")
	var permuteFiles = flag.Bool("files", false, "permute files")
	var save = flag.String("save", "", "Path to save permutation (relative to target)")
	var permName = flag.String("perm", "", "target permutation")
	var inverse = flag.Bool("inv", false, "apply inverse permutation")
	var seed = flag.Int64("seed", 123, "initial random seed")
	flag.Parse()

	rand.Seed(*seed)

	if *target == "" {
		fatal("Please specify the target directory")
	}

	// there can be only one
	if *permuteDirs == *permuteFiles {
		fatal("Please select one and only one of the --dirs and --files options")
	}

	// the gimmeNames also sets working directory to *target
	names, err := gimmeNames(*target, *permuteFiles)
	if err != nil {
		fatal("gimmeNames: %v", err)
	}

	perm, names, err := gimmePerm(*permName, names, *inverse)
	if err != nil {
		fatal("gimmePerm: %v", err)
	}

	if err := permute(names, perm); err != nil {
		fatal("permute: %v", err)
	}

	if *save != "" {
		if err := savePerm(*save, names, perm); err != nil {
			fatal("savePerm: %s: %v", *save, err)
		}
	}
}
