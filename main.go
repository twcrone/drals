package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

type DirAlias struct {
	Alias   string
	Dir     string
}

func main() {
	homeDir, _ := os.UserHomeDir()
	filename := homeDir + "/.drals"
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	dirAliasMap := make(map[string]DirAlias)
	for scanner.Scan() {
		var dirAlias = parse(scanner.Text())
		dirAliasMap[dirAlias.Alias] = dirAlias
	}

	args := os.Args
	if len(args) > 1 {
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
			return
		}
		alias := args[1]
		dir := fmt.Sprintf("'%s'", path)
		dirAlias := DirAlias{Alias: alias, Dir: dir}
		dirAliasMap[alias] = dirAlias

		list := listFrom(dirAliasMap)
		sort.Sort(byAlias(list))

		file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		check(err)
		writer := bufio.NewWriter(file)
		for _, dirAlias := range list {
			alias := fmt.Sprintf("alias %s=%s\n", dirAlias.Alias, dirAlias.Dir)
			_, err := writer.WriteString(alias)
			fmt.Print("Writing ", alias)
			check(err)
		}
		err = writer.Flush()
		check(err)
	}
	err := file.Close()
	check(err)

	dat, _ := ioutil.ReadFile(filename)
	fmt.Println()
	fmt.Println("Current list of directory aliases")
	fmt.Println("---------------------------------")
	fmt.Println(string(dat))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type byAlias []DirAlias

func (a byAlias) Len() int {
	return len(a)
}

func (a byAlias) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byAlias) Less(i, j int) bool {
	return a[i].Alias < a[j].Alias
}

func listFrom(dirAliasMap map[string]DirAlias) []DirAlias {
	values := make([]DirAlias, 0, len(dirAliasMap))

	for _, v := range dirAliasMap {
		values = append(values, v)
	}

	return values
}

func parse(line string) DirAlias {
	segments := strings.Split(line, "=")
	alias := segments[0][6:]
	return DirAlias{Alias: alias, Dir: segments[1]}
}