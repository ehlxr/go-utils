package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	source := flag.String("source", "/Users/ehlxr/ehlxr/blog/Hexo/source/resume/index.md", "source file path")
	dest := flag.String("dest", "/Users/ehlxr/ehlxr/blog/resume/data.json", "destination file path")
	flag.Parse()

	f, err := os.Open(*source)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	m := make(map[string]interface{})
	m["show"] = "1"
	m["content"] = string(fd)
	j, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(j))
	writeFile(*dest, j)
}

func writeFile(fn string, b []byte) {
	file, err := os.OpenFile(fn, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	file.Write(b)
}
