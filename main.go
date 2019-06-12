package main

import (
    "bufio"
    "os"
    "strings"
    "io/ioutil"
    "strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	print("Enter Full Absolute Path: ")
	dir, err := reader.ReadString('\n')
	if err != nil {
		panic("Error: " + err.Error())
	}

	for dir[len(dir) - 1] == '\n' || dir[len(dir) - 1] == '\r' || dir[len(dir) - 1] == '\\' {
		dir = dir[:len(dir) - 1]
	}

	files, _ := ioutil.ReadDir(dir)
	for i, file := range files {
		name := strings.Split(file.Name(), ".")
		if len(name) < 2 {
			continue;
		}

		oldPath := dir + "\\" + file.Name()
		newPath := dir + "\\" + strconv.Itoa(i) + "." + name[1]
		println("Renamed:", oldPath, " to:", newPath)
		os.Rename(oldPath, newPath)
	}
}