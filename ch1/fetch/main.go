package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

//Esse projeto busca uma url da internet, e escreve o resp.Body no terminal, ele também faz uma padronização na url
//Caso nao exista o http://, alé, de buscar e fazer o tratamento de erros, também nos devolve o status da req.

func main() {

	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprint(os.Stderr, "fetch %v\n", err)
			os.Exit(1)
		}

		w, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch %v\n", err)
			os.Exit(1)
		}

		//b, err := io.ReadAll(resp.Body)
		defer resp.Body.Close()

		fmt.Printf("%s\n", string(w))
		fmt.Println("Status:", resp.Status)

	}

}
