package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const origem = "https://google.com.br"
const cacheDir = "./cache"
const port = 8080

func generateHash(input string) string {
	hash := sha1.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

func ProxyCacheHandler(w http.ResponseWriter, r *http.Request) {

	var body []byte

	// Tempo inicial da requisição
	startTime := time.Now()

	// Define o diretório do cache do recurso calculando a hash da URL
	cachePath := filepath.Join(cacheDir, generateHash(r.URL.Path))

	// Verifica se o recurso está em cache
	_, err := os.Stat(cachePath)

	// Caso não esteja, recupera o recurso do servidor
	if os.IsNotExist(err) {
		fmt.Println("Recurso não está presente em cache, buscando na origem:", r.URL.Path)

		// Constroi a URL do recurso
		url := fmt.Sprintf("%s%s", origem, r.URL.Path)

		// Realiza a requisição para a origem para buscar o recurso
		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
			log.Println("Falha ao buscar o recurso na origem:", err)
			return
		}
		defer resp.Body.Close()

		// Lê o conteúdo da resposta
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
			log.Println("Falha ao ler a resposta do servidor:", err)
			return
		}

		// Salva o arquivo em cache com o conteúdo do recurso
		ioutil.WriteFile(cachePath, body, 0644)
	} else {
		// Caso esteja em cache, lê o arquivo e retorna no response
		fmt.Println("Recurso está presente em cache:", r.URL.Path, cachePath)

		// Lê o arquivo em cache
		body, err = ioutil.ReadFile(cachePath)
		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
			log.Println("Falha ao ler o cache:", err)
			return
		}
	}

	// Tempo total da requisição
	fmt.Println(fmt.Sprintf("Tempo total da requisição para o recurso %v:  %v", r.URL.Path, time.Since(startTime)))

	// Resposta cacheada da requisição
	w.Write(body)

}

func main() {

	fmt.Println("Criando diretório de cache:", cacheDir)
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		os.Mkdir(cacheDir, os.ModePerm)
	}

	fmt.Println("Iniciando Proxy")

	http.HandleFunc("/", ProxyCacheHandler)

	fmt.Println("Proxy iniciado na porta:", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
