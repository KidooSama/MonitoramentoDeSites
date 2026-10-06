package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	"strconv"
)

const monitoramento = 2
const delay = 2

func main() {

	 exibeIntroducao()

	 for{

		exibeMenu()
		comando := comandoLido()
		switch comando{
			case 1:
				
				iniciarMonitoramento()
			case 2:
				fmt.Println("Exibindo LOGs")
				imprimeLogs()
			case 0:
				fmt.Println("Sair do Programa")
				os.Exit(0)
			default:
				fmt.Println("Nao existe isso cara")
				os.Exit(-1)
		}
	}
}
func exibeIntroducao() {
	fmt.Println("--------")
	nome := "Talles"
	ver := 1.1
	fmt.Println("Ola Mr.",nome)
	fmt.Println("Este programa está na versão",ver)
	fmt.Println("--------")
}

func exibeMenu(){
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	fmt.Println("--------")

}

func comandoLido() int{
	var comando int
	fmt.Scan(&comando)
	fmt.Println("comando escolhido:",comando)

	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leSitesArquivo()

	for i := 0; i < monitoramento; i++{

		for _,site := range sites{
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("-------")
	}
}

func testaSite(site string){

	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro ao buscar o site:",err)
		
	}
	if resp.StatusCode == 200 {
		//fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		//fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesArquivo()[]string {

	var sites []string

	arquivo,err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um  no arquivo:",err)
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
	return  sites
}

func registraLog(site string, status bool){
	arquivo, err := os.OpenFile("log.txt",os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil{
		fmt.Println("o erro é:", err)
	}
	
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05")+ " - " + site + " / online: "+ strconv.FormatBool(status) + "\n") 

	arquivo.Close()
}

func imprimeLogs(){
	arquivo, err := os.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))

}