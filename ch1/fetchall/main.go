package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	file, err := os.Create("Testes.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		go fetch(url, ch, file) //Para cada url o comando go inicia uma nova goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, file *os.File) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(file, fmt.Sprint(err)+"\n")
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(file, fmt.Sprintf("while reading %s: %v\n", url, err))
		return
	}
	secs := time.Since(start).Seconds()
	fmt.Fprintf(file, fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url))
}
