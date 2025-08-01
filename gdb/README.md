# Debugando com GDB - GNU Project Debugger

## Links úteis

- [gdb](https://sourceware.org/gdb/) Documentação oficial
- [blau_araujo](https://bolha.dev/blau_araujo/gdb-pratico) Repositório do Blau que eu estou tirando como base de aprendizado

## Preparar binário/executável para ser um arquivo 'debugavel' _*(arquivo com simbolos de depuração)*_

O compilador GCC permite gerar um executável que pode ser um arquivo debugadvel (adicionar os simbolos de depuração), mas para isso é necessário gerar este executável especifico.

> Sem os símbolos de depuração, nós teríamos que descobrir quais seriam os endereços de dados em variáveis e de outros elementos do código para inspecionar seus valores e comportamentos.

```bash
# Comando para gerar o binario executavel debugavel

gcc -o <bin_name> -g <filename.c>
```

## Iniciando o GDB

Depois de gerar o executável é necessário utilizar o arquivo com a ferramenta gdb para começar a debugar

```bash
# gdb é a ferramenta de debug
# -q é a flag para tirar as mensagens inicias que aparecem por padrão ao iniciar o debugger, são copyright
gdb -q <bin_name>
```

## Encontrar instruções _*(help)*_
```sh
# O comando help (ou apenas h) teremos uma lista de instruções do gdb
help || h

# O comando apropos <word> podemos buscar instruções que casem com a <word> que passarmos para ele, tipo um Ctrl + F
apropos help
```

## Sair do gdb

```sh
q
```

## Mostrar o código na tela

> Por padrão o comando `list` mostra apenas as 10 primeiras linhas, mas apertando ENTER podemos mostrar o restante do código

```sh 
# Mostrar o código
list

# Mostrar o código do inicio
list .
```

## Adicionando breakpoints (pontos de parada de execução)

> Os breakpoins são onde o código em execução irá parar(pausar) a execução do programa para a análise que desejamos fazer.

```sh 
# <point> é onde queremos adicionar um breakpoint, exemplo função `main`
break main
```

## Executando o programa até ele pausar no breakpoint

```sh 
# Executa o código
run
```

## Continuar a execução linha a linha

> Ao pausar a execução em um breakpoint, podemos continuar a execução linha a linha do código

```sh 
# Comando que passa para a próxima linha
next
```

## Vendo valores que foram definidos nas variáveis

```sh 
# Podemos ver o valor declarado de um variavel com o comando print <name>
print x
```
