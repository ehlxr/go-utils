package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	log.SetOutput(os.Stdout)

	shell := flag.String("s", "", "shell Command or shell file path")
	port := flag.Int("p", 21017, "port")
	host := flag.String("host", "0.0.0.0", "Bind TCP Address")
	flag.Parse()

	if *shell == "" {
		log.Fatalf("No shell Command specified")
	}
	shells := loadData(*shell)

	log.Println("******************************************************************")
	log.Printf("**  %-60s**", "Exec shell Server")
	log.Printf("**  Listen on %-50s**", fmt.Sprintf("%s:%d...", *host, *port))

	addr := fmt.Sprintf("%s:%d", *host, *port)
	if strings.Contains(addr, "0.0.0.0") {
		addr = strings.Replace(addr, "0.0.0.0", "", 1)
		*host = strings.Replace(*host, "0.0.0.0", "127.0.0.1", 1)
	}
	log.Printf("**  Shell Command: %-44s **", fmt.Sprintf("%q", shells))
	log.Printf("**  You can use %-48s**", fmt.Sprintf("http://%s:%d exec the shell", *host, *port))
	log.Println("******************************************************************")

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		for _, str := range shells {
			exec_shell(str)
		}
		res.Write([]byte("Done!"))
	})

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe：", err)
	}
}

func exec_shell(s string) {
	log.Printf("exec shell: %s", s)

	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatalf("exec shell error: %v", err)
	}
	fmt.Printf("%s", out.String())
}

func loadData(p string) []string {
	var values []string

	is, path := isPath(p)
	if is {
		fi, err := os.Open(path)
		if err != nil {
			log.Fatalf("Error: %s\n", err)
		}
		defer fi.Close()

		br := bufio.NewReader(fi)
		for {
			line, _, err := br.ReadLine()
			if err != nil || err == io.EOF {
				break
			}
			if strings.TrimSpace(string(line)) == "" {
				continue
			}
			values = append(values, string(line))
		}
	} else {
		values = append(values, p)
	}
	return values
}

func isPath(path string) (bool, string) {
	absPath, err := filepath.Abs(path)
	_, err = os.Stat(absPath)
	if err == nil {
		return true, absPath
	}
	return false, ""
}
