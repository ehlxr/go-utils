package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.SelectWithAdd{
		// Label:    "What's your source file path",
		// Items:    []string{"/Users/ehlxr/ehlxr/blog/Hexo/source/resume/index.md"},
		AddLabel: "Input your source file path",
	}

	_, source, err := prompt.Run()
	if err != nil {
		fmt.Printf("prompt failed %v\n", err)
		return
	}

	files := getFiles(source)
	for _, file := range files {
		// has, err := hasSlug(file)
		// if err != nil {
		// 	fmt.Printf("hasSlug file: %s failed: %v\n", file, err)
		// 	continue
		// }
		// if has {
		// 	println("file have content slug already")
		// 	continue
		// }
		//
		// bytes, err := handleText(file)
		// if err != nil {
		// 	fmt.Printf("handleText file: %s failed: %v\n", file, err)
		// 	continue
		// }
		//
		// if len(bytes) > 0 {
		// 	err = writeFile(file, bytes)
		// 	if err != nil {
		// 		fmt.Printf("writeFile file: %s failed: %v\n", file, err)
		// 		continue
		// 	}
		//
		// 	fmt.Printf("deal file: %s done!!!\n", file)
		// }

		desc, err := genDesc(file)
		if err != nil {
			fmt.Printf("genDesc file: %s failed: %v\n", file, err)
		}

		bytes, err := genNew(file, string(desc))
		if err != nil {
			fmt.Printf("genNew file: %s failed: %v\n", file, err)
		}

		if len(bytes) > 0 {
			err = writeFile(file, bytes)
			if err != nil {
				fmt.Printf("writeFile file: %s failed: %v\n", file, err)
				continue
			}

			fmt.Printf("deal file: %s done!!!\n", file)
		}
	}
}

func hasSlug(file string) (bool, error) {
	f, err := os.OpenFile(file, os.O_RDONLY, 0644)
	if err != nil {
		return false, err
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		return false, err
	}

	if strings.Index(string(fd), "slug:") > -1 {
		return true, nil
	}

	return false, nil
}

func handleText(path string) ([]byte, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Printf("cannot open text file: %s, err: [%v]", path, err)
		return nil, err
	}
	defer file.Close()

	output := make([]byte, 0)
	// 按行读取文件
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// line := scanner.Bytes()

		output = append(output, line...)
		output = append(output, []byte("\n")...)

		if strings.Index(line, "title:") > -1 {
			newByte := strings.Replace(line, "title:", "slug:", 1)
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("cannot scanner text file: %s, err: [%v]", path, err)
		return nil, err
	}

	return output, nil
}

func genDesc(path string) ([]byte, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Printf("cannot open text file: %s, err: [%v]", path, err)
		return nil, err
	}
	defer file.Close()

	// 按行读取文件
	scanner := bufio.NewScanner(file)
	var flag, ln int
	output := make([]byte, 0)
	for scanner.Scan() {
		line := scanner.Text()
		// line := scanner.Bytes()

		if line == "<!--more-->" {
			break
		}

		if flag < 2 && line == "---" {
			flag++
			continue
		}

		f := flag == 2 && line != ""

		if f && ln > 0 {
			output = append(output, []byte("\n")...)
			// hello-friend 需要两个换行才能换行
			output = append(output, []byte("\n")...)
			output = append(output, []byte("\n")...)
		}

		if f {
			ln++
			line = strings.ReplaceAll(line, "'", "")
			output = append(output, line...)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("cannot scanner text file: %s, err: [%v]", path, err)
		return nil, err
	}

	return output, nil
}

func genNew(path string, des string) ([]byte, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Printf("cannot open text file: %s, err: [%v]", path, err)
		return nil, err
	}
	defer file.Close()

	// 按行读取文件
	scanner := bufio.NewScanner(file)
	var identifier int
	output := make([]byte, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if identifier < 2 && line == "---" {
			identifier++
		}

		if identifier == 2 && line == "---" {
			identifier++

			output = append(output, fmt.Sprintf("description: '%s'", des)...)
			output = append(output, []byte("\n")...)
		}

		output = append(output, line...)
		output = append(output, []byte("\n")...)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("cannot scanner text file: %s, err: [%v]", path, err)
		return nil, err
	}

	return output, nil
}

func writeFile(path string, b []byte) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func getFiles(root string) (files []string) {
	err := filepath.Walk(root, func(p string, f os.FileInfo, err error) error {
		if f == nil {
			return nil
		}
		if p == root || f.IsDir() {
			return nil
		}
		files = append(files, p)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

	return files
}
