package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//Read, clean and split input file
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	actions := strings.Split(strings.TrimSpace(string(input)), "\n")
	fileSystem := actionsToFileSystem(actions)
	dirSizes := fileSystemToDirSizes(fileSystem)

	part1, part2 := 0, dirSizes["/"]
	for _, size := range dirSizes {
		if size <= 100000 {
			part1 += size
		}
		if size+70000000-dirSizes["/"] >= 30000000 && size < part2 {
			part2 = size
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)

}

func actionsToFileSystem(actions []string) map[string]int {
	fileSystem := map[string]int{}
	currentPath := ""
	for _, command := range actions {
		parts := strings.Fields(command)
		if parts[1] == "cd" {
			currentPath = path.Join(currentPath, parts[2])
		} else if unicode.IsDigit(rune(parts[0][0])) {
			fileSystem[path.Join(currentPath, parts[1])], _ = strconv.Atoi(parts[0])
		}
	}
	return fileSystem
}

func fileSystemToDirSizes(fileSystem map[string]int) map[string]int {
	dirSizes := map[string]int{}
	for file, size := range fileSystem {
		directory := path.Dir(file)
		for directory != "/" {
			dirSizes[directory] += size
			directory = path.Dir(directory)
		}
		dirSizes["/"] += size
	}
	return dirSizes
}
