# Aula 1.2 - Entrada de dados STDIN

## Instruções

### equ

A instrução ´equ´, é uma diretiva do montador (nasm), que cria um simbolo constante, que é resolvido quando compilado, parecido com a diretiva `#define` de C. Os valores contido neles não aparecem no binário e não ocupam memória, e no caso, as diretavas em segment .data, servem apenas para padronizar alguns valores que serão substituidos quando compilados, como a syuscall 0x80, o file descriptor de entrada e saída padrão, o line feed, o null e etc.

```asm
segment .data

LF equ 0xA ; Line Feed


section .data

msg db "Some string", LF ; usando a diretiva, na compilação será trocado por 0xA
```

## Mais sobre seções

A seção `.data` é a seção que corresponde a variáveis que serão inicializadas, isso significa que elas terão um valor por padrão.

Exemplo:

```asm
section .data

msg db "Hello, World!", 0xA, 0xD ; variável msg inicializada com uma string
```

A seção `.bss`, é a seção que conterá variáveis que terão seus valores atribuidos na execução, exemplo:

```asm
section .bss

name resb 1

section .text

;...

mov eax, 0x4 ; operação de escrita
mov ebx, 0x0 ; file descriptor de entrada padrão
mov ecx, name ; variavel que será atribuido com o valor da entrada da entrada
mov edx, 0x2 ; quantos bytes poderão ser atribuidos a variavel
```

## Makefile

> Fluxo de execução do Makefile:
>
> Ao rodar make ele vai para o target all, como definimos, o all depende do arquivo input.o, mas não temos ainda, então ele vai para a regra onde está a compilação do arquivo .asm para .o e aí segue o que está dentro do target, no caso, linkedita o arquivo .o, remove qualquer arquivo .o e executa o linkeditado.


```make
# variável com o valor 'input'
NAME = input

# all é o target, quando eu executar o comando make, ele executará deste all
# all: ${NAME}.o - diz que o target(all) depende do arquivo compilado 'input.o' 
all: ${NAME}.o 
	clear # limpa o shell
	ld -s -o ${NAME} ${NAME}.o # lindedita o arquivo objeto compilado
	rm -rf *.o # remove qualquer arquivo com extensão .o, arquivos compilados
	clear # limpa o shell novamente
	./${NAME} # executa o objeto binário linkeditado

# Isso significa que ele vai transformar qualquer arquivo '.asm' em '.o'
%.o: %.asm
    # compilação de código assembly com o compilador nasm
	nasm -f elf64 $< # $< - é uma variavel automatica que significa o primeiro pré-requisito (nesse caso, o arquivo '.asm')
```

