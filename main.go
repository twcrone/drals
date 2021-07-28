package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type DirAlias struct {
	Alias   string
	Dir     string
}

func main() {
	file, _ := os.Open("/Users/twcrone/.drals")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	dirAliasMap := make(map[string]DirAlias)
	for scanner.Scan() {
		var dirAlias = parse(scanner.Text())
		dirAliasMap[dirAlias.Alias] = dirAlias
	}

	file.Close()

	list := listFrom(dirAliasMap)
	sort.Sort(byAlias(list))

	for _, dirAlias := range list {
		fmt.Println(dirAlias.Alias, "->", dirAlias.Dir)
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

func list() {

}

func add(directory string) {

}
