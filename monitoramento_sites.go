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
	//"reflect" usando esta biblioteca (reflect.TypeOf)(variavel ou array), podemos identificar o tipo de variavel ou tipo e tamanho do array ou o tipo de slice)
)

const monitoramentos = 3
const delay = 5

func main() {

	//:= usado para deixar o código enxuto, não sendo necessário iniciar a variável com "var"
	//_ para quando tem uma função que retorna multiplos valores

	exibeIntroducao()

	for {

		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimiLogs()
		case 0:
			sairDoPrograma()
		default:
			fmt.Println("Não conheço este comando.")
			os.Exit(-1)
		}

		/* 	if comando == 1 {
		   		fmt.Println("Monitorando...")
		   	} else if comando == 2 {
		   		fmt.Println("Exibindo logs...")
		   	} else if comando == 0 {
		   		fmt.Println("Saindo do programa")
		   	} else {
		   		fmt.Println("Não conheço este comando")
		   	} */

	}
}

func exibeIntroducao() {
	nome := "Victor"
	versao := 1.2
	fmt.Println("Olá, sr.", nome)

	fmt.Println("Este programa está na versão:", versao )
	fmt.Println("")
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir os logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido) //Scan não precisa inferir o tipo da variavel que será inserida
	//fmt.Println("O endereço da minha variável comando é", &comando) //& é usado na frente da variavel para pegar o endereço dela no sistema
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func exibirLogs() {
	fmt.Println("Exibindo logs...")
}

func sairDoPrograma() {
	fmt.Println("Saindo do programa...")
	os.Exit(0)
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	// sites := []string{"https://random-status-code.herokuapp.com/",
	// 	"https://www.alura.com.br", "https://www.caelum.com.br",
	// 	"https://fullfacelab.com/FFFaceRecognition_nc/tokapi"}

	sites := leSitesDoarquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testaSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoarquivo() []string {

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
		fmt.Println(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
			//fmt.Println("Ocorreu um erro:", err)
		}

	}
	arquivo.Close()
	return sites
}

// restante do código omitido

// restante do código omitido

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site +
		" - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimiLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))

}

/*
func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Bernardo"}
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	nomes = append(nomes, "Aparecida")

	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

}
*/
