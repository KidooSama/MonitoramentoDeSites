# Monitor de Sites

Aplicação de linha de comando desenvolvida em Go para monitorar a disponibilidade de sites através de requisições HTTP.

O projeto foi desenvolvido como forma de praticar os fundamentos da linguagem Go e conceitos como manipulação de arquivos, requisições HTTP, tratamento de erros e criação de logs.

## Funcionalidades

- Monitoramento de sites através de requisições HTTP
- Leitura dos sites a partir de um arquivo de configuração
- Verificação do status HTTP
- Registro dos resultados em arquivo de log
- Exibição dos logs pelo terminal
- Monitoramento periódico dos sites

## Tecnologias

- Go
- HTTP
- Manipulação de arquivos
- CLI (Command Line Interface)

## Como executar

Clone o repositório e execute:

```bash
go run main.go
```

Ou gere o executável:

```bash
go build
```

E execute o arquivo gerado.

## Configuração

Os sites monitorados são definidos no arquivo `sites.txt`, com um endereço por linha:

```text
https://google.com
https://github.com
https://example.com
```

## Objetivo

Este projeto faz parte do meu processo de aprendizado em Go, sendo utilizado para colocar em prática conceitos da linguagem através do desenvolvimento de uma aplicação funcional.
