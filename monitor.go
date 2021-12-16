package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	name := "Wellington Cruz"
	var version float32 = 1.1

	fmt.Println("---- MONITOR ----")
	fmt.Println("version ", version)
	fmt.Println()

	saudacao(name)

	for {
		exibeMenu()
		comando := leComando()
		fmt.Println()

		switch comando {
		case 1:
			inicarMonitoramento()
		case 2:
			imprimirLogs()
		case 0:
			sairDoPrograma()
		default:
			fmt.Println("Comando não reconhecido ...")
			os.Exit(-1)
		}
	}
}

func saudacao(name string) {
	fmt.Println("Olá", name)
	fmt.Println()
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	return comandoLido
}

func exibeMenu() {
	fmt.Println()
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")
	fmt.Println()
}

func sairDoPrograma() {
	os.Exit(0)
}

func inicarMonitoramento() {
	fmt.Println("Monitorando ...")
	sites := lerSitesDoArquivo()

	/**
	for i := 0; i < len(sites); i++ {
		fmt.Println(sites[i])
		testarSite(sites[i])
	}
	*/
	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando Site", i, "-", site)
			testarSite(site)
			fmt.Println()
		}
		time.Sleep(delay * time.Second)
	}
}

func imprimirLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}

func testarSite(site string) {
	resposta, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resposta.StatusCode == 200 {
		fmt.Println("Site:", site, "- Site OK")
		registrarLog(site, true)
	} else {
		fmt.Println("Site:", site, "- Site Com Erro", resposta.StatusCode)
		registrarLog(site, false)
	}
}

func lerSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registrarLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	//https://go.dev/src/time/format.go
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " " + site + " - Online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}
