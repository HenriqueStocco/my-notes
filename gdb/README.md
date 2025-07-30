# Debugando com GDB - GNU Project Debugger

## Links úteis

- [gdb](https://sourceware.org/gdb/) Documentação oficial
- [blau_araujo](https://bolha.dev/blau_araujo/gdb-pratico) Repositório do Blau que eu estou tirando como base de aprendizado

## Preparar binário/executável para ser um arquivo 'debugavel'

O compilador GCC permite gerar um executável que pode ser um arquivo debugadvel, mas para isso é necessário gerar este executável especifico.

```bash
# Comando para gerar o binario executavel debugavel

gcc -o <bin_name> -g <filename.c>
```

## Iniciando o GDB

Depois de gerar o executável é necessário utilizar o arquivo com a ferramenta gdb para começar a debugar

```bash
# gdb é a ferramenta de debug
# -q é a flag para tirar as mensagens inicias que aparecem por padrão ao iniciar o debugger
gdb -q <bin_name>
```
