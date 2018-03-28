package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/manifoldco/promptui"
)

func main() {
	// 	source := flag.String("source", "/Users/ehlxr/ehlxr/blog/Hexo/source/resume/index.md", "source file path")
	// 	dest := flag.String("dest", "/Users/ehlxr/ehlxr/blog/resume/data.json", "destination file path")
	// 	flag.Parse()

	// 	fmt.Printf("is these right? (n/no cancel)\n source file path: %s \n destination file path: %s\n", *source, *dest)
	// var in string
	// fmt.Scanln(&in)
	// if in == "no" || in == "n" {
	// 	fmt.Println("bye!")
	// 	os.Exit(0)
	// }

	prompt := promptui.SelectWithAdd{
		Label:    "What's your source file path",
		Items:    []string{"/Users/ehlxr/ehlxr/blog/Hexo/source/resume/index.md"},
		AddLabel: "Other",
	}

	_, source, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	prompt = promptui.SelectWithAdd{
		Label:    "What's your destination file path",
		Items:    []string{"/Users/ehlxr/ehlxr/blog/resume/data.json"},
		AddLabel: "Other",
	}

	_, dest, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	m := make(map[string]interface{})
	m["show"] = "1"
	m["content"] = string(readFile(source))
	j, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	writeFile(dest, j)
	fmt.Println("Done !")
}

func readFile(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return fd
}

func writeFile(path string, b []byte) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	file.Write(b)
}
